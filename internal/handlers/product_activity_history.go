package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testrunsystem/gen/models"
	"testrunsystem/gen/restapi/operations/product"
	"time"

	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func (h *handler) createProductActivityHistory(
	ctx context.Context,
	tx *gorm.DB,
	historyType string,
	newData,
	oldData *models.Product,
) error {
	logger := h.runtime.Logger.With().
		Str("historyType", historyType).
		Interface("newData", newData).
		Interface("oldData", oldData).
		Logger()

	var note string
	switch historyType {
	case "create":
		newByte, _ := json.Marshal(newData)
		note = fmt.Sprintf("create new data: %v", string(newByte))
	case "update":
		newByte, _ := json.Marshal(newData)
		oldByte, _ := json.Marshal(oldData)
		note = fmt.Sprintf("update data from: %v to %v", string(oldByte), string(newByte))
	case "update_stock":
		note = fmt.Sprintf("update stock from: %d to %d", oldData.Stock, newData.Stock)
	case "delete":
		note = "soft delete data"
	default:
		return h.runtime.SetError(http.StatusBadRequest, "invalid product history type: %s", historyType)
	}

	now := time.Now().UTC()
	err := h.productActivityHistoryRepo.Create(ctx, tx, &models.ProductActivityHistory{
		Note:      note,
		Type:      historyType,
		CreatedAt: strfmt.DateTime(now),
		CreatedBy: fmt.Sprintf("%d", newData.UserID),
		ProductID: *newData.ID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("error productActivityHistoryRepo.Create")
		return err
	}

	return nil
}

func (h *handler) FindProductActivityHistoryPaginate(ctx context.Context, req *product.FindActivityHistoryParams, p *models.Principal) (
	[]*product.FindActivityHistoryOKBodyDataItems0,
	*product.FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata,
	error,
) {
	offset := (*req.Page - 1) * *req.Limit
	data, totalRow, err := h.productActivityHistoryRepo.FindPaginate(ctx, p.UserID, req.ProductID, *req.Limit, offset, *req.Sort, *req.Order)
	if err != nil {
		return nil, nil, err
	}

	var output []*product.FindActivityHistoryOKBodyDataItems0
	for _, v := range data {
		output = append(output, &product.FindActivityHistoryOKBodyDataItems0{
			CreatedAt: v.CreatedAt,
			CreatedBy: v.CreatedBy,
			Note:      v.Note,
			ProductID: v.ProductID,
			Type:      v.Type,
			ID:        v.ID,
		})
	}

	metadata := product.FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata{
		Page:      *req.Page,
		PerPage:   *req.Limit,
		TotalPage: *totalRow / (*req.Limit),
		TotalRow:  *totalRow,
	}

	return output, &metadata, nil
}
