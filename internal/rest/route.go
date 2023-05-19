package rest

import (
	"context"
	"testrunsystem/gen/restapi/operations"
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

	// product
	{
		api.ProductCreateProductHandler = product.CreateProductHandlerFunc(func(cpp product.CreateProductParams) middleware.Responder {
			productID, err := apiHandler.CreateProduct(context.Background(), &cpp)
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

		api.ProductUpdateProductHandler = product.UpdateProductHandlerFunc(func(upp product.UpdateProductParams) middleware.Responder {
			err := apiHandler.UpdateProduct(context.Background(), &upp)
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

		api.ProductDeleteProductHandler = product.DeleteProductHandlerFunc(func(dpp product.DeleteProductParams) middleware.Responder {
			err := apiHandler.DeleteProduct(context.Background(), &dpp)
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

		api.ProductUpdateProductStockHandler = product.UpdateProductStockHandlerFunc(func(upsp product.UpdateProductStockParams) middleware.Responder {
			err := apiHandler.UpdateProductStock(context.Background(), &upsp)
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
