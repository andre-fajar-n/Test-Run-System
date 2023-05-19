package handlers

import (
	"context"
	"fmt"
	"net/http"
	"testrunsystem/gen/models"
	"testrunsystem/gen/restapi/operations/product"
	"time"

	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func (h *handler) CreateProduct(ctx context.Context, req *product.CreateProductParams, p *models.Principal) (*uint64, error) {
	logger := h.runtime.Logger.With().
		Interface("data", req.Data).
		Logger()

	err := h.checkSameName(ctx, *req.Data.Name, nil, &p.UserID)
	if err != nil {
		logger.Error().Err(err).Msg("error checkSameName")
		return nil, err
	}

	var expiryDate *strfmt.DateTime
	if req.Data.ExpiryDate != "" {
		temp, _ := time.Parse("01-02-2006", req.Data.ExpiryDate)
		tempExpDate := strfmt.DateTime(temp)
		expiryDate = &tempExpDate
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	data := &models.Product{
		Name:       *req.Data.Name,
		Stock:      *req.Data.Stock,
		ExpiryDate: expiryDate,
		CreatedAt:  &nowStrfmt,
		CreatedBy:  fmt.Sprintf("%d", p.UserID),
		UserID:     p.UserID,
	}

	err = h.runtime.Db.Transaction(func(tx *gorm.DB) error {
		data, err = h.productRepo.Create(ctx, tx, data)
		if err != nil {
			logger.Error().Err(err).Msg("error productRepo.Create")
			return err
		}

		err = h.createProductActivityHistory(ctx, tx, "create", data, nil)
		if err != nil {
			logger.Error().Err(err).Msg("error createProductActivityHistory")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("error DB Transaction")
		return nil, err
	}

	return data.ID, nil
}

func (h *handler) checkSameName(ctx context.Context, name string, productID, userID *uint64) error {
	logger := h.runtime.Logger.With().
		Str("name", name).
		Logger()

	isNameExist, err := h.productRepo.NameExist(ctx, name, productID, userID)
	if err != nil {
		logger.Error().Err(err).Msg("error productRepo.NameExist")
		return err
	}
	if isNameExist {
		return h.runtime.SetError(http.StatusNotAcceptable, "product name already exist")
	}

	return nil
}

func (h *handler) UpdateProduct(ctx context.Context, req *product.UpdateProductParams, p *models.Principal) error {
	logger := h.runtime.Logger.With().
		Interface("data", req.Data).
		Logger()

	existingData, err := h.productRepo.FindBySingleColumn(ctx, map[string]interface{}{
		"id":      req.ProductID,
		"user_id": p.UserID,
	}, false)
	if err != nil {
		logger.Error().Err(err).Msg("error productRepo.FindBySingleColumn")
		return err
	}

	err = h.checkSameName(ctx, *req.Data.Name, &req.ProductID, &p.UserID)
	if err != nil {
		logger.Error().Err(err).Msg("error checkSameName")
		return err
	}

	var expiryDate *strfmt.DateTime = nil
	if req.Data.ExpiryDate != "" {
		temp, _ := time.Parse("01-02-2006", req.Data.ExpiryDate)
		tempExpDate := strfmt.DateTime(temp)
		expiryDate = &tempExpDate
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	var data models.Product = *existingData

	data.Name = *req.Data.Name
	data.ExpiryDate = expiryDate
	data.UpdatedAt = &nowStrfmt
	data.UpdatedBy = fmt.Sprintf("%d", p.UserID)

	err = h.runtime.Db.Transaction(func(tx *gorm.DB) error {
		err = h.productRepo.Update(ctx, tx, &data)
		if err != nil {
			logger.Error().Err(err).Msg("error productRepo.Update")
			return err
		}

		err = h.createProductActivityHistory(ctx, tx, "update", &data, existingData)
		if err != nil {
			logger.Error().Err(err).Msg("error createProductActivityHistory")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("error DB Transaction")
		return err
	}

	return nil
}

func (h *handler) DeleteProduct(ctx context.Context, req *product.DeleteProductParams, p *models.Principal) error {
	logger := h.runtime.Logger.With().
		Uint64("data", req.ProductID).
		Logger()

	existingData, err := h.productRepo.FindBySingleColumn(ctx, map[string]interface{}{
		"id":      req.ProductID,
		"user_id": p.UserID,
	}, false)
	if err != nil {
		logger.Error().Err(err).Msg("error productRepo.FindBySingleColumn")
		return err
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	existingData.DeletedAt = &nowStrfmt
	existingData.DeletedBy = fmt.Sprintf("%d", p.UserID)
	err = h.runtime.Db.Transaction(func(tx *gorm.DB) error {
		err = h.productRepo.Delete(ctx, tx, existingData)
		if err != nil {
			logger.Error().Err(err).Msg("error productRepo.Delete")
			return err
		}

		err = h.createProductActivityHistory(ctx, tx, "delete", existingData, nil)
		if err != nil {
			logger.Error().Err(err).Msg("error createProductActivityHistory")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("error DB Transaction")
		return err
	}

	return nil
}

func (h *handler) UpdateProductStock(ctx context.Context, req *product.UpdateProductStockParams, p *models.Principal) error {
	logger := h.runtime.Logger.With().
		Interface("data", req.Data).
		Logger()

	existingData, err := h.productRepo.FindBySingleColumn(ctx, map[string]interface{}{
		"id":      req.ProductID,
		"user_id": p.UserID,
	}, false)
	if err != nil {
		logger.Error().Err(err).Msg("error productRepo.FindBySingleColumn")
		return err
	}

	now := time.Now().UTC()

	if existingData.ExpiryDate != nil {
		if time.Time(*existingData.ExpiryDate).Before(now) {
			return h.runtime.SetError(http.StatusBadRequest, "this product was expired")
		}
	}

	nowStrfmt := strfmt.DateTime(now)

	var data models.Product = *existingData

	data.UpdatedAt = &nowStrfmt
	data.UpdatedBy = fmt.Sprintf("%d", p.UserID)
	data.Stock = *req.Data.Stock

	err = h.runtime.Db.Transaction(func(tx *gorm.DB) error {
		err = h.productRepo.Update(ctx, tx, &data)
		if err != nil {
			logger.Error().Err(err).Msg("error productRepo.Update")
			return err
		}

		err = h.createProductActivityHistory(ctx, tx, "update_stock", &data, existingData)
		if err != nil {
			logger.Error().Err(err).Msg("error createProductActivityHistory")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("error DB Transaction")
		return err
	}

	return nil
}
