package handlers

import (
	"context"
	"testrunsystem/gen/restapi/operations/product"
	"testrunsystem/internal/repositories"
	"testrunsystem/runtime"
)

type (
	Handler interface {
		productHandler
	}

	productHandler interface {
		CreateProduct(ctx context.Context, req *product.CreateProductParams) (*uint64, error)
		UpdateProduct(ctx context.Context, req *product.UpdateProductParams) error
		DeleteProduct(ctx context.Context, req *product.DeleteProductParams) error
		UpdateProductStock(ctx context.Context, req *product.UpdateProductStockParams) error
	}
)

func NewHandler(
	rt runtime.Runtime,
	productRepo repositories.Product,
	productActivityHistoryRepo repositories.ProductActivityHistory,
) Handler {
	return &handler{
		rt,
		productRepo,
		productActivityHistoryRepo,
	}
}

type handler struct {
	runtime                    runtime.Runtime
	productRepo                repositories.Product
	productActivityHistoryRepo repositories.ProductActivityHistory
}
