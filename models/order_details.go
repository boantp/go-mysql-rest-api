package models

import (
	"github.com/boantp/go-mysql-rest-api/config"
)

type OrderDetails struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount"`
	OrderStatus int     `json:"order_status"`
}

type TotalOrder struct {
	StoreId        int     `json:"store_id"`
	TotalTaxAmount float64 `json:"total_tax_amount"`
	TotalAmount    float64 `json:"total_amount"`
	GrandTotal     float64 `json:"grand_total"`
}

type DataOrderDetail struct {
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount_item"`
}

func CreateOrderDetails(od OrderDetails) (OrderDetails, error) {
	// insert values
	sqlStr := "INSERT INTO order_details(orders_id, product_name, tax_code_id, tax_name, amount, tax_amount, total_amount, order_status, store_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	//prepare the statement
	stmt, err := config.DB.Prepare(sqlStr)
	checkErr(err)
	//format all vals at once
	_, err = stmt.Exec(0, od.ProductName, od.TaxCode, od.TaxName, od.Amount, od.TaxAmount, od.TotalAmount, od.OrderStatus, od.StoreId)
	checkErr(err)

	return od, nil
}

func FetchOrderDetailsByStoreIdForDraftOrder(storeId int) []DataOrderDetail {
	rows, err := config.DB.Query("SELECT product_name, tax_code_id, tax_name, amount, tax_amount, total_amount FROM order_details WHERE order_status=0 and store_id=?", storeId)
	checkErr(err)

	defer rows.Close()

	ods := make([]DataOrderDetail, 0)
	for rows.Next() {
		od := DataOrderDetail{}
		err := rows.Scan(&od.ProductName, &od.TaxCode, &od.TaxName, &od.Amount, &od.TaxAmount, &od.TotalAmount)
		checkErr(err)

		ods = append(ods, od)
	}
	return ods
}

func TotalBillByStoreIdForDraftOrder(storeId int) TotalOrder {
	total := TotalOrder{}

	row := config.DB.QueryRow("SELECT store_id, SUM(amount) as total_amount, SUM(tax_amount) as total_tax_amount, SUM(total_amount) as grand_total FROM order_details where order_status = 0 and store_id = ?", storeId)

	err := row.Scan(&total.StoreId, &total.TotalAmount, &total.TotalTaxAmount, &total.GrandTotal)
	checkErr(err)

	return total
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
