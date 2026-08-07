package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomRecordRefGet(t *testing.T) {
	req := client.NewCustomRecordRefGetRequest()
	// req.RequestBody().BaseRef.ScriptID = "cseg_nch_property"
	// req.RequestBody().BaseRef.ExternalID = "Apparel"
	// req.RequestBody().BaseRef.InternalID = "9"
	// req.RequestBody().BaseRef.Type = "customSegment"
	// req.RequestBody().BaseRef.XSIType = "platformCore:RecordRef"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
