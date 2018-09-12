package models

import (
	"github.com/boantp/go-mysql-rest-api/config"
)

type Tax struct {
	StoreId     int     `json:"store_id"`
	ProductName string  `json:"product_name"`
	TaxCode     int     `json:"tax_code"`
	TaxName     string  `json:"tax_name"`
	Amount      float64 `json:"amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount_item"`
}

func CreateOrderDetails(od Tax) (Tax, error) {
	// insert values
	sqlStr := "INSERT INTO order_details(orders_id, product_name, tax_code_id, tax_name, amount, tax_amount, total_amount, store_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	//prepare the statement
	stmt, err := config.DB.Prepare(sqlStr)
	checkErr(err)
	//format all vals at once
	_, err = stmt.Exec(0, od.ProductName, od.TaxCode, od.TaxName, od.Amount, od.TaxAmount, od.TotalAmount, od.StoreId)
	checkErr(err)

	return od, nil
}

func FetchOrderDetailsByStoreIdForDraftOrder(storeId int) []Tax {
	rows, err := config.DB.Query("SELECT store_id, product_name, tax_code_id, tax_name, amount, tax_amount, total_amount FROM order_details WHERE order_status=0 and store_id=?", storeId)
	checkErr(err)

	defer rows.Close()

	ods := make([]Tax, 0)
	for rows.Next() {
		od := Tax{}
		err := rows.Scan(&od.StoreId, &od.ProductName, &od.TaxCode, &od.TaxName, &od.Amount, &od.TaxAmount, &od.TotalAmount)
		checkErr(err)

		ods = append(ods, od)
	}
	return ods
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
