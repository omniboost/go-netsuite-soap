package netsuite_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/omniboost/go-netsuite-soap"
	"os"
	"testing"
)

func TestSavedSearch(t *testing.T) {
	req := client.NewSavedSearchRequest()
	savedSearchId := os.Getenv("SAVED_SEARCH_ID")

	criteria := netsuite.SavedSearchCriteria{
		XMLName: xml.Name{
			Local: "lastModifiedDate",
		},
		Operator:    "after",
		SearchValue: []string{"2007-02-10T00:16:17.750Z"},
	}

	req.RequestBody().SearchRecord.SavedSearchId = savedSearchId
	req.RequestBody().SearchRecord.Type = "listRel:CustomerSearchAdvanced"
	req.RequestBody().SearchRecord.Criteria.Basic.SavedSearchCriteria = []netsuite.SavedSearchCriteria{criteria}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
