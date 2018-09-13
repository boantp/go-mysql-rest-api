package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	// Standard library packages

	"github.com/boantp/go-mysql-rest-api/config"
	"github.com/boantp/go-mysql-rest-api/controllers"
	"github.com/boantp/go-mysql-rest-api/models"
	"github.com/julienschmidt/httprouter"
	// Third party packages
)

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

	//Frontend
	router.GET("/", Index)
	router.GET("/order_view/:store_id", ViewBill)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//get tax
	url := "http://localhost:3000/tax_code"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "tax-calculator")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	type TaxCodeResp struct {
		RespCode string           `json:"response_code"`
		RespDesc string           `json:"response_description"`
		Data     []models.TaxCode `json:"data"`
	}

	type TaxCode struct {
		TaxCodeId int    `json:"tax_code_id"`
		Name      string `json:"name"`
	}

	tax := TaxCodeResp{}
	jsonErr := json.Unmarshal(body, &tax)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	taxData := tax.Data
	config.TPL.ExecuteTemplate(w, "create.gohtml", taxData)
}

func ViewBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
