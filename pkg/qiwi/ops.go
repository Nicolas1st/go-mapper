package qiwi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// NewBillCreatedResponse - struct represention of reponse
// returned by the qiwi system when making
// a new bill creation request
type NewBillCreatedResponse struct {
	PayUrl string `json:"payUrl"`
}

// CreateNewBill - creates a new bill in the qiwi system,
// return url to the bill form and error
func (p *Paymenter) CreateNewBill(billID string, sum int, currency Currency, expirationTime time.Time, comment string) (string, error) {
	// marshal bill to json
	bill, err := json.Marshal(NewBill(sum, currency, expirationTime, comment))
	if err != nil {
		return "", err
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, NewBillLink(billID), bytes.NewBuffer(bill))
	if err != nil {
		return "", err
	}

	// add Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+p.PrivateKey)

	// perform the request
	response, err := p.Client.Do(req)
	if err != nil {
		return "", err
	}

	// get the url to the payment form from the reponse
	var billUrlReponse NewBillCreatedResponse
	err = json.NewDecoder(response.Body).Decode(&billUrlReponse)

	return billUrlReponse.PayUrl, err
}

// payment statuses returned by the qiwi system
type PaymentStatus string

const (
	PAID    PaymentStatus = "PAID"
	WAITING PaymentStatus = "WAITING"
)

// BillStatusReponse - represents repspone returned by the qiwi system,
// when getting the status of the request
type BillStatusReponse struct {
	Status StatusField `json:"status"`
}

// StatusField - sub part of BillStatusReponse
type StatusField struct {
	Value PaymentStatus `json:"value"`
}

// GetBillStatus - returns status of the bill
func (p *Paymenter) GetBillStatus(billID string) (PaymentStatus, error) {
	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodGet, NewBillLink(billID), nil)
	if err != nil {
		return WAITING, err
	}

	// add Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+p.PrivateKey)

	// perform the request
	response, err := p.Client.Do(req)
	if err != nil {
		return WAITING, err
	}

	// decode the response
	var billStatus BillStatusReponse
	err = json.NewDecoder(response.Body).Decode(&billStatus)
	if err != nil {
		return WAITING, err
	}

	return billStatus.Status.Value, nil
}
