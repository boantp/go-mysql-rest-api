package models

import (
	"github.com/boantp/go-mysql-rest-api/config"
)

type TaxCode struct {
	TaxCodeId int    `json:"tax_code_id"`
	Name      string `json:"name"`
}

func FetchTaxCode() []TaxCode {
	rows, err := config.DB.Query("SELECT tax_code_id, name FROM tax_code")
	checkErr(err)

	defer rows.Close()

	taxes := make([]TaxCode, 0)
	for rows.Next() {
		tax := TaxCode{}
		err := rows.Scan(&tax.TaxCodeId, &tax.Name)
		checkErr(err)

		taxes = append(taxes, tax)
	}
	return taxes
}
