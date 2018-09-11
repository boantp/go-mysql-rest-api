package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boantp/go-mysql-rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	// TaxCodeController represents the controller for operating on the Tax Code resource
	TaxCodeController struct{}
)

type TaxCodeResp struct {
	RespCode string         `json:"response_code"`
	RespDesc string         `json:"response_description"`
	Data     models.TaxCode `json:"data"`
}

func NewTaxCodeController() *TaxCodeController {
	return &TaxCodeController{}
}

// GetUser retrieves an individual user resource
func (tc TaxCodeController) GetTaxCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Fetch data from model
	//....
	// Stub an example user
	u := models.TaxCode{}
	u.Name = "food"
	u.TaxCodeId = 1

	//define response
	d := TaxCodeResp{"1", "success", u}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
