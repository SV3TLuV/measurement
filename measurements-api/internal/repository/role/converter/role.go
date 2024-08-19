package converter

import (
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToRoleFromRepo(role *entities.Role) *model.Role {
	if role == nil {
		return nil
	}

	return &model.Role{
		ID:    role.ID,
		Title: role.Title,
		Name:  role.Name,
	}
}

func ToRolesFromRepo(roles []*entities.Role) []*model.Role {
	models := make([]*model.Role, 0, len(roles))
	for i := 0; i < len(roles); i++ {
		models = append(models, ToRoleFromRepo(roles[i]))
	}
	return models
}
