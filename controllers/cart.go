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

type ResponseData struct {
	RespCode string `json:"response_code"`
	RespDesc string `json:"response_description"`
	Data     CustomMessage
}

type CustomMessage struct {
	CustomMessage string `json:"custom_message"`
}

func NewCartController() *CartController {
	return &CartController{}
}

// CreateCart creates a new Cart resource
func (cc CartController) CreateCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an order_details to be populated from the body
	u := models.OrderDetails{}

	// Populate the order_details data
	json.NewDecoder(r.Body).Decode(&u)

	//construct function for tax calc
	x := tax{}
	x.amount = u.Amount
	x.taxCode = u.TaxCode
	myTax := setTaxAmount(x)

	//Insert OrderDetails : amount, tax_amount, total_amount
	u.TaxAmount = myTax.taxAmount
	u.TotalAmount = myTax.taxAmount + u.Amount
	//Query insert into model
	//....

	//define response
	d := ResponseData{"1", "success", CustomMessage{"Cart Created"}}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

//setTaxAmount to set tax code type struct for calcTaxAmount
func setTaxAmount(t tax) tax {
	result := tax{}
	result.amount = t.amount
	if t.taxCode == 1 {
		aFood := food{}
		aFood.amount = t.amount
		result.taxAmount = calcTaxAmount(aFood)
	} else if t.taxCode == 2 {
		aTobbaco := tobacco{}
		aTobbaco.amount = t.amount
		result.taxAmount = calcTaxAmount(aTobbaco)
	} else if t.taxCode == 3 {
		aEntertainment := entertainment{}
		aEntertainment.amount = t.amount
		result.taxAmount = calcTaxAmount(aEntertainment)
	} else {
		result.taxAmount = 0
	}
	return result
}

func calcTaxAmount(z calculator) float64 {
	return z.getTaxAmount()
}

type tax struct {
	storeId     int
	productName string
	taxCode     int
	amount      float64
	taxAmount   float64
}

type food struct {
	tax
}

type tobacco struct {
	tax
}

type entertainment struct {
	tax
}

type calculator interface {
	getTaxAmount() float64
}

//polymorphism
//food: 10% of value
func (f food) getTaxAmount() float64 {
	return f.amount * 0.1
}

//tobacco: 10 + (2% of value)
func (t tobacco) getTaxAmount() float64 {
	return 10 + (0.02 * t.amount)
}

//entertainment:0<value<100 ? tax-free : 1% of (value - 100)
func (e entertainment) getTaxAmount() float64 {
	if e.amount >= 100 {
		return (e.amount - 100) * 0.01
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
