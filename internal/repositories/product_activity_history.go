package repositories

import (
	"context"
	"testrunsystem/gen/models"
	"testrunsystem/runtime"

	"gorm.io/gorm"
)

type (
	productActivityHistory struct {
		runtime runtime.Runtime
	}

	ProductActivityHistory interface {
		Create(ctx context.Context, tx *gorm.DB, data *models.ProductActivityHistory) error
	}
)

func NewProductActivityHistory(rt runtime.Runtime) ProductActivityHistory {
	return &productActivityHistory{
		rt,
	}
}

func (r *productActivityHistory) Create(ctx context.Context, tx *gorm.DB, data *models.ProductActivityHistory) error {
	logger := r.runtime.Logger.With().
		Interface("data", data).
		Logger()

	err := tx.Model(&data).Create(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}
