package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/eserilev/migration.winc.services/corporate"
)

const path = "./test/"

func TestCorporateOrder(t *testing.T) {
	corporateOrders := new(corporate.CorporateOrders)
	billingProfile := new(corporate.BillingProfile)
	corporateOrderResponse := new(corporate.CorporateOrderResponse)
	content := ReadCsv()
	corporateOrders.Gifts = corporate.CreateCorporateOrders(content)
	corporateOrders.BrandId = 10
	userGuid := "964b9e0f-31b6-4391-bc12-35ec8a8a0460"
	invoice := true
	billingProfile.Invoice = invoice
	responseContent, success := corporate.PostCorporateOrders(*corporateOrders, userGuid)
	if !success {
		t.Fatalf("Failed to post corporate order")
	}

	json.Unmarshal([]byte(responseContent), &corporateOrderResponse)

	if !corporateOrderResponse.Success {
		t.Fatalf("Corporate Order Request returned a 200, but failed to process")
	}

	t.Log("Corporate Order Submitted Succesfully")
}

func ReadCsv() [][]string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Open(path + file.Name())
		if err != nil {
			log.Fatal("Unable to read input file "+path+file.Name(), err)
		}
		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+path+file.Name(), err)
		}
		return records
	}
	return nil
}
