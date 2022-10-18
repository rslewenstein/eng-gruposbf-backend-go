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
	// 	URI: "api/product/",
	// 	Method: http.MethodPost,
	// 	Function: func(w http.ResponseWriter, r *http.Request){
	// 	},
	// },

}
