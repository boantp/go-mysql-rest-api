package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/boantp/go-mysql-rest-api/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get Controller instance
	tc := controllers.NewTaxCodeController()
	cc := controllers.NewCartController()
	oc := controllers.NewOrderController()

	// Get a tax code resource
	r.GET("/tax_code", tc.GetTaxCode)
	//POST object tax into cart or order_details
	r.POST("/cart", cc.CreateCart)
	//GET my bill with tax calculation result from order_details
	r.GET("/order/:store_id", oc.GetMyBill)
	//POST my bill into orders and update order_status at order_details
	r.POST("/order/:store_id", oc.CreateMyOrder)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
