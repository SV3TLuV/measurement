package converter

import (
	"github.com/doug-martin/goqu/v9"
	"measurements-api/internal/interfaces/converter"
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToPollingStatisticFromService(statistic *model.PollingStatistic) *entities.PollingStatistic {
	if statistic == nil {
		return nil
	}

	return &entities.PollingStatistic{
		ID:            statistic.ID,
		DateTime:      statistic.DateTime,
		Duration:      converter.ToIntervalFromDuration(statistic.Duration),
		PostCount:     statistic.PostCount,
		ReceivedCount: statistic.ReceivedCount,
	}
}

func ToPollingStatisticFromRepo(statistic *entities.PollingStatistic) *model.PollingStatistic {
	if statistic == nil {
		return nil
	}

	return &model.PollingStatistic{
		ID:            statistic.ID,
		DateTime:      statistic.DateTime,
		Duration:      converter.ToDurationFromInterval(statistic.Duration),
		PostCount:     statistic.PostCount,
		ReceivedCount: statistic.ReceivedCount,
	}
}

func ToPollingStatisticsFromRepo(statistics []*entities.PollingStatistic) []*model.PollingStatistic {
	models := make([]*model.PollingStatistic, 0, len(statistics))
	for i := 0; i < len(statistics); i++ {
		models = append(models, ToPollingStatisticFromRepo(statistics[i]))
	}
	return models
}

func ToPollingStatisticRecordFromService(statistic *model.PollingStatistic) *goqu.Record {
	if statistic == nil {
		return nil
	}

	record := goqu.Record{
		"datetime":       statistic.DateTime,
		"duration":       statistic.Duration.String(),
		"post_count":     statistic.PostCount,
		"received_count": statistic.ReceivedCount,
	}

	if statistic.ID > 0 {
		record["polling_statistic_id"] = statistic.ID
	}

	return &record
}
