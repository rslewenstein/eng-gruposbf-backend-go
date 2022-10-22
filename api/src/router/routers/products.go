package routers

import (
	"net/http"

	"go.mod/src/controllers"
)

var productRouters = []Router{
	{
		URI:      "/api/convert/{symbol}&{amount}", // "/api/convert/{productId}&{symbol}&{amount}",
		Method:   http.MethodGet,
		Function: controllers.GetConvertedCurrency,
	},
	{
		URI: "/api/convert/currency",
		Method: http.MethodPost,
		Function: controllers.CreateCurrency,
	},
}
