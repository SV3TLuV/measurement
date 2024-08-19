package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/column/model"
)

func ToColumnViewFromService(column *model.Column) *model2.ColumnView {
	if column == nil {
		return nil
	}

	return &model2.ColumnView{
		ID:         column.ID,
		Title:      column.Title,
		ShortTitle: column.ShortTitle,
		Formula:    column.Formula,
		ObjField:   column.ObjField,
		Code:       column.Code,
	}
}

func ToColumnViewsFromService(columns []*model.Column) []*model2.ColumnView {
	views := make([]*model2.ColumnView, 0, len(columns))
	for i := 0; i < len(columns); i++ {
		views = append(views, ToColumnViewFromService(columns[i]))
	}
	return views
}
