package main

import (
	"testing"

	"github.com/eserilev/migration.winc.services/corporate"
)

const path = "./test/test.csv"
const userGuid = "964b9e0f-31b6-4391-bc12-35ec8a8a0460"
const invoice = true
const billingProfileId = 0
const brandId = 10

func TestCorporateOrder(t *testing.T) {
	result := corporate.ProcessCorporateOrders(path, userGuid, invoice, billingProfileId, brandId)
	if !result.Success {
		t.Log(result)
		t.Fatalf("Failed to post corporate order")
	}
	t.Log("Corporate Order Submitted Succesfully")
}
