package repositories

import (
	"context"
	"testrunsystem/gen/models"
	"testrunsystem/runtime"

	"gorm.io/gorm"
)

type (
	product struct {
		runtime runtime.Runtime
	}

	Product interface {
		Create(ctx context.Context, data *models.Product) (*models.Product, error)
		NameExist(ctx context.Context, name string) (bool, error)
	}
)

func NewRepository(rt runtime.Runtime) Product {
	return &product{
		rt,
	}
}

func (r *product) Create(ctx context.Context, data *models.Product) (*models.Product, error) {
	logger := r.runtime.Logger.With().
		Interface("data", data).
		Logger()

	err := r.runtime.Db.Model(&data).Create(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return data, nil
}

func (r *product) NameExist(ctx context.Context, name string) (bool, error) {
	logger := r.runtime.Logger.With().
		Str("name", name).
		Logger()

	productModel := models.Product{}
	err := r.runtime.Db.Model(&productModel).Where("LOWER(name) = LOWER(?)", name).First(&productModel).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return false, err
	}

	return true, nil
}
