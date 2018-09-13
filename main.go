package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	router.POST("/front/cart/process", FrontCartProcess)
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

	tax := TaxCodeResp{}
	jsonErr := json.Unmarshal(body, &tax)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	taxData := tax.Data
	config.TPL.ExecuteTemplate(w, "create.gohtml", taxData)
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

type Tax struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount_item"`
}

func FrontCartProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get form values
	addTax := Tax{}
	s := r.FormValue("store_id")
	a := r.FormValue("amount")
	t := r.FormValue("tax_code_id")
	f64, err := strconv.ParseFloat(a, 32)
	i, err := strconv.Atoi(s)
	//j, err := strconv.Atoi(a)
	k, err := strconv.Atoi(t)
	addTax.StoreId = i
	addTax.ProductName = r.FormValue("product_name")
	addTax.Amount = float64(f64)
	addTax.TaxCode = k
	uj, _ := json.Marshal(addTax)

	//POST localhost:3000/cart
	url := "http://localhost:3000/cart"
	//fmt.Println("URL:>", url)

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(uj))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	config.TPL.ExecuteTemplate(w, "created.gohtml", resp)
}

func ViewBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
