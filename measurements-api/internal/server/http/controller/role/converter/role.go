package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/role/model"
)

func ToRoleViewFromService(role *model.Role) *model2.RoleView {
	if role == nil {
		return nil
	}

	return &model2.RoleView{
		ID:    role.ID,
		Name:  role.Name,
		Title: role.Title,
	}
}

func ToRoleViewsFromService(roles []*model.Role) []*model2.RoleView {
	views := make([]*model2.RoleView, 0, len(roles))
	for i := 0; i < len(roles); i++ {
		views = append(views, ToRoleViewFromService(roles[i]))
	}
	return views
}
