package converter

import (
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToPermissionFromRepo(permission *entities.Permission) *model.Permission {
	if permission == nil {
		return nil
	}

	permissionModel := &model.Permission{}

	if permission.ID.Valid {
		permissionModel.ID = permission.ID.Int.Uint64()
	}

	if permission.Name.Valid {
		permissionModel.Name = permission.Name.String
	}

	if permission.Title.Valid {
		permissionModel.Title = permission.Title.String
	}

	return permissionModel
}

func ToPermissionsFromRepo(permissions []*entities.Permission) []*model.Permission {
	models := make([]*model.Permission, 0, len(permissions))
	for i := 0; i < len(permissions); i++ {
		models = append(models, ToPermissionFromRepo(permissions[i]))
	}
	return models
}
