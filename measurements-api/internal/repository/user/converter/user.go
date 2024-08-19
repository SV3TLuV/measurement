package converter

import (
	"github.com/doug-martin/goqu/v9"
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
	"measurements-api/internal/repository/role/converter"
)

func ToUserRecordFromService(user *model.User) *goqu.Record {
	if user == nil {
		return nil
	}

	record := goqu.Record{
		"login":         user.Login,
		"password_hash": user.Password,
		"is_blocked":    user.IsBlocked,
		"role_id":       user.Role.ID,
	}

	if user.ID > 0 {
		record["user_id"] = user.ID
	}

	return &record
}

func ToUserFromRepo(user *entities.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:        user.ID,
		Login:     user.Login,
		Password:  user.PasswordHash,
		IsBlocked: user.IsBlocked,
		Role:      *converter.ToRoleFromRepo(&user.Role),
	}
}

func ToUsersFromRepo(users []*entities.User) []*model.User {
	models := make([]*model.User, 0, len(users))
	for i := 0; i < len(users); i++ {
		models = append(models, ToUserFromRepo(users[i]))
	}
	return models
}
