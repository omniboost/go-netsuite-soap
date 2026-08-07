package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-netsuite-soap"
)

// Custom segment POC tests for IN-4714, against the Omniboost Demo sandbox
// (account 8436976_SB1). Fixtures used throughout:
//   - segment scriptId "cseg1" ("TEST_CUSTOM_SEGMENT"), internalId 1
//   - segment's underlying custom record type, internalId 1590
//   - segment values: internalId 1-4 (TEST_VALUE_1..4)

// TestCustomSegmentDiscovery lists all custom segments in the account.
// Requires the "Custom Segments" permission on the integration role.
func TestCustomSegmentDiscovery(t *testing.T) {
	req := client.NewGetCustomizationIDRequest()
	req.RequestBody().CustomizationType.GetCustomizationType = "customSegment"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// TestCustomSegmentDefinition fetches a single custom segment's own
// definition (label, fieldType, recordType, isMandatory, hasGLImpact, ...).
func TestCustomSegmentDefinition(t *testing.T) {
	req := client.NewCustomRecordRefGetRequest()
	req.RequestBody().Xmlns = "urn:messages_2022_2.platform.webservices.netsuite.com"
	req.RequestBody().BaseRef.InternalID = "1"
	req.RequestBody().BaseRef.Type = "customSegment"
	req.RequestBody().BaseRef.XSIType = "platformCore:RecordRef"
	req.RequestBody().BaseRef.Xmlns = "urn:core_2022_2.platform.webservices.netsuite.com"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// TestCustomSegmentValues lists the possible values (rows in the segment's
// underlying custom record type) that can be set on cseg1.
// Requires "value management" permission for the segment.
func TestCustomSegmentValues(t *testing.T) {
	req := client.NewCustomRecordSearchRequest()
	req.RequestBody().SearchRecord.Basic.RecType.InternalID = "1590"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// TestCustomSegmentSetValueOnJournalEntry creates a new Journal Entry with
// cseg1 set to TEST_VALUE_1 (internalId "1"). Requires:
//   - cseg1 applied to the Journal Entry transaction type (Application & Sourcing)
//   - the role's Custom Record permission for cseg1's record type set to Edit/Full
//
// Values must be referenced by internalId (or externalId) - setting only
// CustomFieldValue.Name is silently ignored by NetSuite (no error, no value set).
func TestCustomSegmentSetValueOnJournalEntry(t *testing.T) {
	req := client.NewAddRequest()
	req.RequestBody().Record.Type = "tranGeneral:JournalEntry"
	req.RequestBody().Record.Record = netsuite.JournalEntry{
		Subsidiary: netsuite.RecordRef{InternalID: "1"},
		LineList: netsuite.JournalEntryLines{
			{
				Account: netsuite.RecordRef{InternalID: "58"}, // Expenses
				Debit:   10,
				Memo:    "IN-4714 custom segment POC",
			},
			{
				Account: netsuite.RecordRef{InternalID: "1"}, // ABN AMRO bank account
				Credit:  10,
				Memo:    "IN-4714 custom segment POC",
			},
		},
		CustomFieldList: struct {
			CustomField netsuite.CustomFields `xml:"customField,omitempty"`
		}{
			CustomField: netsuite.CustomFields{
				{
					ScriptID: "cseg1",
					Type:     "platformCore:SelectCustomFieldRef",
					Value:    netsuite.CustomFieldValue{InternalID: "1"}, // TEST_VALUE_1
				},
			},
		},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// TestCustomSegmentGetJournalEntry re-fetches a Journal Entry to confirm
// cseg1's value round-trips through customFieldList.
func TestCustomSegmentGetJournalEntry(t *testing.T) {
	req := client.NewGetRequest()
	req.RequestBody().BaseRef.InternalID = "51646"
	req.RequestBody().BaseRef.Type = "journalEntry"
	req.RequestBody().BaseRef.XSIType = "platformCore:RecordRef"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(resp.ReadResponse.Record.XML))
}
