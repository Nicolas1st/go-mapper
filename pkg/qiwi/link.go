package qiwi

var CreateBillLinkBase = "https://api.qiwi.com/partner/bill/v1/bills/"

// NewBillLinks - creates new link
func NewBillLink(id string) string {
	return CreateBillLinkBase + id
}
