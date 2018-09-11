package models

import (
	"errors"

	"github.com/boantp/go-mysql-rest-api/config"
)

type OrderDetails struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount"`
	OrderStatus int     `json:"order_status"`
}

func CreateOrderDetails(od OrderDetails) (OrderDetails, error) {
	//insert values
	// insert values
	sqlStr := "INSERT INTO order_details(order_id, product_name, tax_code_id, amount, tax_amount, total_amount, order_status, store_id) VALUES (?, ?, ?, ?, ?)"
	//prepare the statement
	stmt, _ := config.DB.Prepare(sqlStr)

	//format all vals at once
	_, err := stmt.Exec(0, od.ProductName, od.TaxCode, od.Amount, od.TaxAmount, od.TotalAmount, od.OrderStatus, od.StoreId)
	if err != nil {
		return od, errors.New("500. Internal Server Error." + err.Error())
	}
	return od, nil
}
