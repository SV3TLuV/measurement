package converter

import (
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToColumnFromRepo(column *entities.Column) *model.Column {
	if column == nil {
		return nil
	}

	columnModel := &model.Column{}

	if column.ID.Valid {
		columnModel.ID = column.ID.Int.Uint64()
	}

	if column.Title.Valid {
		columnModel.Title = column.Title.String
	}

	if column.ShortTitle.Valid {
		columnModel.ShortTitle = column.ShortTitle.String
	}

	if column.Formula.Valid {
		columnModel.Formula = &column.Formula.String
	}

	if column.ObjField.Valid {
		columnModel.ObjField = column.ObjField.String
	}

	if column.Code.Valid {
		columnModel.Code = &column.Code.String
	}

	return columnModel
}

func ToColumnsFromRepo(columns []*entities.Column) []*model.Column {
	models := make([]*model.Column, 0, len(columns))
	for i := 0; i < len(columns); i++ {
		models = append(models, ToColumnFromRepo(columns[i]))
	}
	return models
}
