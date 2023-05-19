package repositories

import (
	"context"
	"fmt"
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
		FindPaginate(ctx context.Context, userID, productID uint64, limit, offset int64, sort, order string) ([]models.ProductActivityHistory, *int64, error)
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

func (r *productActivityHistory) FindPaginate(
	ctx context.Context,
	userID,
	productID uint64,
	limit,
	offset int64,
	sort,
	order string,
) ([]models.ProductActivityHistory, *int64, error) {
	logger := r.runtime.Logger.With().
		Logger()

	var data []models.ProductActivityHistory
	var totalRow int64
	db := r.runtime.Db.Model(&data).Where("product_id", productID)
	err := db.Count(&totalRow).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query count")
		return nil, nil, err
	}

	err = db.Limit(int(limit)).Offset(int(offset)).Order(fmt.Sprintf("%s %s", order, sort)).Find(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query get")
		return nil, nil, err
	}

	return data, &totalRow, nil
}
