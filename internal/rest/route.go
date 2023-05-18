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
}
