package qiwi

import (
	"strconv"
	"time"
)

// Bill - struct representation of the json needed to be sent
// to the qiwi api, to create a bill in the qiwi system
type Bill struct {
	Amount             Amount   `json:"amount"`
	Comment            string   `json:"comment"`
	ExpirationDateTime QiwiTime `json:"expirationDateTime"`
}

// Amount - a sub part of the bill struct
type Amount struct {
	Currency Currency `json:"currency"`
	Value    string   `json:"value"`
}

// NewBill - creates a new bill struct
func NewBill(
	sum int,
	currency Currency,
	expirationTime time.Time,
	comment string,
) Bill {

	return Bill{
		Amount: Amount{
			Currency: currency,
			Value:    strconv.Itoa(sum) + ".00",
		},
		Comment:            comment,
		ExpirationDateTime: ConvertToQiwiTime(expirationTime),
	}
}
