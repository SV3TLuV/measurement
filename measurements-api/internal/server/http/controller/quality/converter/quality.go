package converter

import (
	"measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/quality/model"
)

func ToQualityViewFromService(quality *model.Quality) *model2.QualityView {
	if quality == nil {
		return nil
	}

	return &model2.QualityView{
		ID:       quality.ID,
		Priority: quality.Priority,
		Title:    quality.Title,
		Color:    quality.Color,
		Caption:  quality.Caption,
	}
}

func ToQualityViewsFromService(qualities []*model.Quality) []*model2.QualityView {
	views := make([]*model2.QualityView, 0, len(qualities))
	for i := 0; i < len(qualities); i++ {
		views = append(views, ToQualityViewFromService(qualities[i]))
	}
	return views
}
