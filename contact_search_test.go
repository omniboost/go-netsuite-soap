package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-netsuite-soap"
)

func TestContactSearch(t *testing.T) {
	req := client.NewContactSearchRequest()
	//req.RequestBody().SearchRecord.Basic.Email.Operator = "contains"
	//req.RequestBody().SearchRecord.Basic.Email.SearchValue = ".net"
	//req.RequestBody().SearchRecord.Basic.Email.Operator = "anyOf"
	//req.RequestBody().SearchRecord.Basic.InternalID.Operator = "anyOf"
	//req.RequestBody().SearchRecord.Basic.InternalID.SearchValue = []netsuite.RecordRef{{InternalID: "10242187"}}
	req.RequestBody().SearchRecord.Basic.Company.Operator = "anyOf"
	req.RequestBody().SearchRecord.Basic.Company.SearchValue = []netsuite.RecordRef{{InternalID: "10242087"}}
	// req.RequestBody().SearchRecord.Basic.ExternalID.Operator = "anyOf"
	// req.RequestBody().SearchRecord.Basic.ExternalID.SearchValue = []netsuite.RecordRef{{ExternalID: "10242087"}}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
