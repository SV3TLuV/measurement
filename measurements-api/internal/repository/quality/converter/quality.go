package converter

import (
	"measurements-api/internal/model"
	"measurements-api/internal/repository/entities"
)

func ToQualityFromRepo(quality *entities.Quality) *model.Quality {
	if quality == nil {
		return nil
	}

	return &model.Quality{
		ID:       quality.ID,
		Priority: quality.Priority,
		Title:    quality.Title,
		Color:    quality.Color,
		Caption:  quality.Caption,
	}
}

func ToQualitiesFromRepo(qualities []*entities.Quality) []*model.Quality {
	models := make([]*model.Quality, 0, len(qualities))
	for i := 0; i < len(qualities); i++ {
		models = append(models, ToQualityFromRepo(qualities[i]))
	}
	return models
}
