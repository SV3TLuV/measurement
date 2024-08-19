package converter

import (
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
	"measurements-api/internal/repository/post_info/converter"
)

func ToObjectFromRepo(object *entities.Object) *model.Object {
	if object == nil {
		return nil
	}

	obj := &model.Object{
		Type:     *ToObjectTypeFromRepo(&object.Type),
		PostInfo: converter.ToPostInfoFromRepo(object.PostInfo),
		Children: ToObjectsFromRepo(object.Children),
	}

	if object.ID.Valid {
		obj.ID = object.ID.Int.Uint64()
	}

	if object.Title.Valid {
		obj.Title = object.Title.String
	}

	if object.Address.Valid {
		obj.Address = &object.Address.String
	}

	if object.Lat.Valid {
		obj.Lat = &object.Lat.Float64
	}

	if object.Lon.Valid {
		obj.Lon = &object.Lon.Float64
	}

	if object.Laboratory.Valid {
		obj.Laboratory = &object.Laboratory.String
	}

	if object.City.Valid {
		obj.City = &object.City.String
	}

	if object.ParentID.Valid {
		parentID := object.ParentID.Int.Uint64()
		obj.ParentID = &parentID
	}

	return obj
}

func ToObjectsFromRepo(objects []*entities.Object) []*model.Object {
	models := make([]*model.Object, 0, len(objects))
	for i := 0; i < len(objects); i++ {
		models = append(models, ToObjectFromRepo(objects[i]))
	}
	return models
}

func ToObjectTypeFromRepo(objectType *entities.ObjectType) *model.ObjectType {
	if objectType == nil {
		return nil
	}

	return &model.ObjectType{
		ID:    objectType.ID,
		Title: objectType.Title,
	}
}

func ToObjectWithOperationFromRepo(object *entities.ObjectWithOperation) *model.Object {
	if object == nil {
		return nil
	}

	obj := &model.Object{
		Type:     *ToObjectTypeFromRepo(&object.Type),
		PostInfo: converter.ToPostInfoFromRepo(object.PostInfo),
		Children: ToObjectsFromRepo(object.Children),
		Status:   &object.Operation,
	}

	if object.ID.Valid {
		obj.ID = object.ID.Int.Uint64()
	}

	if object.Title.Valid {
		obj.Title = object.Title.String
	}

	if object.Address.Valid {
		obj.Address = &object.Address.String
	}

	if object.Lat.Valid {
		obj.Lat = &object.Lat.Float64
	}

	if object.Lon.Valid {
		obj.Lon = &object.Lon.Float64
	}

	if object.Laboratory.Valid {
		obj.Laboratory = &object.Laboratory.String
	}

	if object.City.Valid {
		obj.City = &object.City.String
	}

	if object.ParentID.Valid {
		parentID := object.ParentID.Int.Uint64()
		obj.ParentID = &parentID
	}

	return obj
}

func ToObjectsWithOperationFromRepo(objects []*entities.ObjectWithOperation) []*model.Object {
	models := make([]*model.Object, 0, len(objects))
	for i := 0; i < len(objects); i++ {
		models = append(models, ToObjectWithOperationFromRepo(objects[i]))
	}
	return models
}
