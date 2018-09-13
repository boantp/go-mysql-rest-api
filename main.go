package main

import (
	"fmt"
	"log"
	"net/http"
	// Standard library packages

	"github.com/boantp/go-mysql-rest-api/controllers"

	"github.com/julienschmidt/httprouter"
	// Third party packages
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()

	// Get Controller instance
	tc := controllers.NewTaxCodeController()
	cc := controllers.NewCartController()
	oc := controllers.NewOrderController()

	// Get a tax code resource
	router.GET("/tax_code", tc.GetTaxCode)
	//POST object tax into cart or order_details
	router.POST("/cart", cc.CreateCart)
	//GET my bill with tax calculation result from order_details
	router.GET("/order/:store_id", oc.GetMyBill)

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":3000", router))
}
