package rest

import (
	"testrunsystem/gen/restapi/operations"
	"testrunsystem/gen/restapi/operations/health"
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

}
