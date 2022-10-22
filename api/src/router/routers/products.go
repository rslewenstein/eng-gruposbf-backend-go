package routers

import (
	"net/http"

	"go.mod/src/controllers"
)

var productRouters = []Router{
	{
		URI:      "/api/convert",
		Method:   http.MethodGet,
		Function: controllers.GetConverterAllCurrency,
	},
	{
		URI:      "/api/convert/{productId}&{symbol}&{amount}",
		Method:   http.MethodGet,
		Function: controllers.GetConverterAllCurrency,
	},
	// {
	// 	URI: "api/product/currency",
	// 	Method: http.MethodPost,
	// 	Function: controllers.CreateCurrency;
	// },

}
