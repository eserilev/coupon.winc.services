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

func CreateCorporateRecords(records [][]string) []CorporateOrder {
	corporateRecords := make([]CorporateOrder, len(records)-1)
	for i, record := range records[1:] {
		corporateRecords[i] = CreateCorporateRecordObject(record)
	}
	return corporateRecords
}

func CreateCorporateRecordObject(record []string) CorporateOrder {
	corporateRecord := new(CorporateOrder)
	corporateRecord.OrderId = 0
	giftAmount, _ := strconv.Atoi(record[1])
	corporateRecord.GiftAmount = giftAmount
	corporateRecord.ShippingMethod = 4
	corporateRecord.Email = record[3]
	corporateRecord.UserName = record[4]
	corporateRecord.FirstName = record[5]
	corporateRecord.LastName = record[6]
	corporateRecord.Company = record[7]
	corporateRecord.Street1 = record[8]
	corporateRecord.Street2 = record[9]
	corporateRecord.City = record[10]
	corporateRecord.State = record[11]
	corporateRecord.PostalCode = record[12]
	corporateRecord.Phone = record[13]
	corporateRecord.GiftMessage = record[14]
	corporateRecord.Bundle = record[15]
	corporateRecord.Tag = record[16]
	corporateRecord.Coupon = record[17]
	corporateRecord.Credits = record[18]
	return *corporateRecord
}
