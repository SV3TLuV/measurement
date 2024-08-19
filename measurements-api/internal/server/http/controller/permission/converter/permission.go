package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/permission/model"
)

func ToPermissionViewFromService(permission *model.Permission) *model2.PermissionView {
	if permission == nil {
		return nil
	}

	return &model2.PermissionView{
		ID:    permission.ID,
		Name:  permission.Name,
		Title: permission.Title,
	}
}

func ToPermissionViewsFromService(permissions []*model.Permission) []*model2.PermissionView {
	views := make([]*model2.PermissionView, 0, len(permissions))
	for i := 0; i < len(permissions); i++ {
		views = append(views, ToPermissionViewFromService(permissions[i]))
	}
	return views
}
