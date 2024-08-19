package converter

import (
	"github.com/doug-martin/goqu/v9"
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToSessionRecordFromService(session *model.Session) *goqu.Record {
	if session == nil {
		return nil
	}

	return &goqu.Record{
		"session_id":    session.ID,
		"user_id":       session.UserID,
		"created":       session.Created,
		"updated":       session.Updated,
		"refresh_token": session.RefreshToken,
	}
}

func ToSessionFromRepo(session *entities.Session) *model.Session {
	if session == nil {
		return nil
	}

	return &model.Session{
		ID:           session.ID,
		UserID:       session.UserID,
		Created:      session.Created,
		Updated:      session.Updated,
		RefreshToken: session.RefreshToken,
	}
}

func ToSessionsFromRepo(sessions []*entities.Session) []*model.Session {
	models := make([]*model.Session, 0, len(sessions))
	for i := 0; i < len(sessions); i++ {
		models = append(models, ToSessionFromRepo(sessions[i]))
	}
	return models
}
