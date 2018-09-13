package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/boantp/go-mysql-rest-api/config"
	"github.com/boantp/go-mysql-rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// TaxCodeController represents the controller for operating on the Tax Code resource
	WebController struct{}
)

func NewWebController() *WebController {
	return &WebController{}
}

//custom var
type BaseUrl string

const (
	development BaseUrl = "http://localhost:3000/"
	staging     BaseUrl = "http://localhost:4000/"
	production  BaseUrl = "http://localhost:5000/"
)

func (web WebController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//get tax code
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("tax_code")
	url := buffer.String()
	body := RequestGet(url)

	tax := TaxCodeResp{}
	jsonErr := json.Unmarshal(body, &tax)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	taxData := tax.Data

	config.TPL.ExecuteTemplate(w, "create.gohtml", taxData)
}

func (web WebController) FrontCartProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get form values
	formStoreId := r.FormValue("store_id")
	formAmount := r.FormValue("amount")
	formTaxCodeId := r.FormValue("tax_code_id")
	formProductName := r.FormValue("product_name")
	// string to int
	storeId, err := strconv.Atoi(formStoreId)
	if err != nil {
		log.Fatal(err)
	}
	// string to int
	taxCodeId, err := strconv.Atoi(formTaxCodeId)
	if err != nil {
		log.Fatal(err)
	}
	// convert form values
	f64, err := strconv.ParseFloat(formAmount, 32)
	if err != nil {
		log.Fatal(err)
	}
	//store new tax object
	addTax := models.Tax{}
	addTax.StoreId = storeId
	addTax.ProductName = formProductName
	addTax.TaxCode = taxCodeId
	addTax.Amount = float64(f64)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(addTax)

	//POST localhost:3000/cart
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("cart")
	url := buffer.String()

	//make post request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(uj))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	data := struct {
		RespStatus string
		RespBody   string
	}{
		resp.Status,
		string(body),
	}

	config.TPL.ExecuteTemplate(w, "created.gohtml", data)
}

func (web WebController) ViewBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//get order by store_id
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("order/1")
	url := buffer.String()

	body := RequestGet(url)

	order := OrderResp{}
	jsonErr := json.Unmarshal(body, &order)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	config.TPL.ExecuteTemplate(w, "orders.gohtml", order)
}

func RequestGet(url string) []byte {
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

	return body
}
