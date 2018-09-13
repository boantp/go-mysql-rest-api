package main

import (
	"log"
	"net/http"

	"github.com/boantp/go-mysql-rest-api/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Get Controller instance
	tc := controllers.NewTaxCodeController()
	cc := controllers.NewCartController()
	oc := controllers.NewOrderController()
	web := controllers.NewWebController()

	/** REST API **/
	// Get a tax code resource
	router.GET("/tax_code", tc.GetTaxCode)
	//POST object tax into cart or order_details
	router.POST("/cart", cc.CreateCart)
	//GET my bill with tax calculation result from order_details
	router.GET("/order/:store_id", oc.GetMyBill)

	/** Web Frontend **/
	router.GET("/", web.Index)
	router.POST("/front/cart/process", web.FrontCartProcess)
	router.GET("/order_view/:store_id", web.ViewBill)

	log.Fatal(http.ListenAndServe(":3000", router))
}
