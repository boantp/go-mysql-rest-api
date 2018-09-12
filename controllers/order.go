package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/boantp/go-mysql-rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// OrderController represents the controller for create order and display calculate tax
	OrderController struct{}
)

type OrderResp struct {
	RespCode string       `json:"response_code"`
	RespDesc string       `json:"response_description"`
	Data     []models.Tax `json:"data"`
	models.Orders
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

// GetMyBill retrieves an order_details resource
func (oc OrderController) GetMyBill(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab store_id
	i := p.ByName("store_id")
	store_id, _ := strconv.Atoi(i)

	//Get from order_details with that store_id that order_status = 0
	bill := models.FetchOrderDetailsByStoreIdForDraftOrder(store_id)

	//Get all total from order_details with that store_id that order_status = 0
	totalBill := models.TotalBillByStoreIdForDraftOrder(store_id)

	//define response
	d := OrderResp{"1", "success", bill, totalBill}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
