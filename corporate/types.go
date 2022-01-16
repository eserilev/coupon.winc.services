package corporate

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
