package netsuite

import (
	"encoding/xml"
	"strings"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
)

type RecordRef struct {
	InternalID string `xml:"internalId,attr,omitempty"`
	ExternalID string `xml:"externalId,attr,omitempty"`
	Name       string `xml:"name,omitempty"`
}

func (r RecordRef) IsEmpty() bool {
	return zero.IsZero(r)
}

type Customers []Customer

type Customer struct {
	InternalID           string    `xml:"internalId,attr,omitempty"`
	ExternalID           string    `xml:"externalId,attr,omitempty"`
	Type                 string    `xml:"type,attr,omitempty"`
	ListRel              string    `xml:"listRel,attr,omitempty"`
	EntityID             string    `xml:"entityId,omitempty"`
	IsPerson             bool      `xml:"isPerson,omitempty"`
	CompanyName          string    `xml:"companyName,omitempty"`
	EntityStatus         RecordRef `xml:"entityStatus,omitempty"`
	Phone                string    `xml:"phone,omitempty"`
	Email                string    `xml:"email,omitempty"`
	DefaultAddress       string    `xml:"defaultAddress,omitempty"`
	IsInactive           string    `xml:"isInactive,omitempty"`
	Language             string    `xml:"language,omitempty"`
	DateCreated          Date      `xml:"dateCreated,omitempty"`
	Subsidiary           RecordRef `xml:"subsidiary,omitempty"`
	VatRegNumber         string    `xml:"vatRegNumber,omitempty"`
	Terms                RecordRef `xml:"terms,omitempty"`
	CreditLimit          string    `xml:"creditLimit,omitempty"`
	UnbilledOrders       string    `xml:"unbilledOrders,omitempty"`
	Currency             RecordRef `xml:"currency,omitempty"`
	ShipComplete         string    `xml:"shipComplete,omitempty"`
	AlcoholRecipientType string    `xml:"alcoholRecipientType,omitempty"`
	ReceivablesAccount   RecordRef `xml:"receivablesAccount,omitempty"`
	LastModifiedDate     DateTime  `xml:"lastModifiedDate,omitempty"`
	Stage                string    `xml:"stage,omitempty"`
	IsBudgetApproved     string    `xml:"isBudgetApproved,omitempty"`
	FirstName            string    `xml:"firstName,omitempty"`
	LastName             string    `xml:"lastName,omitempty"`
	CustomFieldList      struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

func (c Customer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c Customer) IsEmpty() bool {
	return zero.IsZero(c)
}

type CustomFields []CustomField

type CustomField struct {
	InternalID string           `xml:"internalId,attr,omitempty"`
	ScriptID   string           `xml:"scriptId,attr,omitempty"`
	Type       string           `xml:"xsi:type,attr,omitempty"`
	Value      CustomFieldValue `xml:"value,omitempty"`
}

func (c *CustomField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias CustomField
	s := struct {
		Type string `xml:"type,attr,omitempty"`
		*Alias
	}{Alias: (*Alias)(c)}

	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	c.Type = s.Type
	return nil
}

type Accounts []Account

type Account struct {
	InternalID      string    `xml:"internalId,attr"`
	ExternalID      string    `xml:"externalId,attr"`
	Type            string    `xml:"type,attr"`
	ListAcct        string    `xml:"listAcct,attr"`
	AcctType        string    `xml:"acctType"`
	AcctNumber      string    `xml:"acctNumber"`
	AcctName        string    `xml:"acctName"`
	IncludeChildren string    `xml:"includeChildren"`
	GeneralRate     string    `xml:"generalRate"`
	Parent          RecordRef `xml:"parent"`
	CashFlowRate    string    `xml:"cashFlowRate"`
	IsInactive      string    `xml:"isInactive"`
	Inventory       string    `xml:"inventory"`
	Eliminate       string    `xml:"eliminate"`
	Revalue         string    `xml:"revalue"`
	CustomFieldList struct {
		CustomFields CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
	Currency RecordRef `xml:"currency"`
}

type Departments []Department

type Department struct {
	InternalID      string    `xml:"internalId,attr"`
	Type            string    `xml:"type,attr"`
	ListAcct        string    `xml:"listAcct,attr"`
	Name            string    `xml:"name"`
	IncludeChildren string    `xml:"includeChildren"`
	IsInactive      string    `xml:"isInactive"`
	Parent          RecordRef `xml:"parent"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type Record struct {
	Type   string `xml:"xsi:type,attr"`
	Record interface{}
}

func (r Record) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: r.Type})
	return omitempty.MarshalXML(r.Record, e, start)
}

type Classifications []Classification

type Classification struct {
	InternalID      string `xml:"internalId,attr"`
	Type            string `xml:"type,attr"`
	ListAcct        string `xml:"listAcct,attr"`
	Name            string `xml:"name"`
	IncludeChildren string `xml:"includeChildren"`
	IsInactive      string `xml:"isInactive"`
	CustomFieldList struct {
		CustomFields CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type JournalEntry struct {
	TranGeneral      string            `xml:"tranGeneral,attr,omitempty"`
	InternalId       string            `xml:"internalId,attr,omitempty"`
	ExternalId       string            `xml:"externalId,attr,omitempty"`
	Type             string            `xml:"type,attr,omitempty"`
	TranDate         Date              `xml:"tranDate,omitempty"`
	Currency         RecordRef         `xml:"currency,omitempty"`
	ExchangeRate     string            `xml:"exchangeRate,omitempty"`
	TranId           string            `xml:"tranId,omitempty"`
	Subsidiary       RecordRef         `xml:"subsidiary,omitempty"`
	CreatedDate      Date              `xml:"createdDate,omitempty"`
	LastModifiedDate Date              `xml:"lastModifiedDate,omitempty"`
	LineList         JournalEntryLines `xml:"lineList>line"`
	Memo             string            `xml:"memo,omitempty"`
	Location         RecordRef         `xml:"location,omitempty"`
	CustomFieldList  struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
}

func (j JournalEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j JournalEntry) IsEmpty() bool {
	return zero.IsZero(j)
}

// type JournalEntryLineList struct {
// 	Line JournalEntryLines
// }

// func (j JournalEntryLineList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	return omitempty.MarshalXML(j, e, start)
// }

// func (j JournalEntryLineList) IsEmpty() bool {
// 	return zero.IsZero(j)
// }

type JournalEntryLines []JournalEntryLine

type JournalEntryLine struct {
	XMLName xml.Name `xml:"line"`

	Account         RecordRef `xml:"account"`
	Line            int       `xml:"line,omitempty"`
	Debit           float64   `xml:"debit"`
	Memo            string    `xml:"memo"`
	Eliminate       string    `xml:"eliminate,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	Credit     float64   `xml:"credit"`
	Location   RecordRef `xml:"location,omitempty"`
	Department RecordRef `xml:"department,omitempty"`
	Class      RecordRef `xml:"class,omitempty"`
	TaxCode    RecordRef `xml:"taxCode,omitempty"`
	Tax1Amt    *float64  `xml:"tax1Amt,omitempty"`
	Tax1Acct   RecordRef `xml:"tax1Acct,omitempty"`
}

func (j JournalEntryLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j JournalEntryLine) IsEmpty() bool {
	return zero.IsZero(j)
}

type SalesTaxItems []SalesTaxItem

type SalesTaxItem struct {
	ListAcct       string    `xml:"listAcct,attr"`
	InternalID     string    `xml:"internalId,attr"`
	Type           string    `xml:"type,attr"`
	ItemID         string    `xml:"itemId"`
	Description    string    `xml:"description"`
	Rate           string    `xml:"rate"`
	TaxType        RecordRef `xml:"taxType"`
	TaxAgency      RecordRef `xml:"taxAgency"`
	IsInactive     string    `xml:"isInactive"`
	SubsidiaryList struct {
		RecordRef RecordRef `xml:"recordRef"`
	} `xml:"subsidiaryList"`
	IncludeChildren       string `xml:"includeChildren"`
	Eccode                string `xml:"eccode"`
	ReverseCharge         string `xml:"reverseCharge"`
	Service               string `xml:"service"`
	Exempt                string `xml:"exempt"`
	IsDefault             string `xml:"isDefault"`
	ExcludeFromTaxReports string `xml:"excludeFromTaxReports"`
	Available             string `xml:"available"`
	Export                string `xml:"export"`
	CustomFieldList       struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	Parent RecordRef `xml:"parent"`
}

type Invoice struct {
	TranSales             string          `xml:"tranSales,attr,omitempty"`
	InternalId            string          `xml:"internalId,attr,omitempty"`
	Type                  string          `xml:"type,attr,omitempty"`
	ExternalId            string          `xml:"externalId,attr,omitempty"`
	CreatedDate           Date            `xml:"createdDate,omitempty"`
	LastModifiedDate      Date            `xml:"lastModifiedDate,omitempty"`
	Entity                RecordRef       `xml:"entity,omitempty"`
	TranDate              Date            `xml:"tranDate,omitempty"`
	TranId                string          `xml:"tranId,omitempty"`
	Department            RecordRef       `xml:"department,omitempty"`
	Terms                 RecordRef       `xml:"terms,omitempty"`
	Subsidiary            RecordRef       `xml:"subsidiary,omitempty"`
	Currency              RecordRef       `xml:"currency,omitempty"`
	DueDate               Date            `xml:"dueDate,omitempty"`
	OtherRefNum           string          `xml:"otherRefNum,omitempty"`
	Location              RecordRef       `xml:"location,omitempty"`
	Memo                  string          `xml:"memo,omitempty"`
	TotalCostEstimate     string          `xml:"totalCostEstimate,omitempty"`
	EstGrossProfit        string          `xml:"estGrossProfit,omitempty"`
	EstGrossProfitPercent string          `xml:"estGrossProfitPercent,omitempty"`
	Account               RecordRef       `xml:"account,omitempty"`
	ExchangeRate          string          `xml:"exchangeRate,omitempty"`
	CurrencyName          string          `xml:"currencyName,omitempty"`
	BillingAddress        Address         `xml:"billingAddress,omitempty"`
	ShippingAddress       Address         `xml:"shippingAddress,omitempty"`
	ShipIsResidential     string          `xml:"shipIsResidential,omitempty"`
	SubTotal              string          `xml:"subTotal,omitempty"`
	CanHaveStackable      string          `xml:"canHaveStackable,omitempty"`
	TaxTotal              string          `xml:"taxTotal,omitempty"`
	Total                 string          `xml:"total,omitempty"`
	Status                string          `xml:"status,omitempty"`
	Email                 string          `xml:"email,omitempty"`
	VatRegNum             string          `xml:"vatRegNum,omitempty"`
	ItemCostDiscPrint     string          `xml:"itemCostDiscPrint"`
	ExpCostDiscPrint      string          `xml:"expCostDiscPrint"`
	ItemList              InvoiceItemList `xml:"itemList>item"`
	OverrideInstallments  string          `xml:"overrideInstallments"`
	CustomFieldList       struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type Transactions []Transaction

type Transaction struct {
	Text             string `xml:",chardata"`
	TranSales        string `xml:"tranSales,attr"`
	InternalID       string `xml:"internalId,attr"`
	Type             string `xml:"type,attr"`
	ExternalID       string `xml:"externalId,attr"`
	CreatedDate      string `xml:"createdDate"`
	LastModifiedDate string `xml:"lastModifiedDate"`
	Entity           struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"entity"`
	TranDate string `xml:"tranDate"`
	TranID   string `xml:"tranId"`
	Source   string `xml:"source"`
	Terms    struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"terms"`
	Location struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"location"`
	Subsidiary struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"subsidiary"`
	Currency struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"currency"`
	DueDate               string `xml:"dueDate"`
	OtherRefNum           string `xml:"otherRefNum"`
	Memo                  string `xml:"memo"`
	TotalCostEstimate     string `xml:"totalCostEstimate"`
	EstGrossProfit        string `xml:"estGrossProfit"`
	EstGrossProfitPercent string `xml:"estGrossProfitPercent"`
	Account               struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"account"`
	ExchangeRate   string `xml:"exchangeRate"`
	CurrencyName   string `xml:"currencyName"`
	BillingAddress struct {
		Text            string `xml:",chardata"`
		PlatformCommon  string `xml:"platformCommon,attr"`
		InternalID      string `xml:"internalId"`
		Country         string `xml:"country"`
		Attention       string `xml:"attention"`
		Addr1           string `xml:"addr1"`
		City            string `xml:"city"`
		Zip             string `xml:"zip"`
		AddrText        string `xml:"addrText"`
		Override        string `xml:"override"`
		CustomFieldList struct {
			CustomField CustomFields `xml:"customField"`
		} `xml:"customFieldList"`
	} `xml:"billingAddress"`
	BillAddressList struct {
		Text       string `xml:",chardata"`
		InternalID string `xml:"internalId,attr"`
		Name       string `xml:"name"`
	} `xml:"billAddressList"`
	ShippingAddress struct {
		Text            string `xml:",chardata"`
		PlatformCommon  string `xml:"platformCommon,attr"`
		InternalID      string `xml:"internalId"`
		Country         string `xml:"country"`
		Attention       string `xml:"attention"`
		Addr1           string `xml:"addr1"`
		City            string `xml:"city"`
		Zip             string `xml:"zip"`
		AddrText        string `xml:"addrText"`
		Override        string `xml:"override"`
		CustomFieldList struct {
			CustomField CustomFields `xml:"customField"`
		} `xml:"customFieldList"`
	} `xml:"shippingAddress"`
	ShipIsResidential string `xml:"shipIsResidential"`
	SubTotal          string `xml:"subTotal"`
	CanHaveStackable  string `xml:"canHaveStackable"`
	TaxTotal          string `xml:"taxTotal"`
	Total             string `xml:"total"`
	Status            string `xml:"status"`
	VatRegNum         string `xml:"vatRegNum"`
	ItemCostDiscPrint string `xml:"itemCostDiscPrint"`
	ExpCostDiscPrint  string `xml:"expCostDiscPrint"`
	TimeDiscPrint     string `xml:"timeDiscPrint"`
	ItemList          struct {
		Text string `xml:",chardata"`
		Item []struct {
			Text string `xml:",chardata"`
			Item struct {
				Text       string `xml:",chardata"`
				InternalID string `xml:"internalId,attr"`
				Name       string `xml:"name"`
			} `xml:"item"`
			Line        string `xml:"line"`
			Description string `xml:"description"`
			Amount      string `xml:"amount"`
			Quantity    string `xml:"quantity"`
			Price       struct {
				Text       string `xml:",chardata"`
				InternalID string `xml:"internalId,attr"`
			} `xml:"price"`
			Department struct {
				Text       string `xml:",chardata"`
				InternalID string `xml:"internalId,attr"`
				Name       string `xml:"name"`
			} `xml:"department"`
			GrossAmt string `xml:"grossAmt"`
			Tax1Amt  string `xml:"tax1Amt"`
			TaxCode  struct {
				Text       string `xml:",chardata"`
				InternalID string `xml:"internalId,attr"`
				Name       string `xml:"name"`
			} `xml:"taxCode"`
			TaxRate1        string `xml:"taxRate1"`
			CustomFieldList struct {
				CustomField CustomFields `xml:"customField"`
			} `xml:"customFieldList"`
		} `xml:"item"`
	} `xml:"itemList"`
	OverrideInstallments string            `xml:"overrideInstallments"`
	LineList             JournalEntryLines `xml:"lineList>line"`
	CustomFieldList      struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
}

type Address struct {
	PlatformCommon  string `xml:"platformCommon,attr"`
	InternalID      string `xml:"internalId"`
	Country         string `xml:"country"`
	Attention       string `xml:"attention"`
	Addressee       string `xml:"addressee"`
	Addr1           string `xml:"addr1"`
	City            string `xml:"city"`
	Zip             string `xml:"zip"`
	AddrText        string `xml:"addrText"`
	Override        string `xml:"override"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
	Addr2 string `xml:"addr2"`
}

type InvoiceItemList []InvoiceItem

type InvoiceItem struct {
	Item            RecordRef `xml:"item,omitempty"`
	Line            string    `xml:"line,omitempty"`
	Description     string    `xml:"description,omitempty"`
	Amount          float64   `xml:"amount,omitempty"`
	Quantity        string    `xml:"quantity,omitempty"`
	Price           RecordRef `xml:"price,omitempty"`
	Rate            float64   `xml:"rate,omitempty"`
	Department      RecordRef `xml:"department,omitempty"`
	Location        RecordRef `xml:"location,omitempty"`
	GrossAmt        float64   `xml:"grossAmt,omitempty"`
	Tax1Amt         float64   `xml:"tax1Amt,omitempty"`
	Tax1Acct        RecordRef `xml:"tax1Acct,omitempty"`
	TaxCode         RecordRef `xml:"taxCode,omitempty"`
	TaxRate1        float64   `xml:"taxRate1,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
}

func (i InvoiceItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

func (i InvoiceItem) IsEmpty() bool {
	return zero.IsZero(i)
}

type Status struct {
	IsSuccess    bool `xml:"isSuccess,attr"`
	StatusDetail struct {
		Type    string `xml:"type,attr"`
		Code    string `xml:"code"`
		Message string `xml:"message"`
	} `xml:"statusDetail"`
}

func (s Status) Error() string {
	if s.IsSuccess == false && s.StatusDetail.Message != "" {
		s := []string{s.StatusDetail.Type, s.StatusDetail.Code, s.StatusDetail.Message}
		return strings.Join(s, ", ")
	}

	return ""
}

type CustomFieldValue struct {
	InternalID string `xml:"internalId,attr,omitempty"`
	TypeID     string `xml:"typeId,attr,omitempty"`
	Name       string `xml:"name,omitempty"`
	Text       string `xml:",chardata"`
}

func (c CustomFieldValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c CustomFieldValue) IsEmpty() bool {
	return zero.IsZero(c)
}

type CreditMemo struct {
	TranSales             string    `xml:"tranSales,attr,omitempty"`
	InternalID            string    `xml:"internalId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"`
	ExternalID            string    `xml:"externalId,attr,omitempty"`
	CreatedDate           Date      `xml:"createdDate,omitempty"`
	LastModifiedDate      Date      `xml:"lastModifiedDate,omitempty"`
	Entity                RecordRef `xml:"entity,omitempty"`
	TranDate              Date      `xml:"tranDate,omitempty"`
	TranID                string    `xml:"tranId,omitempty"`
	Department            RecordRef `xml:"department,omitempty"`
	Terms                 RecordRef `xml:"terms,omitempty"`
	Subsidiary            RecordRef `xml:"subsidiary,omitempty"`
	Currency              RecordRef `xml:"currency,omitempty"`
	DueDate               Date      `xml:"dueDate,omitempty"`
	OtherRefNum           string    `xml:"otherRefNum,omitempty"`
	Location              RecordRef `xml:"location,omitempty"`
	Memo                  string    `xml:"memo,omitempty"`
	TotalCostEstimate     string    `xml:"totalCostEstimate,omitempty"`
	EstGrossProfit        string    `xml:"estGrossProfit,omitempty"`
	EstGrossProfitPercent string    `xml:"estGrossProfitPercent,omitempty"`
	Account               RecordRef `xml:"account,omitempty"`
	ExchangeRate          string    `xml:"exchangeRate,omitempty"`
	CurrencyName          string    `xml:"currencyName,omitempty"`
	BillingAddress        Address   `xml:"billingAddress,omitempty"`
	ShippingAddress       Address   `xml:"shippingAddress,omitempty"`
	ShipIsResidential     string    `xml:"shipIsResidential,omitempty"`
	SubTotal              string    `xml:"subTotal,omitempty"`
	CanHaveStackable      string    `xml:"canHaveStackable,omitempty"`
	TaxTotal              string    `xml:"taxTotal,omitempty"`
	Total                 string    `xml:"total,omitempty"`
	Status                string    `xml:"status,omitempty"`
	Email                 string    `xml:"email,omitempty"`
	VatRegNum             string    `xml:"vatRegNum,omitempty"`
	// ItemCostDiscPrint     string         `xml:"itemCostDiscPrint"`
	// ExpCostDiscPrint     string             `xml:"expCostDiscPrint"`
	ItemList CreditMemoItemList `xml:"itemList>item"`
	// OverrideInstallments string             `xml:"overrideInstallments"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField"`
	} `xml:"customFieldList"`
	// Check this box to automatically apply credit to the oldest open
	// receivable. If you clear this box, the credit amount defaults to
	// unapplied and you must later apply the credit amount to an open
	// receivable. Once you check or uncheck the Auto Apply box on a credit memo
	// and save it, the box retains the setting the next time you enter a credit
	// memo.
	// AutoApply   bool                `xml:"autoapply"`
	// CreatedFrom int                 `xml:"createdfrom,omitempty"`
	// ApplyList   CreditMemoApplyList `xml:"applyList"`
}

type CreditMemoItemList []CreditMemoItem

type CreditMemoItem struct {
	Item            RecordRef `xml:"item,omitempty"`
	Line            string    `xml:"line,omitempty"`
	Description     string    `xml:"description,omitempty"`
	Amount          float64   `xml:"amount,omitempty"`
	Quantity        string    `xml:"quantity,omitempty"`
	Price           RecordRef `xml:"price,omitempty"`
	Rate            float64   `xml:"rate,omitempty"`
	Department      RecordRef `xml:"department,omitempty"`
	Location        RecordRef `xml:"location,omitempty"`
	GrossAmt        float64   `xml:"grossAmt,omitempty"`
	Tax1Amt         float64   `xml:"tax1Amt,omitempty"`
	Tax1Acct        RecordRef `xml:"tax1Acct,omitempty"`
	TaxCode         RecordRef `xml:"taxCode,omitempty"`
	TaxRate1        float64   `xml:"taxRate1,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
}

func (i CreditMemoItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

func (i CreditMemoItem) IsEmpty() bool {
	return zero.IsZero(i)
}

// https://www.netsuite.com/help/helpcenter/en_US/srbrowser/Browser2017_1/schema/record/statisticaljournalentry.html
type StatisticalJournalEntry struct {
	// Approved        bool `xml:"approved"`
	CreatedDate     Date `xml:"createdDate,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	// Select the custom journal entry record, if required.
	CustomForm         RecordRef                    `xml:"customForm,omitempty"`
	LastModifiedDate   Date                         `xml:"lastModifiedDate,omitempty"`
	LineList           StatisticalJournalEntryLines `xml:"lineList>statisticalJournalEntryLine"`
	Memo               string                       `xml:"memo,omitempty"`
	ParentExpenseAlloc RecordRef                    `xml:"parentExpenseAlloc,omitempty"`
	PostingPeriod      RecordRef                    `xml:"postingPeriod,omitempty"`
	ReversalDate       Date                         `xml:"reversalDate,omitempty"`
	ReversalDefer      bool                         `xml:"reversalDefer,omitempty"`
	Subsidiary         RecordRef                    `xml:"subsidiary,omitempty"`
	TranDate           Date                         `xml:"tranDate,omitempty"`
	TranID             string                       `xml:"tranId,omitempty"`
	InternalId         string                       `xml:"internalId,attr,omitempty"`
	ExternalId         string                       `xml:"externalId,attr,omitempty"`
	// This field displays the unit type to associate with this statistical journal entry.
	UnitsType RecordRef `xml:"unitsType,omitempty"`
}

func (j StatisticalJournalEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j StatisticalJournalEntry) IsEmpty() bool {
	return zero.IsZero(j)
}

type StatisticalJournalEntryLines []StatisticalJournalEntryLine

type StatisticalJournalEntryLine struct {
	XMLName xml.Name `xml:"line"`

	Account         RecordRef `xml:"account"`
	Class           RecordRef `xml:"class,omitempty"`
	CustomFieldList struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	Debit        float64   `xml:"debit"`
	Department   RecordRef `xml:"department,omitempty"`
	Entity       RecordRef `xml:"entity,omitempty"`
	Line         int       `xml:"line,omitempty"`
	LineUnit     RecordRef `xml:"lineUnit,omitempty"`
	Location     RecordRef `xml:"location,omitempty"`
	Memo         string    `xml:"memo,omitempty"`
	PreviewDebit string    `xml:"previewDebit,omitempty"`
	ScheduleNum  RecordRef `xml:"scheduleNum,omitempty"`
}

func (j StatisticalJournalEntryLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

func (j StatisticalJournalEntryLine) IsEmpty() bool {
	return zero.IsZero(j)
}

type CustomRecords []CustomRecord

type CustomRecord struct {
}

type CustomFieldTypes []CustomFieldType

type CustomFieldType struct {
}

type CreditMemoApplyList struct {
	Apply []CreditMemoApply `xml:"apply"`
}

type CreditMemoApply struct {
	Apply     bool     `xml:"apply"`
	Doc       int      `xml:"doc"`
	ApplyDate DateTime `xml:"applyDate"`
	Job       string   `xml:"job"`
	Type      string   `xml:"type"`
	RefNum    string   `xml:"refNum"`
	Total     float64  `xml:"total"`
	Due       float64  `xml:"due"`
	Currency  string   `xml:"currency"`
	Amount    float64  `xml:"amount"`
	Line      int      `xml:"line"`
}

type Currencies []Currency

type Currency struct {
	InternalID             string `xml:"internalId,attr"`
	ExternalID             string `xml:"externalId,attr"`
	Type                   string `xml:"type,attr"`
	ListAcct               string `xml:"listAcct,attr"`
	Name                   string `xml:"name"`
	Symbol                 string `xml:"symbol"`
	IsBaseCurrency         string `xml:"isBaseCurrency"`
	IsInactive             string `xml:"isInactive"`
	OverrideCurrencyFormat string `xml:"overrideCurrencyFormat"`
	DisplaySymbol          string `xml:"displaySymbol"`
	SymbolPlacement        string `xml:"symbolPlacement"`
	Locale                 string `xml:"locale"`
	FormatSample           string `xml:"formatSample"`
	ExchangeRate           string `xml:"exchangeRate"`
	FxRateUpdateTimezone   string `xml:"fxRateUpdateTimezone"`
	CurrencyPrecision      string `xml:"currencyPrecision"`
}

type CustomizationRefs []CustomizationRef

type CustomizationRef struct {
	ScriptID   string `xml:"scriptId,attr"`
	InternalID string `xml:"internalId,attr"`
	Type       string `xml:"type,attr"`
	Name       string `xml:"name"`
}

type RecordType struct {
	InternalID   string `xml:"internalId,attr"`
	PlatformCore string `xml:"platformCore,attr"`
	Name         string `xml:"name"`
}

type BaseRef struct {
	ScriptID   string `xml:"scriptId,attr,omitempty"`
	ExternalID string `xml:"externalId,attr,omitempty"`
	InternalID string `xml:"internalId,attr,omitempty"`
	Type       string `xml:"type,attr"`
	XSIType    string `xml:"xsi:type,attr"`
	TypeID     string `xml:"typeId,attr,omitempty"`
	// Xmlns      string `xml:"xmlns:q1,attr"`
}

type TransactionBodyCustomField struct {
	Text        string `xml:",chardata"`
	InternalId  string `xml:"internalId,attr"`
	Type        string `xml:"type,attr"`
	SetupCustom string `xml:"setupCustom,attr"`
	Label       string `xml:"label"`
	Owner       struct {
		Text         string `xml:",chardata"`
		InternalId   string `xml:"internalId,attr"`
		PlatformCore string `xml:"platformCore,attr"`
		Name         string `xml:"name"`
	} `xml:"owner"`
	StoreValue   bool `xml:"storeValue"`
	ShowInList   bool `xml:"showInList"`
	GlobalSearch bool `xml:"globalSearch"`
	IsParent     bool `xml:"isParent"`
	Subtab       struct {
		Text         string `xml:",chardata"`
		InternalId   string `xml:"internalId,attr"`
		PlatformCore string `xml:"platformCore,attr"`
		Name         string `xml:"name"`
	} `xml:"subtab"`
	DisplayType              string `xml:"displayType"`
	IsMandatory              bool   `xml:"isMandatory"`
	CheckSpelling            bool   `xml:"checkSpelling"`
	DefaultChecked           bool   `xml:"defaultChecked"`
	IsFormula                bool   `xml:"isFormula"`
	BodyPurchase             bool   `xml:"bodyPurchase"`
	BodySale                 bool   `xml:"bodySale"`
	BodyJournal              bool   `xml:"bodyJournal"`
	BodyExpenseReport        bool   `xml:"bodyExpenseReport"`
	BodyTransferOrder        bool   `xml:"bodyTransferOrder"`
	BodyItemReceipt          bool   `xml:"bodyItemReceipt"`
	BodyItemReceiptOrder     bool   `xml:"bodyItemReceiptOrder"`
	BodyItemFulfillment      bool   `xml:"bodyItemFulfillment"`
	BodyItemFulfillmentOrder bool   `xml:"bodyItemFulfillmentOrder"`
	BodyInventoryAdjustment  bool   `xml:"bodyInventoryAdjustment"`
	BodyPrintFlag            bool   `xml:"bodyPrintFlag"`
	BodyPickingTicket        bool   `xml:"bodyPickingTicket"`
	BodyOtherTransaction     bool   `xml:"bodyOtherTransaction"`
	BodyPrintPackingSlip     bool   `xml:"bodyPrintPackingSlip"`
	BodyCustomerPayment      bool   `xml:"bodyCustomerPayment"`
	BodyVendorPayment        bool   `xml:"bodyVendorPayment"`
	BodyDeposit              bool   `xml:"bodyDeposit"`
	BodyPrintStatement       bool   `xml:"bodyPrintStatement"`
	AccessLevel              string `xml:"accessLevel"`
	SearchLevel              string `xml:"searchLevel"`
	TranslationsList         struct {
		Text         string   `xml:",chardata"`
		Translations []string `xml:"translations"`
	} `xml:"translationsList"`
	FieldType string `xml:"fieldType"`
	ScriptId  string `xml:"scriptId"`
}

type CustomLists []CustomList

type CustomList struct {
	Name                  string                     `xml:"name"`
	Owner                 RecordRef                  `xml:"owner"`
	IsOrdered             bool                       `xml:"isOrdered"`
	Description           string                     `xml:"description"`
	IsMatrixOption        bool                       `xml:"isMaxtrixOption"`
	ScriptID              string                     `xml:"scriptId"`
	ConvertToCustomRecord bool                       `xml:"convertToCustomRecord"`
	CustomValueList       CustomListCustomValueList  `xml:"customValueList"`
	TranslationsList      CustomListTranslationsList `xml:"translationsList>translations"`
}

type CustomListCustomValueList struct {
	CustomValue []CustomListCustomValue `xml:"customValue"`
}

type CustomListCustomValue struct {
	Value                  string            `xml:"value"`
	Abbreviation           string            `xml:"abbreviation"`
	IsInactive             bool              `xml:"isInactive"`
	ValueID                int               `xml:"valueId"`
	ValueLanguageValueList LanguageValueList `xml:"valueLanguageValueList"`
}

type CustomListTranslationsList struct {
	Translations CustomListTranslations `xml:"translations"`
}

type CustomListTranslations struct {
	Locale            Language `xml:"locale"`
	LocaleDescription string   `xml:"localeDescription"`
	Name              string   `xml:"name"`
}

type Language struct {
}

type LanguageValueList struct {
	LanguageValue LanguageValue `xml:"languageValue"`
}

type LanguageValue struct {
	Locale Language `xml:"locale"`
	Value  string   `xml:"value"`
}

type Items []Item

type Item struct {
}

type ItemSearchRecordBasic struct {
	XMLName xml.Name `xml:"platformMsgs:searchRecord"`

	Type  string                  `xml:"xsi:type,attr"`
	Basic SalesTaxItemSearchBasic `xml:"basic"`
}

type Locations []Location

type Location RecordRef

// CategoryList represents a list of categories.
type CategoryList struct {
	Category []*RecordRef `xml:"category,omitempty"`
}

// ContactAddress represents an address in the contact's address book.
type ContactAddress struct {
	DefaultShipping bool    `xml:"defaultShipping,omitempty"`
	DefaultBilling  bool    `xml:"defaultBilling,omitempty"`
	Label           string  `xml:"label,omitempty"`
	Address         Address `xml:"addressbookAddress,omitempty"`
	InternalID      string  `xml:"internalId,omitempty"`
}

// ContactAddressbookList represents a list of addresses in the contact's address book.
type ContactAddressbookList struct {
	Addressbook []*ContactAddress `xml:"addressbook,omitempty"`
	ReplaceAll  bool              `xml:"replaceAll,attr,omitempty"`
}

// Subscriptions represents subscription details.
type Subscriptions struct {
	Subscribed       *bool      `xml:"subscribed,omitempty"`
	Subscription     *RecordRef `xml:"subscription,omitempty"`
	LastModifiedDate Date       `xml:"lastModifiedDate,omitempty"`
}

// SubscriptionsList represents a list of subscriptions.
type SubscriptionsList struct {
	Subscriptions []*Subscriptions `xml:"subscriptions,omitempty"`
	ReplaceAll    bool             `xml:"replaceAll,attr,omitempty"`
}

type Contacts []Contact

// Contact represents the structure of the Contact based on the provided XSD schema.
type Contact struct {
	CustomForm        *RecordRef              `xml:"customForm,omitempty"`
	EntityID          string                  `xml:"entityId,omitempty"`
	ContactSource     *RecordRef              `xml:"contactSource,omitempty"`
	Company           *RecordRef              `xml:"company,omitempty"`
	Salutation        string                  `xml:"salutation,omitempty"`
	FirstName         string                  `xml:"firstName,omitempty"`
	MiddleName        string                  `xml:"middleName,omitempty"`
	LastName          string                  `xml:"lastName,omitempty"`
	Title             string                  `xml:"title,omitempty"`
	Phone             string                  `xml:"phone,omitempty"`
	Fax               string                  `xml:"fax,omitempty"`
	Email             string                  `xml:"email,omitempty"`
	DefaultAddress    string                  `xml:"defaultAddress,omitempty"`
	IsPrivate         *bool                   `xml:"isPrivate,omitempty"`
	IsInactive        *bool                   `xml:"isInactive,omitempty"`
	Subsidiary        *RecordRef              `xml:"subsidiary,omitempty"`
	PhoneticName      string                  `xml:"phoneticName,omitempty"`
	CategoryList      *CategoryList           `xml:"categoryList,omitempty"`
	AltEmail          string                  `xml:"altEmail,omitempty"`
	OfficePhone       string                  `xml:"officePhone,omitempty"`
	HomePhone         string                  `xml:"homePhone,omitempty"`
	MobilePhone       string                  `xml:"mobilePhone,omitempty"`
	Supervisor        *RecordRef              `xml:"supervisor,omitempty"`
	SupervisorPhone   string                  `xml:"supervisorPhone,omitempty"`
	Assistant         *RecordRef              `xml:"assistant,omitempty"`
	AssistantPhone    string                  `xml:"assistantPhone,omitempty"`
	Comments          string                  `xml:"comments,omitempty"`
	Image             *RecordRef              `xml:"image,omitempty"`
	BillPay           *bool                   `xml:"billPay,omitempty"`
	DateCreated       *DateTime               `xml:"dateCreated,omitempty"`
	LastModifiedDate  *DateTime               `xml:"lastModifiedDate,omitempty"`
	AddressbookList   *ContactAddressbookList `xml:"addressbookList,omitempty"`
	SubscriptionsList *SubscriptionsList      `xml:"subscriptionsList,omitempty"`
	CustomFieldList   struct {
		CustomField CustomFields `xml:"customField,omitempty"`
	} `xml:"customFieldList,omitempty"`
	InternalID string `xml:"internalId,attr"`
	ExternalID string `xml:"externalId,attr"`
}
