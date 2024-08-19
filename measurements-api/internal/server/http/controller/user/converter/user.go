package converter

import (
	"measurements-api/internal/model"
	model4 "measurements-api/internal/repository/object/model"
	model5 "measurements-api/internal/repository/user/model"
	model3 "measurements-api/internal/server/http/controller/object/model"
	"measurements-api/internal/server/http/controller/role/converter"
	model2 "measurements-api/internal/server/http/controller/user/model"
	model6 "measurements-api/internal/server/http/model"
)

func ToUserViewFromService(user *model.User) *model2.UserView {
	if user == nil {
		return nil
	}

	return &model2.UserView{
		ID:            user.ID,
		Login:         user.Login,
		IsBlocked:     user.IsBlocked,
		Role:          *converter.ToRoleViewFromService(&user.Role),
		PermissionIds: user.PermissionIds,
		ColumnIds:     user.ColumnIds,
		PostIds:       user.PostIds,
	}
}

func ToUserViewsFromService(users []*model.User) []*model2.UserView {
	views := make([]*model2.UserView, 0, len(users))
	for i := 0; i < len(users); i++ {
		views = append(views, ToUserViewFromService(users[i]))
	}
	return views
}

func ToGetObjectsParamsFromQuery(query *model3.FetchObjectListQuery) *model4.GetObjectsQueryParams {
	if query == nil {
		return nil
	}

	return &model4.GetObjectsQueryParams{
		TypeID:    query.TypeId,
		Search:    query.Search,
		ParentIds: query.ParentIds,
	}
}

func ToGetUsersParamsFromRequest(
	request *model2.FetchUserListRequest) *model5.GetUsersQueryParams {
	if request == nil {
		return nil
	}

	return &model5.GetUsersQueryParams{
		Search:   request.Search,
		Page:     request.Page,
		PageSize: request.PageSize,
		RoleIds:  request.RoleIds,
	}
}

func ToUserPagedListViewFromService(
	pagedList *model.PagedList[model.User]) *model6.PagedListView[model2.UserView] {
	if pagedList == nil {
		return nil
	}

	return &model6.PagedListView[model2.UserView]{
		Page:     pagedList.Page,
		PageSize: pagedList.PageSize,
		Total:    pagedList.Total,
		Items:    ToUserViewsFromService(pagedList.Items),
	}
}

func ToUserFromCreateRequest(request *model2.CreateUserRequest) *model.User {
	if request == nil {
		return nil
	}

	return &model.User{
		Login:         request.Login,
		Password:      request.Password,
		Role:          model.Role{ID: request.RoleID},
		PermissionIds: request.PermissionIDs,
		ColumnIds:     request.ColumnIDs,
		PostIds:       request.PostIDs,
	}
}

func ToUserFromUpdateRequest(request *model2.UpdateUserRequest) *model.User {
	if request == nil {
		return nil
	}

	userData := &model.User{
		ID:            request.UserID,
		PermissionIds: request.PermissionIDs,
		ColumnIds:     request.ColumnIDs,
		PostIds:       request.PostIDs,
	}

	if request.Login != nil {
		userData.Login = *request.Login
	}

	if request.RoleID != nil {
		userData.Role = model.Role{ID: *request.RoleID}
	}

	return userData
}
