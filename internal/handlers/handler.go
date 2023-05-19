package handlers

import (
	"context"
	"testrunsystem/gen/models"
	"testrunsystem/gen/restapi/operations/authentication"
	"testrunsystem/gen/restapi/operations/product"
	"testrunsystem/internal/repositories"
	"testrunsystem/runtime"
)

type (
	Handler interface {
		userHandler
		productHandler
		producActivitytHandler
	}

	productHandler interface {
		CreateProduct(ctx context.Context, req *product.CreateProductParams, p *models.Principal) (*uint64, error)
		UpdateProduct(ctx context.Context, req *product.UpdateProductParams, p *models.Principal) error
		DeleteProduct(ctx context.Context, req *product.DeleteProductParams, p *models.Principal) error
		UpdateProductStock(ctx context.Context, req *product.UpdateProductStockParams, p *models.Principal) error
	}

	userHandler interface {
		Register(ctx context.Context, req authentication.RegisterParams) (*uint64, error)
		Login(ctx context.Context, req *authentication.LoginParams) (token, expiredAt *string, err error)
	}

	producActivitytHandler interface {
		FindProductActivityHistoryPaginate(ctx context.Context, req *product.FindActivityHistoryParams, p *models.Principal) (
			[]*product.DataTuple0,
			*product.FindActivityHistoryOKBodyFindActivityHistoryOKBodyAO1Metadata,
			error,
		)
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
