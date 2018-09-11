package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boantp/go-mysql-rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// OrderController represents the controller for create order and display calculate tax
	OrderController struct{}
)

type OrderCodeResp struct {
	RespCode       string                `json:"response_code"`
	RespDesc       string                `json:"response_description"`
	Data           []models.OrderDetails `json:"data"`
	StoreId        int                   `json:"store_id"`
	TotalAmount    float64               `json:"total_amount"`
	TotalTaxAmount float64               `json:"total_tax_amount"`
	GrandTotal     float64               `json:"grand_total"`
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

// GetMyBill retrieves an order_details resource
func (oc OrderController) GetMyBill(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab store_id
	//store_id := p.ByName("store_id")

	//Get from order_details with that store_id that order_status = 0
	//...
	// Stub an example user
	u := models.TaxCode{
		TaxCodeId: 1,
		Name:      "food",
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateMyOrder creates a new orders resource
func (oc OrderController) CreateMyOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an order_details to be populated from the body
	u := models.OrderDetails{}

	// Populate the order_details data
	json.NewDecoder(r.Body).Decode(&u)

	// Add Cart Object
	u.StoreId = 1
	u.ProductName = "Foo"
	u.TaxCode = 1
	u.Amount = 10000

	//Insert Orders

	//Update Order Details

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

//calculate total_amount
func calculate_total_amount() {

}

//calculate grand_total
func calculate_grand_total() {

}
