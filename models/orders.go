package models

type Orders struct {
	StoreId        int     `json:"store_id"`
	StoreName      string  `json:"store_name"`
	TotalAmount    float64 `json:"total_amount"`
	TotalTaxAmount float64 `json:"total_tax_amount"`
	GrandTotal     float64 `json:"grand_total"`
}
