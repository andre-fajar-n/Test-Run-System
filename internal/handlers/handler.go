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
	}
)

func NewHandler(
	rt runtime.Runtime,
	productRepo repositories.Product,
) Handler {
	return &handler{
		rt,
		productRepo,
	}
}

type handler struct {
	runtime     runtime.Runtime
	productRepo repositories.Product
}
