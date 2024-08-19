package converter

import (
	"github.com/doug-martin/goqu/v9"
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToPostInfoRecordFromService(postInfo *model.PostInfo) *goqu.Record {
	if postInfo == nil {
		return nil
	}

	return &goqu.Record{
		"object_id":              postInfo.ObjectID,
		"last_polling_date_time": postInfo.LastPollingDateTime,
		"is_listened":            postInfo.IsListened,
	}
}

func ToPostInfoRecordsFromService(postInfos []*model.PostInfo) []*goqu.Record {
	records := make([]*goqu.Record, 0, len(postInfos))
	for i := 0; i < len(postInfos); i++ {
		records = append(records, ToPostInfoRecordFromService(postInfos[i]))
	}
	return records
}

func ToPostInfoFromRepo(info *entities.PostInfo) *model.PostInfo {
	if info == nil {
		return nil
	}

	postInfo := &model.PostInfo{
		LastPollingDateTime: info.LastPollingDateTime,
	}

	if info.ObjectID.Valid {
		postInfo.ObjectID = info.ObjectID.Int.Uint64()
	}

	if info.IsListened.Valid {
		postInfo.IsListened = info.IsListened.Bool
	}

	return postInfo
}
