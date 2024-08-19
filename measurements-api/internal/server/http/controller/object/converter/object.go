package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/object/model"
)

func ToObjectViewFromService(object *model.Object) *model2.ObjectView {
	if object == nil {
		return nil
	}

	view := &model2.ObjectView{
		ID:         object.ID,
		Title:      object.Title,
		Address:    object.Address,
		Lat:        object.Lat,
		Lon:        object.Lon,
		Type:       ToObjectTypeViewFromService(&object.Type),
		ParentID:   object.ParentID,
		Children:   ToObjectViewsFromService(object.Children),
		Laboratory: object.Laboratory,
		City:       object.City,
		Status:     object.Status,
	}

	if object.PostInfo != nil {
		view.LastPollingDateTime = object.PostInfo.LastPollingDateTime
		view.IsListened = &object.PostInfo.IsListened
	}

	return view
}

func ToObjectViewsFromService(objects []*model.Object) []*model2.ObjectView {
	if len(objects) == 0 {
		return nil
	}

	views := make([]*model2.ObjectView, 0, len(objects))
	for i := 0; i < len(objects); i++ {
		views = append(views, ToObjectViewFromService(objects[i]))
	}
	return views
}

func ToObjectTypeViewFromService(objType *model.ObjectType) *model2.ObjectTypeView {
	if objType == nil {
		return nil
	}

	return &model2.ObjectTypeView{
		ID:    objType.ID,
		Title: objType.Title,
	}
}
