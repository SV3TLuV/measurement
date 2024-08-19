package jobs

import (
	"context"
	"log"
	"measurements-api/internal/service"
)

func NewDeletingOutdatedMeasurementJob(
	measurementService service.MeasurementService) func() {
	return func() {
		ctx := context.Background()
		if err := measurementService.DeleteOutdatedMeasurements(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
