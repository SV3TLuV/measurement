package jobs

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log"
	"measurements-api/internal/interfaces/converter"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	"measurements-api/internal/service"
	"measurements-api/pkg/asoiza"
	"time"
)

type Collector struct {
}

const count = 1024

func NewCollectMeasurementJob(
	objectService service.ObjectService,
	asoizaClient asoiza.Client,
	postInfoRepo repository.PostInfoRepository,
	measurementService service.MeasurementService,
	collectorService service.CollectorService) func() {
	return func() {
		started := time.Now()
		state := model.CollectorState{
			Status:         model.Collecting,
			Started:        &started,
			PollingPercent: uint64(0),
		}

		ctx := context.Background()
		posts, err := objectService.GetPosts(ctx)
		if err != nil {
			log.Fatal(err)
		}

		state.PostCount = uint64(len(posts))
		err = collectorService.UpdateState(ctx, state)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(posts); i++ {
			received, err := collectFromPost(ctx,
				posts[i].ID,
				asoizaClient,
				postInfoRepo,
				measurementService)
			if err != nil {
				log.Fatal(err)
			}

			state.PolledPostCount++
			state.ReceivedCount += received
			state.PollingPercent = uint64(float64(state.PolledPostCount) / float64(state.PostCount) * 100)
			err = collectorService.UpdateState(ctx, state)
			if err != nil {
				log.Fatal(err)
			}
		}

		ended := time.Now()
		state.Ended = &ended
		state.Status = model.Pending
		err = collectorService.UpdateState(ctx, state)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func collectFromPost(ctx context.Context,
	postID uint64,
	asoizaClient asoiza.Client,
	postInfoRepo repository.PostInfoRepository,
	measurementService service.MeasurementService) (uint64, error) {
	last, err := asoizaClient.GetNewestMeasurement(ctx, postID)
	if err != nil {
		return 0, errors.Wrap(err, "get last measurement from asoiza")
	}
	if last == nil {
		return 0, nil
	}

	lastInDb, err := measurementService.GetByID(ctx, last.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return 0, errors.Wrap(err, "get last measurement from database")
	}
	if lastInDb != nil {
		return 0, nil
	}

	postInfo, err := postInfoRepo.GetById(ctx, postID)
	if err != nil {
		return 0, errors.Wrap(err, "get post info from database")
	}

	updatedAt := time.Now()
	postInfo.LastPollingDateTime = &updatedAt
	err = postInfoRepo.SaveOne(ctx, postInfo)
	if err != nil {
		return 0, errors.Wrap(err, "save post info to database")
	}

	received, skip := uint64(0), 0
	for {
		measurements, err := asoizaClient.GetMeasurements(ctx, postID, count, skip)
		if err != nil {
			return 0, errors.Wrap(err, "get post measurements")
		}
		if len(measurements) == 0 {
			return 0, nil
		}

		models := converter.ToMeasurementsFromAsoiza(measurements)
		insertedCount, err := measurementService.Save(ctx, models)
		if err != nil {
			return 0, errors.Wrap(err, "save measurements")
		}
		if insertedCount != nil {
			received += *insertedCount
		}

		lastMeasurement := models[len(measurements)-1]
		if lastMeasurement.IsOld() {
			return received, nil
		}

		skip += len(measurements)
	}
}
