package handlers

import (
	"context"
	"testrunsystem/gen/restapi/operations/authentication"
	"testrunsystem/gen/restapi/operations/product"
	"testrunsystem/internal/repositories"
	"testrunsystem/runtime"
)

type (
	Handler interface {
		userHandler
		productHandler
	}

	productHandler interface {
		CreateProduct(ctx context.Context, req *product.CreateProductParams) (*uint64, error)
		UpdateProduct(ctx context.Context, req *product.UpdateProductParams) error
		DeleteProduct(ctx context.Context, req *product.DeleteProductParams) error
		UpdateProductStock(ctx context.Context, req *product.UpdateProductStockParams) error
	}

	userHandler interface {
		Register(ctx context.Context, req authentication.RegisterParams) (*uint64, error)
		Login(ctx context.Context, req *authentication.LoginParams) (token, expiredAt *string, err error)
	}
)

func NewHandler(
	rt runtime.Runtime,
	userRepo repositories.User,
	productRepo repositories.Product,
	productActivityHistoryRepo repositories.ProductActivityHistory,
) Handler {
	return &handler{
		rt,
		userRepo,
		productRepo,
		productActivityHistoryRepo,
	}
}

type handler struct {
	runtime                    runtime.Runtime
	userRepo                   repositories.User
	productRepo                repositories.Product
	productActivityHistoryRepo repositories.ProductActivityHistory
}
