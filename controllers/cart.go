package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boantp/go-mysql-rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// CartController represents the controller for create cart and calculate tax
	CartController struct{}
)

func NewCartController() *CartController {
	return &CartController{}
}

type CartResp struct {
	RespCode string `json:"response_code"`
	RespDesc string `json:"response_description"`
	Data     CustomMessage
}

type CustomMessage struct {
	CustomMessage string `json:"custom_message"`
}

// CreateCart creates a new Cart resource
func (cc CartController) CreateCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an order_details to be populated from the body
	u := models.Tax{}

	// Populate the order_details data
	json.NewDecoder(r.Body).Decode(&u)

	//construct function for tax calc
	myTax := setTaxAmount(u)

	//Inject taxamount and totalamount into object
	u.TaxAmount = myTax.TaxAmount
	u.TotalAmount = myTax.TaxAmount + u.Amount
	u.TaxName = myTax.TaxName

	//Query insert into model
	_, err := models.CreateOrderDetails(u)
	if err != nil {
		fmt.Println(err.Error())
	}

	//define response
	d := CartResp{"1", "success", CustomMessage{"Cart Created"}}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

type food struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount_item"`
}

type tobacco struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount_item"`
}

type entertainment struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount_item"`
}

type calculator interface {
	getTaxAmount() float64
}

//setTaxAmount to set tax code type struct for calcTaxAmount
func setTaxAmount(t models.Tax) models.Tax {
	result := models.Tax{}
	result.Amount = t.Amount
	if t.TaxCode == 1 {
		aFood := food{}
		aFood.Amount = t.Amount
		result.TaxAmount = calcTaxAmount(aFood)
		result.TaxName = "Food"
	} else if t.TaxCode == 2 {
		aTobbaco := tobacco{}
		aTobbaco.Amount = t.Amount
		result.TaxAmount = calcTaxAmount(aTobbaco)
		result.TaxName = "Tobacco"
	} else if t.TaxCode == 3 {
		aEntertainment := entertainment{}
		aEntertainment.Amount = t.Amount
		result.TaxAmount = calcTaxAmount(aEntertainment)
		result.TaxName = "Entertainment"
	} else {
		result.TaxAmount = 0
	}
	return result
}

func calcTaxAmount(z calculator) float64 {
	return z.getTaxAmount()
}

//polymorphism
//food: 10% of value
func (f food) getTaxAmount() float64 {
	return f.Amount * 0.1
}

//tobacco: 10 + (2% of value)
func (t tobacco) getTaxAmount() float64 {
	return 10 + (0.02 * t.Amount)
}

//entertainment:0<value<100 ? tax-free : 1% of (value - 100)
func (e entertainment) getTaxAmount() float64 {
	if e.Amount >= 100 {
		return (e.Amount - 100) * 0.01
	}
	return 0
}

// a new method which takes the INTERFACE type
func totalTaxAmount(calculators ...calculator) float64 {
	var totalTaxAmount float64
	for _, c := range calculators {
		totalTaxAmount += c.getTaxAmount()
	}

	return totalTaxAmount
}
