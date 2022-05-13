package qiwi

import "net/http"

// Paymenter - struct containing information needed
// by all functions making requests to the qiwi api
type Paymenter struct {
	PrivateKey string
	Client     *http.Client
}

// NewPaymenter - creates a new Paymenter struct,
// returns its pointer
func NewPaymenter(privateKey string) *Paymenter {
	return &Paymenter{
		PrivateKey: privateKey,
		Client:     &http.Client{},
	}
}
