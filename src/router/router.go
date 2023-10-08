package router

import (
	"github.com/Joaovitordebrito/wex-purchase-service/src/router/routes"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
