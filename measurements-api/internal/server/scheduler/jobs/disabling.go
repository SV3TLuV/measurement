package jobs

import (
	"context"
	"log"
	model2 "measurements-api/internal/model"
	"measurements-api/internal/repository/object/model"
	"measurements-api/internal/service"
	"measurements-api/pkg/asoiza"
	"time"
)

func NewDisablingPostsJob(
	configuration service.ConfigurationService,
	measurementService service.MeasurementService,
	objectService service.ObjectService,
	asoizaClient asoiza.Client) func() {
	return func() {
		ctx := context.Background()
		config, err := configuration.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}

		day := config.DisablingThreshold / (24 * 60 * 60)
		threshold := time.Now().UTC().AddDate(0, 0, -int(day))
		postKey := uint64(model2.PostKey)
		posts, err := objectService.GetObjects(ctx, &model.GetObjectsQueryParams{
			TypeID: &postKey,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, post := range posts {
			switch post.PostInfo.IsListened {
			case true:
				measurement, err := measurementService.GetLastPostMeasurement(ctx, post.ID)
				if err != nil {
					log.Fatal(err)
				}

				if measurement != nil && measurement.Created.Before(threshold) {
					err = objectService.Disable(ctx, post.ID)
					if err != nil {
						log.Fatal(err)
					}
				}
			case false:
				measurement, err := asoizaClient.GetNewestMeasurement(ctx, post.ID)
				if err != nil {
					log.Fatal(err)
				}

				if measurement != nil && measurement.Created.After(threshold) {
					err = objectService.Enable(ctx, post.ID)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
}
