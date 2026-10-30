package netsuite_test

import (
	"context"
	"io"
	"strings"
	"testing"

	netsuite "github.com/omniboost/go-netsuite-soap"
)

func addRequestXML(t *testing.T, c *netsuite.Client) string {
	t.Helper()

	req := c.NewAddRequest()
	req.RequestBody().Record.Type = "tranGeneral:StatisticalJournalEntry"
	req.RequestBody().Record.Record = netsuite.StatisticalJournalEntry{}

	httpReq, err := c.NewRequest(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(httpReq.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}

func TestRunServerSuiteScriptAndTriggerWorkflowsPreference(t *testing.T) {
	c := netsuite.NewClient(nil)
	c.SetAccountID("123_SB1")

	if xml := addRequestXML(t, c); strings.Contains(xml, "runServerSuiteScriptAndTriggerWorkflows") {
		t.Errorf("preference must be omitted by default, got:\n%s", xml)
	}

	c.Preferences.RunServerSuiteScriptAndTriggerWorkflows = true
	xml := addRequestXML(t, c)
	want := "<platformMsgs:runServerSuiteScriptAndTriggerWorkflows>true</platformMsgs:runServerSuiteScriptAndTriggerWorkflows>"
	if !strings.Contains(xml, want) {
		t.Errorf("want %s in header, got:\n%s", want, xml)
	}
}

func TestRequestPreferencesDoNotLeakToClient(t *testing.T) {
	c := netsuite.NewClient(nil)
	c.SetAccountID("123_SB1")

	req := c.NewAddRequest()
	req.RequestBody().Record.Type = "tranGeneral:StatisticalJournalEntry"
	req.RequestBody().Record.Record = netsuite.StatisticalJournalEntry{}
	req.SetPreferences(netsuite.Preferences{RunServerSuiteScriptAndTriggerWorkflows: true})

	httpReq, err := c.NewRequest(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(httpReq.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "<platformMsgs:runServerSuiteScriptAndTriggerWorkflows>true<") {
		t.Errorf("request-level preference missing from header:\n%s", b)
	}

	// another request on the same client must not carry the preference
	if xml := addRequestXML(t, c); strings.Contains(xml, "runServerSuiteScriptAndTriggerWorkflows") {
		t.Errorf("request-level preference leaked to the client, got:\n%s", xml)
	}
}
