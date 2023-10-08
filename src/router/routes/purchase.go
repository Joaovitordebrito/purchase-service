package routes

import (
	"net/http"

	"github.com/Joaovitordebrito/wex-purchase-service/src/controller"
)

var userRoutes = []Route{
	{
		URI:      "/purchase",
		Method:   http.MethodPost,
		Function: controller.CreatePurchase,
		Auth:     false,
	},
	{
		URI:      "/converted/currency/{uuid}/{country}",
		Method:   http.MethodGet,
		Function: controller.GetConvertedCurrency,
		Auth:     false,
	},
}
