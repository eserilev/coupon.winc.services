//
// Corporate Order Helpers
//

package corporate

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

//
// DefaultClient
//

var DefaultClient *http.Client = &http.Client{
	CheckRedirect: nil,
	Jar:           nil,
	Timeout:       30 * time.Second,
}

func ProcessCorporateOrders(r *http.Request) ([]byte, bool) {
	content := ReadCsvFile(r)
	userGuid := GetUserGuid(r)
	billingProfile := GetBillingProfile(r)
	corporateOrders := new(CorporateOrders)
	corporateOrders.Gifts = CreateCorporateRecords(content)
	corporateOrders.BillingProfile = billingProfile
	corporateOrders.BrandId = GetBrandId(r)
	resultString, success := PostCorporateOrders(*corporateOrders, userGuid)
	if !success {
		return nil, success
	}

	resultBytes, _ := json.Marshal(resultString)
	return resultBytes, success
}

func ReadCsvFile(r *http.Request) [][]string {
	reader := csv.NewReader(r.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	return records
}

func GetUserGuid(r *http.Request) string {
	return r.URL.Query().Get("userGuid")
}

func GetBrandId(r *http.Request) int {
	brandId, _ := strconv.Atoi(r.URL.Query().Get("brandId"))
	return brandId
}

func GetBillingProfile(r *http.Request) BillingProfile {
	billingProfile := new(BillingProfile)
	queryParams := r.URL.Query()
	invoice, _ := strconv.ParseBool(queryParams.Get("invoice"))
	billingProfileId, _ := strconv.Atoi(queryParams.Get("billingProfileId"))
	paymentType := queryParams.Get("paymentType")
	billingProfile.Invoice = invoice
	billingProfile.PaymentType = paymentType
	billingProfile.BillingProfileId = billingProfileId

	return *billingProfile
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

func PostCorporateOrders(corporateOrders CorporateOrders, userGuid string) (string, bool) {
	success := false
	responseContent := ""
	dest := "http://cwapi-staging.cloudapp.net/winc/users/" + userGuid + "/gift-checkout"
	data, _ := json.Marshal(corporateOrders)
	response, err := Post(dest, data)
	if err != nil {
		fmt.Println(err)
		return err.Error(), success
	}

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		responseContent = bodyString
		success = true
	} else {
		responseContent = response.Status
	}

	return responseContent, success
}

func Post(dest string, data []byte) (*http.Response, error) {
	fmt.Println("Attempting Authorization", dest)
	req, err := http.NewRequest("POST", dest, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic "+basicAuth("api@clubw.com", "randomPassword"))
	res, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
