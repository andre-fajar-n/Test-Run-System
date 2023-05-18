package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testrunsystem/gen/models"
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

	var action string
	switch historyType {
	case "create":
		newByte, _ := json.Marshal(newData)
		action = fmt.Sprintf("create new data: %v", string(newByte))
	case "update":
		newByte, _ := json.Marshal(newData)
		oldByte, _ := json.Marshal(oldData)
		action = fmt.Sprintf("update data from: %v to %v", string(newByte), string(oldByte))
	default:
		return h.runtime.SetError(http.StatusBadRequest, "invalid product history type: %s", historyType)
	}

	now := time.Now().UTC()
	err := h.productActivityHistoryRepo.Create(ctx, tx, &models.ProductActivityHistory{
		Action:    action,
		CreatedAt: strfmt.DateTime(now),
		ProductID: *newData.ID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("error productActivityHistoryRepo.Create")
		return err
	}

	return nil
}
