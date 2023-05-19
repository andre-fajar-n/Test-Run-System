package rest

import (
	"context"
	"testrunsystem/gen/models"
	"testrunsystem/gen/restapi/operations"
	"testrunsystem/gen/restapi/operations/authentication"
	"testrunsystem/gen/restapi/operations/health"
	"testrunsystem/gen/restapi/operations/product"
	"testrunsystem/internal/handlers"
	"testrunsystem/runtime"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *runtime.Runtime, api *operations.ServerAPI, apiHandler handlers.Handler) {
	//  health
	api.HealthHealthHandler = health.HealthHandlerFunc(func(hp health.HealthParams) middleware.Responder {
		return health.NewHealthOK().WithPayload(&health.HealthOKBody{
			Message: "Server up",
		})
	})

	// auth
	{
		api.AuthenticationRegisterHandler = authentication.RegisterHandlerFunc(func(rp authentication.RegisterParams) middleware.Responder {
			userID, err := apiHandler.Register(context.Background(), rp)
			if err != nil {
				errRes := rt.GetError(err)
				return authentication.NewRegisterDefault(int(errRes.Code())).WithPayload(&authentication.RegisterDefaultBody{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return authentication.NewRegisterCreated().WithPayload(&authentication.RegisterCreatedBody{
				Message: "success register",
				UserID:  *userID,
			})
		})

		api.AuthenticationLoginHandler = authentication.LoginHandlerFunc(func(lp authentication.LoginParams) middleware.Responder {
			token, expiredAt, err := apiHandler.Login(context.Background(), &lp)
			if err != nil {
				errRes := rt.GetError(err)
				return authentication.NewLoginDefault(int(errRes.Code())).WithPayload(&authentication.LoginDefaultBody{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return authentication.NewLoginCreated().WithPayload(&authentication.LoginCreatedBody{
				Message:   "success login",
				ExpiredAt: *expiredAt,
			}).WithToken(*token)
		})
	}

	// product
	{
		api.ProductCreateProductHandler = product.CreateProductHandlerFunc(func(cpp product.CreateProductParams, p *models.Principal) middleware.Responder {
			productID, err := apiHandler.CreateProduct(context.Background(), &cpp, p)
			if err != nil {
				errRes := rt.GetError(err)
				return product.NewCreateProductDefault(int(errRes.Code())).WithPayload(&product.CreateProductDefaultBody{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return product.NewCreateProductCreated().WithPayload(&product.CreateProductCreatedBody{
				Message:   "success create product",
				ProductID: *productID,
			})
		})

		api.ProductUpdateProductHandler = product.UpdateProductHandlerFunc(func(upp product.UpdateProductParams, p *models.Principal) middleware.Responder {
			err := apiHandler.UpdateProduct(context.Background(), &upp, p)
			if err != nil {
				errRes := rt.GetError(err)
				return product.NewUpdateProductDefault(int(errRes.Code())).WithPayload(&product.UpdateProductDefaultBody{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return product.NewUpdateProductCreated().WithPayload(&product.UpdateProductCreatedBody{
				Message: "success update product",
			})
		})

		api.ProductDeleteProductHandler = product.DeleteProductHandlerFunc(func(dpp product.DeleteProductParams, p *models.Principal) middleware.Responder {
			err := apiHandler.DeleteProduct(context.Background(), &dpp, p)
			if err != nil {
				errRes := rt.GetError(err)
				return product.NewDeleteProductDefault(int(errRes.Code())).WithPayload(&product.DeleteProductDefaultBody{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return product.NewDeleteProductCreated().WithPayload(&product.DeleteProductCreatedBody{
				Message: "success delete product",
			})
		})

		api.ProductUpdateProductStockHandler = product.UpdateProductStockHandlerFunc(func(upsp product.UpdateProductStockParams, p *models.Principal) middleware.Responder {
			err := apiHandler.UpdateProductStock(context.Background(), &upsp, p)
			if err != nil {
				errRes := rt.GetError(err)
				return product.NewUpdateProductStockDefault(int(errRes.Code())).WithPayload(&product.UpdateProductStockDefaultBody{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return product.NewUpdateProductStockCreated().WithPayload(&product.UpdateProductStockCreatedBody{
				Message: "success update product stock",
			})
		})
	}
}
