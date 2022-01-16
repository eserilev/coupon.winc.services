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

	"github.com/eserilev/migration.winc.services/config"
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

func PostCorporateOrders(corporateOrders CorporateOrders, userGuid string) (string, bool) {
	success := false
	responseContent := ""
	dest := config.CwApiBaseUrl() + config.CorporateOrderRelativePath(userGuid)
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
	req.Header.Add("Authorization", "Basic "+basicAuth(config.CwApiUserName(), config.CwApiPassword()))
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
