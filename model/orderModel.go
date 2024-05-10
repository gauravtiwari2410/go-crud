package model

import "time"

type Order struct {
	Id            int `json:"id"`
	CashierID     int `json:"cashierID"`
	PaymentTypeId int `json:"paymentTypeId"`

	TotalPrice  int       `json:"totalPrice"`
	TotalPaid   int       `json:"totalPaid"`
	TotalReturn int       `json:"totalReturn"`
	ReceiptId   string    `json:"receiptId"`
	IsDownload  int       `json:"is_download"`
	ProductId   string    `json:"productId"`
	Quantities  string    `json:"quantities"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
