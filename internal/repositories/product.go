package repositories

import (
	"context"
	"fmt"
	"net/http"
	"testrunsystem/gen/models"
	"testrunsystem/runtime"

	"gorm.io/gorm"
)

type (
	product struct {
		runtime runtime.Runtime
	}

	Product interface {
		Create(ctx context.Context, tx *gorm.DB, data *models.Product) (*models.Product, error)
		NameExist(ctx context.Context, name string, productID *uint64) (bool, error)
		FindBySingleColumn(ctx context.Context, column string, value interface{}) (*models.Product, error)
		Update(ctx context.Context, tx *gorm.DB, data *models.Product) error
	}
)

func NewProduct(rt runtime.Runtime) Product {
	return &product{
		rt,
	}
}

func (r *product) Create(ctx context.Context, tx *gorm.DB, data *models.Product) (*models.Product, error) {
	logger := r.runtime.Logger.With().
		Interface("data", data).
		Logger()

	err := tx.Model(&data).Create(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return data, nil
}

func (r *product) NameExist(ctx context.Context, name string, productID *uint64) (bool, error) {
	logger := r.runtime.Logger.With().
		Str("name", name).
		Logger()

	productModel := models.Product{}

	db := r.runtime.Db.Model(&productModel).Where("LOWER(name) = LOWER(?)", name)

	if productID != nil {
		db = db.Where("id <> ?", *productID)
	}

	err := db.First(&productModel).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return false, err
	}

	return true, nil
}

func (r *product) FindBySingleColumn(ctx context.Context, column string, value interface{}) (*models.Product, error) {
	logger := r.runtime.Logger.With().
		Str("column", column).
		Interface("value", value).
		Logger()

	productModel := models.Product{}
	err := r.runtime.Db.Model(&productModel).Where(fmt.Sprintf("%s = ?", column), value).First(&productModel).Error
	if err == gorm.ErrRecordNotFound {
		return nil, r.runtime.SetError(http.StatusNotFound, "product not found")
	}
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return &productModel, nil
}

func (r *product) Update(ctx context.Context, tx *gorm.DB, data *models.Product) error {
	logger := r.runtime.Logger.With().
		Interface("data", data).
		Logger()

	err := tx.Model(&data).Updates(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}
