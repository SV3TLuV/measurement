package converter

import (
	"measurements-api/internal/model"
	"measurements-api/pkg/asoiza"
	"strconv"
)

func ToObjectFromAsoiza(object *asoiza.Node) *model.Object {
	if object == nil {
		return nil
	}

	id, _ := strconv.ParseUint(object.ID, 10, 64)
	objType := ToObjectTypeFromAsoiza(object.ObjectType)

	var parentID *uint64
	switch {
	case object.CityID != nil:
		parentId, _ := strconv.ParseUint(*object.CityID, 10, 64)
		parentID = &parentId
	case object.LabID != nil:
		parentId, _ := strconv.ParseUint(*object.LabID, 10, 64)
		parentID = &parentId
	}

	return &model.Object{
		ID:       id,
		Title:    object.Title,
		Address:  object.Address,
		Lat:      object.Lat,
		Lon:      object.Lon,
		Type:     *objType,
		ParentID: parentID,
	}
}

func ToObjectTypeFromAsoiza(objType string) *model.ObjectType {
	switch objType {
	case "labs":
		return &model.ObjectType{ID: uint64(model.LaboratoryKey)}
	case "sys_localities":
		return &model.ObjectType{ID: uint64(model.CityKey)}
	case "objects":
		return &model.ObjectType{ID: uint64(model.PostKey)}
	default:
		return nil
	}
}
