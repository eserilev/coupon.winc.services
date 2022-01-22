package corporate

import "strconv"

type CorporateOrder struct {
	OrderId        int    `json:"orderId"`
	GiftAmount     int    `json:"giftAmount"`
	ShippingMethod int    `json:"shippingMethod"`
	Email          string `json:"email"`
	UserName       string `json:"userName"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Company        string `json:"company"`
	Street1        string `json:"street1"`
	Street2        string `json:"street2"`
	City           string `json:"city"`
	State          string `json:"state"`
	PostalCode     string `json:"postalCode"`
	Phone          string `json:"phone"`
	GiftMessage    string `json:"giftMessage"`
	Bundle         string `json:"bundle"`
	Tag            string `json:"tag"`
	Coupon         string `json:"coupon"`
	Credits        string `json:"credits"`
}

type BillingProfile struct {
	Invoice          bool   `json:"invoice"`
	PaymentType      string `json:"paymentType"`
	BillingProfileId int    `json:"billingProfileId"`
}

type CorporateOrders struct {
	BillingProfile BillingProfile   `json:"billingProfile"`
	Gifts          []CorporateOrder `json:"gifts"`
	BrandId        int              `json:"brandId"`
}

type CorporateOrderResponse struct {
	Success bool             `json:"success"`
	Gifts   []CorporateOrder `json:"gifts"`
	Message string           `json:"message"`
}

func CreateCorporateOrders(records [][]string) []CorporateOrder {
	corporateOrders := make([]CorporateOrder, len(records)-1)
	for i, record := range records[1:] {
		corporateOrders[i] = CreateCorporateOrder(record)
	}
	return corporateOrders
}

func CreateCorporateOrder(record []string) CorporateOrder {
	corporateOrders := new(CorporateOrder)
	corporateOrders.OrderId = 0
	giftAmount, _ := strconv.Atoi(record[1])
	corporateOrders.GiftAmount = giftAmount
	corporateOrders.ShippingMethod = 4
	corporateOrders.Email = record[3]
	corporateOrders.UserName = record[4]
	corporateOrders.FirstName = record[5]
	corporateOrders.LastName = record[6]
	corporateOrders.Company = record[7]
	corporateOrders.Street1 = record[8]
	corporateOrders.Street2 = record[9]
	corporateOrders.City = record[10]
	corporateOrders.State = record[11]
	corporateOrders.PostalCode = record[12]
	corporateOrders.Phone = record[13]
	corporateOrders.GiftMessage = record[14]
	corporateOrders.Bundle = record[15]
	corporateOrders.Tag = record[16]
	corporateOrders.Coupon = record[17]
	corporateOrders.Credits = record[18]
	return *corporateOrders
}

func CreateBillingProfile(billingProfileId int, invoice bool) BillingProfile {
	billingProfile := new(BillingProfile)
	billingProfile.Invoice = invoice
	billingProfile.BillingProfileId = billingProfileId
	return *billingProfile
}
