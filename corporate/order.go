//
// Corporate Order Helpers
//

package corporate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/eserilev/utilities.winc.services/winc_csv"
)

//
// DefaultClient
//

var DefaultClient *http.Client = &http.Client{
	CheckRedirect: nil,
	Jar:           nil,
	Timeout:       30 * time.Second,
}

func ProcessCorporateOrders(filePath string, userGuid string, invoice bool, billingProfileId int, brandId int) *CorporateOrderResponse {
	content := winc_csv.ReadCsv(filePath)
	corporateOrders := new(CorporateOrders)
	corporateOrders.Gifts = CreateCorporateOrders(content)
	corporateOrders.BillingProfile = CreateBillingProfile(billingProfileId, invoice)
	corporateOrders.BrandId = brandId
	result := PostCorporateOrders(*corporateOrders, userGuid)
	return result
}

func PostCorporateOrders(corporateOrders CorporateOrders, userGuid string) *CorporateOrderResponse {
	corporateOrderResponse := new(CorporateOrderResponse)
	responseContent := ""
	dest := CwApiBaseUrl() + CorporateOrderRelativePath(userGuid)
	data, _ := json.Marshal(corporateOrders)
	response, err := Post(dest, data)

	if err != nil {
		fmt.Println(err)
		return corporateOrderResponse
	}

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		responseContent = string(bodyBytes)
		json.Unmarshal([]byte(responseContent), &corporateOrderResponse)
	}

	return corporateOrderResponse
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
