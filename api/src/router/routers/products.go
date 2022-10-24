package routers

import (
	"net/http"

	"go.mod/src/controllers"
)

var productRouters = []Router{
	{
		URI:      "/api/convert/{symbol}&{amount}",
		Method:   http.MethodGet,
		Function: controllers.GetConvertedCurrency,
	},
	{
		URI: "/api/convert/currency",
		Method: http.MethodPost,
		Function: controllers.CreateCurrency,
	},
}
