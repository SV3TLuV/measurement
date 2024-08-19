package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/collector/model"
)

func ToInformationViewFromService(info *model.CollectorInformation) *model2.CollectorInformationView {
	if info == nil {
		return nil
	}

	return &model2.CollectorInformationView{
		ListenedPostCount:   info.ListenedPostCount,
		PostCount:           info.PostCount,
		PollingInterval:     info.PollingInterval,
		LastPollingDateTime: info.LastPollingDateTime,
		UntilNextPolling:    info.UntilNextPolling,
	}
}

func ToControllerStateViewFromService(state *model.CollectorState) *model2.CollectorStateView {
	if state == nil {
		return nil
	}

	return &model2.CollectorStateView{
		Status:          string(state.Status),
		PolledPostCount: state.PolledPostCount,
		PostCount:       state.PostCount,
		PollingPercent:  state.PollingPercent,
		ReceivedCount:   state.ReceivedCount,
		Started:         state.Started,
	}
}
