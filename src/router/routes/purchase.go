package routes

import (
	"net/http"

	"github.com/Joaovitordebrito/wex-purchase-service/src/controller"
)

var purchaseController controller.PurchaseController = &controller.PurchaseControllerImpl{}

var userRoutes = []Route{
	{
		URI:      "/purchase",
		Method:   http.MethodPost,
		Function: purchaseController.CreatePurchase,
		Auth:     false,
	},
	{
		URI:      "/converted/currency/{uuid}/{country}",
		Method:   http.MethodGet,
		Function: purchaseController.GetConvertedCurrency,
		Auth:     false,
	},
}
