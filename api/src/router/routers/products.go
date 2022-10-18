package routers

import (
	"net/http"

	"go.mod/src/controllers"
)

var productRouters = []Router{
	{
		URI:      "/api/product/convert-all-currency",
		Method:   http.MethodGet,
		Function: controllers.GetConverterAllCurrency,
	},
	// {
	// 	URI:      "/api/product/convert/{productId}/{currency}",
	// 	Method:   http.MethodGet,
	// 	Function: controllers.GetConverterCurrency,
	// },
	// {
	// 	URI: "api/product/currency",
	// 	Method: http.MethodPost,
	// 	Function: controllers.CreateCurrency;
	// },

}
