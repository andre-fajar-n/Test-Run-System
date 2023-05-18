package handlers

import (
	"context"
	"net/http"
	"testrunsystem/gen/models"
	"testrunsystem/gen/restapi/operations/product"
	"time"

	"github.com/go-openapi/strfmt"
)

func (h *handler) CreateProduct(ctx context.Context, req *product.CreateProductParams) (*uint64, error) {
	logger := h.runtime.Logger.With().
		Interface("data", req.Data).
		Logger()

	isNameExist, err := h.productRepo.NameExist(ctx, *req.Data.Name)
	if err != nil {
		logger.Error().Err(err).Msg("error productRepo.NameExist")
		return nil, err
	}
	if isNameExist {
		return nil, h.runtime.SetError(http.StatusNotAcceptable, "product name already exist")
	}

	var expiryDate strfmt.DateTime
	if req.Data.ExpiryDate != "" {
		temp, _ := time.Parse("01-02-2006", req.Data.ExpiryDate)
		expiryDate = strfmt.DateTime(temp)
	}

	now := time.Now().UTC()

	data := &models.Product{
		Name:       *req.Data.Name,
		Stock:      *req.Data.Stock,
		ExpiryDate: expiryDate,
		CreatedAt:  strfmt.DateTime(now),
	}
	data, err = h.productRepo.Create(ctx, data)
	if err != nil {
		logger.Error().Err(err).Msg("error productRepo.Create")
		return nil, err
	}

	return data.ID, nil
}
