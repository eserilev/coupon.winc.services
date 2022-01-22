//
// Corporate Order Helpers
//

package corporate

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func ProcessOrders(filePath string, userGuid string, invoice bool, billingProfileId int, brandId int) *CorporateOrderResponse {
	corporateOrderResponse := new(CorporateOrderResponse)
	content := ReadCsv(filePath)
	corporateOrders := new(CorporateOrders)
	corporateOrders.Gifts = CreateCorporateOrders(content)
	corporateOrders.BillingProfile = CreateBillingProfile(billingProfileId, invoice)
	corporateOrders.BrandId = brandId
	resultString, success := PostCorporateOrders(*corporateOrders, userGuid)
	if !success {
		return nil
	}

	resultBytes, _ := json.Marshal(resultString)

	json.Unmarshal(resultBytes, &corporateOrderResponse)

	return corporateOrderResponse
}

func ReadCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func ReadCsvFile(r *http.Request) [][]string {
	reader := csv.NewReader(r.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	return records
}

func PostCorporateOrders(corporateOrders CorporateOrders, userGuid string) (string, bool) {
	success := false
	corporateOrderResponse := new(CorporateOrderResponse)
	responseContent := ""
	dest := CwApiBaseUrl() + CorporateOrderRelativePath(userGuid)
	data, _ := json.Marshal(corporateOrders)
	response, err := Post(dest, data)

	if err != nil {
		fmt.Println(err)
		return err.Error(), success
	}

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		responseContent = string(bodyBytes)
		json.Unmarshal([]byte(responseContent), &corporateOrderResponse)
		success = corporateOrderResponse.Success
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
	req.Header.Add("X-Api-Key", "todo")
	res, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
