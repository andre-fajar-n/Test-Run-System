package rest

import (
	"testrunsystem/gen/restapi/operations"

	"github.com/go-openapi/runtime"
)

func Mime(api *operations.ServerAPI) {
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
}
