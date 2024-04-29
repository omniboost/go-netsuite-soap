package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-netsuite-soap"
)

func TestTransactionSearch(t *testing.T) {
	client.SearchPreferences.PageSize = 10
	// client.SearchPreferences.ReturnSearchColumns = true
	client.SearchPreferences.BodyFieldsOnly = false
	req := client.NewTransactionSearchRequest()
	// req.RequestBody().SearchRecord.Basic.Name.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.Name.SearchValue = "e"

	// req.RequestBody().SearchRecord.Basic.Number.Operator = "equalTo"
	// req.RequestBody().SearchRecord.Basic.Number.SearchValue = 20002200445

	// req.RequestBody().SearchRecord.Basic.TranID.Operator = "is"
	// req.RequestBody().SearchRecord.Basic.TranID.SearchValue = "99536"

	// req.RequestBody().SearchRecord.Basic.Number.Operator = "equalTo"
	// req.RequestBody().SearchRecord.Basic.Number.SearchValue = "2240087"

	// req.RequestBody().SearchRecord.Basic.Memo.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.Memo.SearchValue = "Mews"

	// req.RequestBody().SearchRecord.Basic.TranDate.Operator = "on"
	// req.RequestBody().SearchRecord.Basic.TranDate.SearchValue = "2023-09-14T15:00:00.000-08:00"

	// req.RequestBody().SearchRecord.Basic.DateCreated.Operator = "after"
	// req.RequestBody().SearchRecord.Basic.DateCreated.SearchValue = "2023-09-15T15:00:00.000-00:00"

	req.RequestBody().SearchRecord.Basic.InternalID.Operator = "anyOf"
	req.RequestBody().SearchRecord.Basic.InternalID.SearchValue = []netsuite.RecordRef{
		{InternalID: "4201593"},
	}

	// req.RequestBody().SearchRecord.Basic.Type.Operator = "anyOf"
	// req.RequestBody().SearchRecord.Basic.Type.SearchValue = "_journal"
	//
	// req.RequestBody().SearchRecord.Basic.MemoMain.Operator = "startsWith"
	// req.RequestBody().SearchRecord.Basic.MemoMain.SearchValue = "Mews 2023-10-17"
	//
	// req.RequestBody().SearchRecord.Basic.TranDate.Operator = "on"
	// req.RequestBody().SearchRecord.Basic.TranDate.SearchValue = "2023-10-17T00:00:00"
	//
	// req.RequestBody().SearchRecord.Basic.Subsidiary.Operator = "anyOf"
	// req.RequestBody().SearchRecord.Basic.Subsidiary.SearchValue = []netsuite.RecordRef{
	// 	{InternalID: "655"},
	// }

	// req.RequestBody().SearchRecord.Basic.CustomFieldList = netsuite.SearchCustomFieldList{
	// 	// netsuite.SearchMultiSelectCustomField{
	// 	// 	ScriptID: "custbody_nch_source",
	// 	// 	Operator: "anyOf",
	// 	// 	SearchValue: netsuite.ListOrRecordRef{
	// 	// 		InternalID: "1",
	// 	// 	},
	// 	// },
	// 	netsuite.SearchMultiSelectCustomField{
	// 		ScriptID: "cseg_nch_property",
	// 		Operator: "anyOf",
	// 		SearchValue: netsuite.ListOrRecordRef{
	// 			InternalID: "4",
	// 		},
	// 	},
	// }

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
