package converter

import (
	"fmt"
	model3 "measurements-api/internal/model"
	model2 "measurements-api/internal/repository/polling_statistic/model"
	"measurements-api/internal/server/http/controller/collector/model"
)

func ToGetPollingStatisticsParamsFromRequest(request *model.FetchPollingStatisticQuery) *model2.GetPollingStatisticParams {
	if request == nil {
		return nil
	}

	return &model2.GetPollingStatisticParams{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
}

func ToPollingStatisticViewFromService(statistic *model3.PollingStatistic) *model.PollingStatisticView {
	if statistic == nil {
		return nil
	}

	minutes := int(statistic.Duration.Minutes())
	seconds := int(statistic.Duration.Seconds()) % 60

	return &model.PollingStatisticView{
		ID:            statistic.ID,
		DateTime:      statistic.DateTime,
		Duration:      fmt.Sprintf("%02d:%02d", minutes, seconds),
		PostCount:     statistic.PostCount,
		ReceivedCount: statistic.ReceivedCount,
	}
}

func ToPollingStatisticViewsFromService(statistics []*model3.PollingStatistic) []*model.PollingStatisticView {
	views := make([]*model.PollingStatisticView, 0, len(statistics))
	for i := 0; i < len(statistics); i++ {
		views = append(views, ToPollingStatisticViewFromService(statistics[i]))
	}
	return views
}
