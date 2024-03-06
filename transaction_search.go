package netsuite

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
	"github.com/omniboost/go-netsuite-soap/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewTransactionSearchRequest() TransactionSearchRequest {
	r := TransactionSearchRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TransactionSearchRequest struct {
	client      *Client
	queryParams *TransactionSearchRequestQueryParams
	pathParams  *TransactionSearchRequestPathParams
	method      string
	headers     http.Header
	requestBody TransactionSearchRequestBody
}

func (r TransactionSearchRequest) SOAPAction() string {
	return "search"
}

func (r TransactionSearchRequest) NewQueryParams() *TransactionSearchRequestQueryParams {
	return &TransactionSearchRequestQueryParams{}
}

type TransactionSearchRequestQueryParams struct {
}

func (p TransactionSearchRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TransactionSearchRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r TransactionSearchRequest) NewPathParams() *TransactionSearchRequestPathParams {
	return &TransactionSearchRequestPathParams{}
}

type TransactionSearchRequestPathParams struct {
}

func (p *TransactionSearchRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TransactionSearchRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *TransactionSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *TransactionSearchRequest) Method() string {
	return r.method
}

func (r TransactionSearchRequest) NewRequestBody() TransactionSearchRequestBody {
	return TransactionSearchRequestBody{
		SearchRecord: SearchRecordBasic{
			Type: "tranSales:TransactionSearch",
		},
	}
}

type TransactionSearchRequestBody struct {
	XMLName xml.Name `xml:"platformMsgs:search"`

	SearchRecord SearchRecordBasic `xml:"platformMsgs:searchRecord"`
}

func (r *TransactionSearchRequest) RequestBody() *TransactionSearchRequestBody {
	return &r.requestBody
}

func (r *TransactionSearchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *TransactionSearchRequest) SetRequestBody(body TransactionSearchRequestBody) {
	r.requestBody = body
}

func (r *TransactionSearchRequest) NewResponseBody() *TransactionSearchRequestResponseBody {
	return &TransactionSearchRequestResponseBody{}
}

type TransactionSearchRequestResponseBody struct {
	XMLName xml.Name `xml:"searchResponse"`

	SearchResult struct {
		PlatformCore string `xml:"platformCore,attr"`
		TotalRecords string `xml:"totalRecords"`
		PageSize     string `xml:"pageSize"`
		TotalPages   string `xml:"totalPages"`
		PageIndex    string `xml:"pageIndex"`
		SearchId     string `xml:"searchId"`
		RecordList   struct {
			Record Transactions `xml:"record"`
		} `xml:"recordList"`
	} `xml:"searchResult"`
}

func (r *TransactionSearchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("", r.PathParams())
	return &u, err
}

func (r *TransactionSearchRequest) Do() (TransactionSearchRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}

type TransactionSearchBasic struct {
	Type                         SearchEnumMultiSelectField `xml:"type,omitempty"`
	Account                      SearchMultiSelectField     `xml:"platformCore:account,omitempty"`
	AccountType                  SearchEnumMultiSelectField `xml:"platformCore:accountType,omitempty"`
	AcctCorpCardExp              SearchMultiSelectField     `xml:"platformCore:acctCorpCardExp,omitempty"`
	ActualProductionEndDate      SearchDateField            `xml:"platformCore:actualProductionEndDate,omitempty"`
	ActualProductionStartDate    SearchDateField            `xml:"platformCore:actualProductionStartDate,omitempty"`
	ActualShipDate               SearchDateField            `xml:"platformCore:actualShipDate,omitempty"`
	AltSalesAmount               SearchDoubleField          `xml:"platformCore:altSalesAmount,omitempty"`
	AltSalesNetAmount            SearchDoubleField          `xml:"platformCore:altSalesNetAmount,omitempty"`
	Amount                       SearchDoubleField          `xml:"platformCore:amount,omitempty"`
	AmountPaid                   SearchDoubleField          `xml:"platformCore:amountPaid,omitempty"`
	AmountRemaining              SearchDoubleField          `xml:"platformCore:amountRemaining,omitempty"`
	AmountUnbilled               SearchDoubleField          `xml:"platformCore:amountUnbilled,omitempty"`
	AnyLineItem                  SearchMultiSelectField     `xml:"platformCore:anyLineItem,omitempty"`
	AppliedToForeignAmount       SearchDoubleField          `xml:"platformCore:appliedToForeignAmount,omitempty"`
	AppliedToIsFxVariance        SearchBooleanField         `xml:"platformCore:appliedToIsFxVariance,omitempty"`
	AppliedToLinkAmount          SearchDoubleField          `xml:"platformCore:appliedToLinkAmount,omitempty"`
	AppliedToLinkType            SearchEnumMultiSelectField `xml:"platformCore:appliedToLinkType,omitempty"`
	AppliedToTransaction         SearchMultiSelectField     `xml:"platformCore:appliedToTransaction,omitempty"`
	ApplyingForeignAmount        SearchDoubleField          `xml:"platformCore:applyingForeignAmount,omitempty"`
	ApplyingIsFxVariance         SearchBooleanField         `xml:"platformCore:applyingIsFxVariance,omitempty"`
	ApplyingLinkAmount           SearchDoubleField          `xml:"platformCore:applyingLinkAmount,omitempty"`
	ApplyingLinkType             SearchEnumMultiSelectField `xml:"platformCore:applyingLinkType,omitempty"`
	ApplyingTransaction          SearchMultiSelectField     `xml:"platformCore:applyingTransaction,omitempty"`
	ApprovalStatus               SearchEnumMultiSelectField `xml:"platformCore:approvalStatus,omitempty"`
	AuthCode                     SearchStringField          `xml:"platformCore:authCode,omitempty"`
	AutoCalculateLag             SearchBooleanField         `xml:"platformCore:autoCalculateLag,omitempty"`
	AvsStreetMatch               SearchEnumMultiSelectField `xml:"platformCore:avsStreetMatch,omitempty"`
	AvsZipMatch                  SearchEnumMultiSelectField `xml:"platformCore:avsZipMatch,omitempty"`
	Billable                     SearchBooleanField         `xml:"platformCore:billable,omitempty"`
	BillAddress                  SearchStringField          `xml:"platformCore:billAddress,omitempty"`
	BillAddressee                SearchStringField          `xml:"platformCore:billAddressee,omitempty"`
	BillAttention                SearchStringField          `xml:"platformCore:billAttention,omitempty"`
	BillCity                     SearchStringField          `xml:"platformCore:billCity,omitempty"`
	BillCountry                  SearchEnumMultiSelectField `xml:"platformCore:billCountry,omitempty"`
	BillCounty                   SearchStringField          `xml:"platformCore:billCounty,omitempty"`
	BilledDate                   SearchDateField            `xml:"platformCore:billedDate,omitempty"`
	BillingAccount               SearchMultiSelectField     `xml:"platformCore:billingAccount,omitempty"`
	BillingSchedule              SearchMultiSelectField     `xml:"platformCore:billingSchedule,omitempty"`
	BillingStatus                SearchBooleanField         `xml:"platformCore:billingStatus,omitempty"`
	BillingTransaction           SearchMultiSelectField     `xml:"platformCore:billingTransaction,omitempty"`
	BillPhone                    SearchStringField          `xml:"platformCore:billPhone,omitempty"`
	BillState                    SearchStringField          `xml:"platformCore:billState,omitempty"`
	BillVarianceStatus           SearchEnumMultiSelectField `xml:"platformCore:billVarianceStatus,omitempty"`
	BillZip                      SearchStringField          `xml:"platformCore:billZip,omitempty"`
	BinNumber                    SearchStringField          `xml:"platformCore:binNumber,omitempty"`
	BinNumberQuantity            SearchDoubleField          `xml:"platformCore:binNumberQuantity,omitempty"`
	BomQuantity                  SearchDoubleField          `xml:"platformCore:bomQuantity,omitempty"`
	BookSpecificTransaction      SearchBooleanField         `xml:"platformCore:bookSpecificTransaction,omitempty"`
	BuildEntireAssembly          SearchBooleanField         `xml:"platformCore:buildEntireAssembly,omitempty"`
	BuildVariance                SearchDoubleField          `xml:"platformCore:buildVariance,omitempty"`
	Built                        SearchDoubleField          `xml:"platformCore:built,omitempty"`
	CanHaveStackablePromotions   SearchBooleanField         `xml:"platformCore:canHaveStackablePromotions,omitempty"`
	CatchUpPeriod                SearchMultiSelectField     `xml:"platformCore:catchUpPeriod,omitempty"`
	CcCustomerCode               SearchStringField          `xml:"platformCore:ccCustomerCode,omitempty"`
	CcExpireDate                 SearchDateField            `xml:"platformCore:ccExpireDate,omitempty"`
	CcName                       SearchStringField          `xml:"platformCore:ccName,omitempty"`
	CcNumber                     SearchStringField          `xml:"platformCore:ccNumber,omitempty"`
	ChargeType                   SearchEnumMultiSelectField `xml:"platformCore:chargeType,omitempty"`
	Class                        SearchMultiSelectField     `xml:"platformCore:class,omitempty"`
	Cleared                      SearchBooleanField         `xml:"platformCore:cleared,omitempty"`
	Closed                       SearchBooleanField         `xml:"platformCore:closed,omitempty"`
	CloseDate                    SearchDateField            `xml:"platformCore:closeDate,omitempty"`
	Cogs                         SearchBooleanField         `xml:"platformCore:cogs,omitempty"`
	CommissionEffectiveDate      SearchDateField            `xml:"platformCore:commissionEffectiveDate,omitempty"`
	Commit                       SearchEnumMultiSelectField `xml:"platformCore:commit,omitempty"`
	ComponentYield               SearchDoubleField          `xml:"platformCore:componentYield,omitempty"`
	ConfirmationNumber           SearchStringField          `xml:"platformCore:confirmationNumber,omitempty"`
	Contribution                 SearchLongField            `xml:"platformCore:contribution,omitempty"`
	CostComponentAmount          SearchDoubleField          `xml:"platformCore:costComponentAmount,omitempty"`
	CostComponentCategory        SearchMultiSelectField     `xml:"platformCore:costComponentCategory,omitempty"`
	CostComponentItem            SearchMultiSelectField     `xml:"platformCore:costComponentItem,omitempty"`
	CostComponentQuantity        SearchDoubleField          `xml:"platformCore:costComponentQuantity,omitempty"`
	CostComponentStandardCost    SearchDoubleField          `xml:"platformCore:costComponentStandardCost,omitempty"`
	CostEstimate                 SearchDoubleField          `xml:"platformCore:costEstimate,omitempty"`
	CostEstimateRate             SearchDoubleField          `xml:"platformCore:costEstimateRate,omitempty"`
	CostEstimateType             SearchEnumMultiSelectField `xml:"platformCore:costEstimateType,omitempty"`
	CreatedBy                    SearchMultiSelectField     `xml:"platformCore:createdBy,omitempty"`
	CreatedFrom                  SearchMultiSelectField     `xml:"platformCore:createdFrom,omitempty"`
	CreditAmount                 SearchDoubleField          `xml:"platformCore:creditAmount,omitempty"`
	CscMatch                     SearchEnumMultiSelectField `xml:"platformCore:cscMatch,omitempty"`
	Currency                     SearchMultiSelectField     `xml:"platformCore:currency,omitempty"`
	CustomerSubOf                SearchMultiSelectField     `xml:"platformCore:customerSubOf,omitempty"`
	CustomForm                   SearchMultiSelectField     `xml:"platformCore:customForm,omitempty"`
	CustomGL                     SearchBooleanField         `xml:"platformCore:customGL,omitempty"`
	CustType                     SearchMultiSelectField     `xml:"platformCore:custType,omitempty"`
	DateCreated                  SearchDateField            `xml:"platformCommon:dateCreated,omitempty"`
	DaysOpen                     SearchLongField            `xml:"platformCore:daysOpen,omitempty"`
	DaysOverdue                  SearchLongField            `xml:"platformCore:daysOverdue,omitempty"`
	DebitAmount                  SearchDoubleField          `xml:"platformCore:debitAmount,omitempty"`
	DeferredRevenue              SearchDoubleField          `xml:"platformCore:deferredRevenue,omitempty"`
	DeferRevRec                  SearchBooleanField         `xml:"platformCore:deferRevRec,omitempty"`
	Department                   SearchMultiSelectField     `xml:"platformCore:department,omitempty"`
	DepositDate                  SearchDateField            `xml:"platformCore:depositDate,omitempty"`
	DepositTransaction           SearchMultiSelectField     `xml:"platformCore:depositTransaction,omitempty"`
	DrAccount                    SearchMultiSelectField     `xml:"platformCore:drAccount,omitempty"`
	DueDate                      SearchDateField            `xml:"platformCore:dueDate,omitempty"`
	Email                        SearchStringField          `xml:"platformCore:email,omitempty"`
	Employee                     SearchMultiSelectField     `xml:"platformCore:employee,omitempty"`
	EndDate                      SearchDateField            `xml:"platformCore:endDate,omitempty"`
	Entity                       SearchMultiSelectField     `xml:"platformCore:entity,omitempty"`
	EntityStatus                 SearchMultiSelectField     `xml:"platformCore:entityStatus,omitempty"`
	EstGrossProfit               SearchDoubleField          `xml:"platformCore:estGrossProfit,omitempty"`
	EstGrossProfitPct            SearchDoubleField          `xml:"platformCore:estGrossProfitPct,omitempty"`
	ExchangeRate                 SearchDoubleField          `xml:"platformCore:exchangeRate,omitempty"`
	ExcludeCommission            SearchBooleanField         `xml:"platformCore:excludeCommission,omitempty"`
	ExcludeFromRateRequest       SearchBooleanField         `xml:"platformCore:excludeFromRateRequest,omitempty"`
	ExpectedCloseDate            SearchDateField            `xml:"platformCore:expectedCloseDate,omitempty"`
	ExpectedReceiptDate          SearchDateField            `xml:"platformCore:expectedReceiptDate,omitempty"`
	ExpenseCategory              SearchMultiSelectField     `xml:"platformCore:expenseCategory,omitempty"`
	ExpenseDate                  SearchDateField            `xml:"platformCore:expenseDate,omitempty"`
	ExternalId                   SearchMultiSelectField     `xml:"platformCommon:externalId,omitempty"`
	ExternalIdString             SearchStringField          `xml:"platformCore:externalIdString,omitempty"`
	FinChrg                      SearchBooleanField         `xml:"platformCore:finChrg,omitempty"`
	Firmed                       SearchBooleanField         `xml:"platformCore:firmed,omitempty"`
	ForecastType                 SearchEnumMultiSelectField `xml:"platformCore:forecastType,omitempty"`
	FulfillingTransaction        SearchMultiSelectField     `xml:"platformCore:fulfillingTransaction,omitempty"`
	FxAccount                    SearchMultiSelectField     `xml:"platformCore:fxAccount,omitempty"`
	FxAmount                     SearchDoubleField          `xml:"platformCore:fxAmount,omitempty"`
	FxCostEstimate               SearchDoubleField          `xml:"platformCore:fxCostEstimate,omitempty"`
	FxCostEstimateRate           SearchDoubleField          `xml:"platformCore:fxCostEstimateRate,omitempty"`
	FxEstGrossProfit             SearchDoubleField          `xml:"platformCore:fxEstGrossProfit,omitempty"`
	FxTranCostEstimate           SearchDoubleField          `xml:"platformCore:fxTranCostEstimate,omitempty"`
	FxVsoeAllocation             SearchDoubleField          `xml:"platformCore:fxVsoeAllocation,omitempty"`
	FxVsoeAmount                 SearchDoubleField          `xml:"platformCore:fxVsoeAmount,omitempty"`
	FxVsoePrice                  SearchDoubleField          `xml:"platformCore:fxVsoePrice,omitempty"`
	GcoAvailabelToCharge         SearchBooleanField         `xml:"platformCore:gcoAvailabelToCharge,omitempty"`
	GcoAvailableToRefund         SearchBooleanField         `xml:"platformCore:gcoAvailableToRefund,omitempty"`
	GcoAvsStreetMatch            SearchEnumMultiSelectField `xml:"platformCore:gcoAvsStreetMatch,omitempty"`
	GcoAvsZipMatch               SearchEnumMultiSelectField `xml:"platformCore:gcoAvsZipMatch,omitempty"`
	GcoBuyerAccountAge           SearchLongField            `xml:"platformCore:gcoBuyerAccountAge,omitempty"`
	GcoBuyerIp                   SearchStringField          `xml:"platformCore:gcoBuyerIp,omitempty"`
	GcoChargeAmount              SearchDoubleField          `xml:"platformCore:gcoChargeAmount,omitempty"`
	GcoChargebackAmount          SearchDoubleField          `xml:"platformCore:gcoChargebackAmount,omitempty"`
	GcoConfirmedChargedTotal     SearchDoubleField          `xml:"platformCore:gcoConfirmedChargedTotal,omitempty"`
	GcoConfirmedRefundedTotal    SearchDoubleField          `xml:"platformCore:gcoConfirmedRefundedTotal,omitempty"`
	GcoCreditcardNumber          SearchStringField          `xml:"platformCore:gcoCreditcardNumber,omitempty"`
	GcoCscMatch                  SearchEnumMultiSelectField `xml:"platformCore:gcoCscMatch,omitempty"`
	GcoFinancialState            SearchStringField          `xml:"platformCore:gcoFinancialState,omitempty"`
	GcoFulfillmentState          SearchStringField          `xml:"platformCore:gcoFulfillmentState,omitempty"`
	GcoOrderId                   SearchStringField          `xml:"platformCore:gcoOrderId,omitempty"`
	GcoOrderTotal                SearchDoubleField          `xml:"platformCore:gcoOrderTotal,omitempty"`
	GcoPromotionAmount           SearchDoubleField          `xml:"platformCore:gcoPromotionAmount,omitempty"`
	GcoPromotionName             SearchStringField          `xml:"platformCore:gcoPromotionName,omitempty"`
	GcoRefundAmount              SearchDoubleField          `xml:"platformCore:gcoRefundAmount,omitempty"`
	GcoShippingTotal             SearchDoubleField          `xml:"platformCore:gcoShippingTotal,omitempty"`
	GcoStateChangedDetail        SearchStringField          `xml:"platformCore:gcoStateChangedDetail,omitempty"`
	GiftCertificate              SearchStringField          `xml:"platformCore:giftCertificate,omitempty"`
	GrossAmount                  SearchDoubleField          `xml:"platformCore:grossAmount,omitempty"`
	IncludeInForecast            SearchBooleanField         `xml:"platformCore:includeInForecast,omitempty"`
	Incoterm                     SearchMultiSelectField     `xml:"platformCore:incoterm,omitempty"`
	IntercoStatus                SearchEnumMultiSelectField `xml:"platformCore:intercoStatus,omitempty"`
	IntercoTransaction           SearchMultiSelectField     `xml:"platformCore:intercoTransaction,omitempty"`
	InternalID                   SearchMultiSelectField     `xml:"platformCommon:internalId,omitempty"`
	InternalIdNumber             SearchLongField            `xml:"platformCore:internalIdNumber,omitempty"`
	InventoryLocation            SearchMultiSelectField     `xml:"platformCore:inventoryLocation,omitempty"`
	InventorySubsidiary          SearchMultiSelectField     `xml:"platformCore:inventorySubsidiary,omitempty"`
	InVsoeBundle                 SearchBooleanField         `xml:"platformCore:inVsoeBundle,omitempty"`
	IsAllocation                 SearchBooleanField         `xml:"platformCore:isAllocation,omitempty"`
	IsBackflush                  SearchBooleanField         `xml:"platformCore:isBackflush,omitempty"`
	IsGcoChargeback              SearchBooleanField         `xml:"platformCore:isGcoChargeback,omitempty"`
	IsGcoChargeConfirmed         SearchBooleanField         `xml:"platformCore:isGcoChargeConfirmed,omitempty"`
	IsGcoPaymentGuaranteed       SearchBooleanField         `xml:"platformCore:isGcoPaymentGuaranteed,omitempty"`
	IsGcoRefundConfirmed         SearchBooleanField         `xml:"platformCore:isGcoRefundConfirmed,omitempty"`
	IsInsideDelivery             SearchBooleanField         `xml:"platformCore:isInsideDelivery,omitempty"`
	IsInsidePickup               SearchBooleanField         `xml:"platformCore:isInsidePickup,omitempty"`
	IsIntercompanyAdjustment     SearchBooleanField         `xml:"platformCore:isIntercompanyAdjustment,omitempty"`
	IsInTransitPayment           SearchBooleanField         `xml:"platformCore:isInTransitPayment,omitempty"`
	IsMultiShipTo                SearchBooleanField         `xml:"platformCore:isMultiShipTo,omitempty"`
	IsPayPalMeth                 SearchBooleanField         `xml:"platformCore:isPayPalMeth,omitempty"`
	IsReversal                   SearchBooleanField         `xml:"platformCore:isReversal,omitempty"`
	IsRevRecTransaction          SearchBooleanField         `xml:"platformCore:isRevRecTransaction,omitempty"`
	IsScrap                      SearchBooleanField         `xml:"platformCore:isScrap,omitempty"`
	IsShipAddress                SearchBooleanField         `xml:"platformCore:isShipAddress,omitempty"`
	IsTransferPriceCosting       SearchBooleanField         `xml:"platformCore:isTransferPriceCosting,omitempty"`
	IsVsoeAlloc                  SearchBooleanField         `xml:"platformCore:isVsoeAlloc,omitempty"`
	IsWip                        SearchBooleanField         `xml:"platformCore:isWip,omitempty"`
	Item                         SearchMultiSelectField     `xml:"platformCore:item,omitempty"`
	ItemFulfillmentChoice        SearchEnumMultiSelectField `xml:"platformCore:itemFulfillmentChoice,omitempty"`
	ItemRevision                 SearchMultiSelectField     `xml:"platformCore:itemRevision,omitempty"`
	ItemSubOf                    SearchMultiSelectField     `xml:"platformCore:itemSubOf,omitempty"`
	LandedCostPerLine            SearchBooleanField         `xml:"platformCore:landedCostPerLine,omitempty"`
	LastModifiedDate             SearchDateField            `xml:"platformCore:lastModifiedDate,omitempty"`
	LeadSource                   SearchMultiSelectField     `xml:"platformCore:leadSource,omitempty"`
	LineUniqueKey                SearchLongField            `xml:"platformCore:lineUniqueKey,omitempty"`
	Location                     SearchMultiSelectField     `xml:"platformCore:location,omitempty"`
	LocationAutoAssigned         SearchBooleanField         `xml:"platformCore:locationAutoAssigned,omitempty"`
	MainLine                     SearchBooleanField         `xml:"platformCore:mainLine,omitempty"`
	MainName                     SearchMultiSelectField     `xml:"platformCore:mainName,omitempty"`
	ManufacturingRouting         SearchMultiSelectField     `xml:"platformCore:manufacturingRouting,omitempty"`
	MatchBillToReceipt           SearchBooleanField         `xml:"platformCore:matchBillToReceipt,omitempty"`
	Memo                         SearchStringField          `xml:"platformCommon:memo,omitempty"`
	MemoMain                     SearchStringField          `xml:"platformCommon:memoMain,omitempty"`
	Memorized                    SearchBooleanField         `xml:"platformCore:memorized,omitempty"`
	MerchantAccount              SearchStringField          `xml:"platformCore:merchantAccount,omitempty"`
	Message                      SearchStringField          `xml:"platformCore:message,omitempty"`
	MultiSubsidiary              SearchBooleanField         `xml:"platformCore:multiSubsidiary,omitempty"`
	Name                         SearchStringField          `xml:"name,omitempty"`
	NameText                     SearchStringField          `xml:"platformCore:nameText,omitempty"`
	NetAmount                    SearchDoubleField          `xml:"platformCore:netAmount,omitempty"`
	NextApprover                 SearchMultiSelectField     `xml:"platformCore:nextApprover,omitempty"`
	NextBillDate                 SearchDateField            `xml:"platformCore:nextBillDate,omitempty"`
	Nexus                        SearchMultiSelectField     `xml:"platformCore:nexus,omitempty"`
	NoAutoAssignLocation         SearchBooleanField         `xml:"platformCore:noAutoAssignLocation,omitempty"`
	NonReimbursable              SearchBooleanField         `xml:"platformCore:nonReimbursable,omitempty"`
	Number                       SearchLongField            `xml:"platformCommon:number,omitempty"`
	OneTimeTotal                 SearchDoubleField          `xml:"platformCore:oneTimeTotal,omitempty"`
	Opportunity                  SearchMultiSelectField     `xml:"platformCore:opportunity,omitempty"`
	OrderAllocationStrategy      SearchMultiSelectField     `xml:"platformCore:orderAllocationStrategy,omitempty"`
	OrderPriority                SearchDoubleField          `xml:"platformCore:orderPriority,omitempty"`
	OtherRefNum                  SearchTextNumberField      `xml:"platformCore:otherRefNum,omitempty"`
	OtherRefNumNonDeposit        SearchTextNumberField      `xml:"platformCore:otherRefNumNonDeposit,omitempty"`
	OverheadParentItem           SearchMultiSelectField     `xml:"platformCore:overheadParentItem,omitempty"`
	OverrideInstallments         SearchBooleanField         `xml:"platformCore:overrideInstallments,omitempty"`
	PackageCount                 SearchLongField            `xml:"platformCore:packageCount,omitempty"`
	PaidTransaction              SearchMultiSelectField     `xml:"platformCore:paidTransaction,omitempty"`
	Parent                       SearchLongField            `xml:"platformCore:parent,omitempty"`
	Partner                      SearchMultiSelectField     `xml:"platformCore:partner,omitempty"`
	PartnerContribution          SearchLongField            `xml:"platformCore:partnerContribution,omitempty"`
	PartnerRole                  SearchMultiSelectField     `xml:"platformCore:partnerRole,omitempty"`
	PartnerTeamMember            SearchMultiSelectField     `xml:"platformCore:partnerTeamMember,omitempty"`
	PayingTransaction            SearchMultiSelectField     `xml:"platformCore:payingTransaction,omitempty"`
	PaymentApproved              SearchBooleanField         `xml:"platformCore:paymentApproved,omitempty"`
	PaymentEventDate             SearchDateField            `xml:"platformCore:paymentEventDate,omitempty"`
	PaymentEventHoldReason       SearchEnumMultiSelectField `xml:"platformCore:paymentEventHoldReason,omitempty"`
	PaymentEventPurchaseCardUsed SearchBooleanField         `xml:"platformCore:paymentEventPurchaseCardUsed,omitempty"`
	PaymentEventPurchaseDataSent SearchBooleanField         `xml:"platformCore:paymentEventPurchaseDataSent,omitempty"`
	PaymentEventResult           SearchEnumMultiSelectField `xml:"platformCore:paymentEventResult,omitempty"`
	PaymentEventType             SearchEnumMultiSelectField `xml:"platformCore:paymentEventType,omitempty"`
	PaymentHold                  SearchBooleanField         `xml:"platformCore:paymentHold,omitempty"`
	PaymentMethod                SearchMultiSelectField     `xml:"platformCore:paymentMethod,omitempty"`
	PaymentOption                SearchStringField          `xml:"platformCore:paymentOption,omitempty"`
	PayPalPending                SearchBooleanField         `xml:"platformCore:payPalPending,omitempty"`
	PayPalStatus                 SearchStringField          `xml:"platformCore:payPalStatus,omitempty"`
	PayPalTranId                 SearchStringField          `xml:"platformCore:payPalTranId,omitempty"`
	PnRefNum                     SearchStringField          `xml:"platformCore:pnRefNum,omitempty"`
	PoAsText                     SearchStringField          `xml:"platformCore:poAsText,omitempty"`
	PolicyViolated               SearchBooleanField         `xml:"platformCore:policyViolated,omitempty"`
	Posting                      SearchBooleanField         `xml:"platformCore:posting,omitempty"`
	PostingPeriod                RecordRef                  `xml:"platformCore:postingPeriod,omitempty"`
	// PostingPeriodRelative        PostingPeriodDate          `xml:"platformCore:postingPeriodRelative,omitempty"`
	PriceLevel                 SearchMultiSelectField     `xml:"platformCore:priceLevel,omitempty"`
	PrintedPickingTicket       SearchBooleanField         `xml:"platformCore:printedPickingTicket,omitempty"`
	Probability                SearchLongField            `xml:"platformCore:probability,omitempty"`
	ProjectedAmount            SearchDoubleField          `xml:"platformCore:projectedAmount,omitempty"`
	ProjectTask                SearchMultiSelectField     `xml:"platformCore:projectTask,omitempty"`
	PromoCode                  SearchMultiSelectField     `xml:"platformCore:promoCode,omitempty"`
	PurchaseOrder              SearchMultiSelectField     `xml:"platformCore:purchaseOrder,omitempty"`
	Quantity                   SearchDoubleField          `xml:"platformCore:quantity,omitempty"`
	QuantityBilled             SearchDoubleField          `xml:"platformCore:quantityBilled,omitempty"`
	QuantityCommitted          SearchDoubleField          `xml:"platformCore:quantityCommitted,omitempty"`
	QuantityPacked             SearchDoubleField          `xml:"platformCore:quantityPacked,omitempty"`
	QuantityPicked             SearchDoubleField          `xml:"platformCore:quantityPicked,omitempty"`
	QuantityRevCommitted       SearchDoubleField          `xml:"platformCore:quantityRevCommitted,omitempty"`
	QuantityShipRecv           SearchDoubleField          `xml:"platformCore:quantityShipRecv,omitempty"`
	RecognizedRevenue          SearchDoubleField          `xml:"platformCore:recognizedRevenue,omitempty"`
	RecordType                 SearchStringField          `xml:"platformCore:recordType,omitempty"`
	RecurAnnuallyTotal         SearchDoubleField          `xml:"platformCore:recurAnnuallyTotal,omitempty"`
	RecurMonthlyTotal          SearchDoubleField          `xml:"platformCore:recurMonthlyTotal,omitempty"`
	RecurQuarterlyTotal        SearchDoubleField          `xml:"platformCore:recurQuarterlyTotal,omitempty"`
	RecurringBill              SearchBooleanField         `xml:"platformCore:recurringBill,omitempty"`
	RecurWeeklyTotal           SearchDoubleField          `xml:"platformCore:recurWeeklyTotal,omitempty"`
	RefNumber                  SearchLongField            `xml:"platformCore:refNumber,omitempty"`
	RequestedDate              SearchDateField            `xml:"platformCore:requestedDate,omitempty"`
	RevCommitStatus            SearchEnumMultiSelectField `xml:"platformCore:revCommitStatus,omitempty"`
	RevCommittingStatus        SearchBooleanField         `xml:"platformCore:revCommittingStatus,omitempty"`
	RevCommittingTransaction   SearchMultiSelectField     `xml:"platformCore:revCommittingTransaction,omitempty"`
	RevenueStatus              SearchEnumMultiSelectField `xml:"platformCore:revenueStatus,omitempty"`
	ReversalDate               SearchDateField            `xml:"platformCore:reversalDate,omitempty"`
	ReversalNumber             SearchStringField          `xml:"platformCore:reversalNumber,omitempty"`
	RevRecEndDate              SearchDateField            `xml:"platformCore:revRecEndDate,omitempty"`
	RevRecOnRevCommitment      SearchBooleanField         `xml:"platformCore:revRecOnRevCommitment,omitempty"`
	RevRecStartDate            SearchDateField            `xml:"platformCore:revRecStartDate,omitempty"`
	RevRecTermInMonths         SearchLongField            `xml:"platformCore:revRecTermInMonths,omitempty"`
	SalesEffectiveDate         SearchDateField            `xml:"platformCore:salesEffectiveDate,omitempty"`
	SalesOrder                 SearchMultiSelectField     `xml:"platformCore:salesOrder,omitempty"`
	SalesRep                   SearchMultiSelectField     `xml:"platformCore:salesRep,omitempty"`
	SalesTeamMember            SearchMultiSelectField     `xml:"platformCore:salesTeamMember,omitempty"`
	SalesTeamRole              SearchMultiSelectField     `xml:"platformCore:salesTeamRole,omitempty"`
	SchedulingMethod           SearchEnumMultiSelectField `xml:"platformCore:schedulingMethod,omitempty"`
	SerialNumber               SearchStringField          `xml:"platformCore:serialNumber,omitempty"`
	SerialNumberCost           SearchDoubleField          `xml:"platformCore:serialNumberCost,omitempty"`
	SerialNumberCostAdjustment SearchDoubleField          `xml:"platformCore:serialNumberCostAdjustment,omitempty"`
	SerialNumberQuantity       SearchDoubleField          `xml:"platformCore:serialNumberQuantity,omitempty"`
	SerialNumbers              SearchStringField          `xml:"platformCore:serialNumbers,omitempty"`
	ShipAddress                SearchStringField          `xml:"platformCore:shipAddress,omitempty"`
	ShipAddressee              SearchStringField          `xml:"platformCore:shipAddressee,omitempty"`
	ShipAttention              SearchStringField          `xml:"platformCore:shipAttention,omitempty"`
	ShipCarrier                SearchEnumMultiSelectField `xml:"platformCore:shipCarrier,omitempty"`
	ShipCity                   SearchStringField          `xml:"platformCore:shipCity,omitempty"`
	ShipComplete               SearchBooleanField         `xml:"platformCore:shipComplete,omitempty"`
	ShipCountry                SearchEnumMultiSelectField `xml:"platformCore:shipCountry,omitempty"`
	ShipCounty                 SearchStringField          `xml:"platformCore:shipCounty,omitempty"`
	ShipDate                   SearchDateField            `xml:"platformCore:shipDate,omitempty"`
	ShipGroup                  SearchLongField            `xml:"platformCore:shipGroup,omitempty"`
	ShipMethod                 SearchMultiSelectField     `xml:"platformCore:shipMethod,omitempty"`
	ShipPhone                  SearchStringField          `xml:"platformCore:shipPhone,omitempty"`
	Shipping                   SearchBooleanField         `xml:"platformCore:shipping,omitempty"`
	ShipRecvStatus             SearchBooleanField         `xml:"platformCore:shipRecvStatus,omitempty"`
	ShipRecvStatusLine         SearchBooleanField         `xml:"platformCore:shipRecvStatusLine,omitempty"`
	ShipState                  SearchMultiSelectField     `xml:"platformCore:shipState,omitempty"`
	ShipTo                     SearchMultiSelectField     `xml:"platformCore:shipTo,omitempty"`
	ShipZip                    SearchStringField          `xml:"platformCore:shipZip,omitempty"`
	Source                     SearchEnumMultiSelectField `xml:"platformCore:source,omitempty"`
	StartDate                  SearchDateField            `xml:"platformCore:startDate,omitempty"`
	Statistical                SearchBooleanField         `xml:"platformCore:statistical,omitempty"`
	Status                     SearchEnumMultiSelectField `xml:"platformCore:status,omitempty"`
	Subscription               SearchMultiSelectField     `xml:"platformCore:subscription,omitempty"`
	SubscriptionLine           SearchMultiSelectField     `xml:"platformCore:subscriptionLine,omitempty"`
	Subsidiary                 SearchMultiSelectField     `xml:"platformCommon:subsidiary,omitempty"`
	SubsidiaryTaxRegNum        SearchMultiSelectField     `xml:"platformCore:subsidiaryTaxRegNum,omitempty"`
	TaxItem                    SearchMultiSelectField     `xml:"platformCore:taxItem,omitempty"`
	TaxLine                    SearchBooleanField         `xml:"platformCore:taxLine,omitempty"`
	// taxPeriod RecordRef `xml:"platformCore:taxPeriod,omitempty"`
	// taxPeriodRelative PostingPeriodDate `xml:"platformCore:taxPeriodRelative,omitempty"`
	TaxPointDate                   SearchDateField            `xml:"platformCore:taxPointDate,omitempty"`
	TaxRate                        SearchDoubleField          `xml:"platformCore:taxRate,omitempty"`
	Terms                          SearchMultiSelectField     `xml:"platformCore:terms,omitempty"`
	TermsOfSale                    SearchEnumMultiSelectField `xml:"platformCore:termsOfSale,omitempty"`
	Title                          SearchStringField          `xml:"platformCore:title,omitempty"`
	ToBeEmailed                    SearchBooleanField         `xml:"platformCore:toBeEmailed,omitempty"`
	ToBePrinted                    SearchBooleanField         `xml:"platformCore:toBePrinted,omitempty"`
	ToSubsidiary                   SearchMultiSelectField     `xml:"platformCore:toSubsidiary,omitempty"`
	TotalAmount                    SearchDoubleField          `xml:"platformCore:totalAmount,omitempty"`
	TrackingNumbers                SearchStringField          `xml:"platformCore:trackingNumbers,omitempty"`
	TranCostEstimate               SearchDoubleField          `xml:"platformCore:tranCostEstimate,omitempty"`
	TranDate                       SearchDateField            `xml:"platformCommon:tranDate,omitempty"`
	TranEstGrossProfit             SearchDoubleField          `xml:"platformCore:tranEstGrossProfit,omitempty"`
	TranEstGrossProfitPct          SearchDoubleField          `xml:"platformCore:tranEstGrossProfitPct,omitempty"`
	TranFxEstGrossProfit           SearchDoubleField          `xml:"platformCore:tranFxEstGrossProfit,omitempty"`
	TranID                         SearchStringField          `xml:"platformCommon:tranId,omitempty"`
	TranIsVsoeBundle               SearchBooleanField         `xml:"platformCore:tranIsVsoeBundle,omitempty"`
	TransactionDiscount            SearchBooleanField         `xml:"platformCore:transactionDiscount,omitempty"`
	TransactionLineType            SearchEnumMultiSelectField `xml:"platformCore:transactionLineType,omitempty"`
	TransactionNumber              SearchStringField          `xml:"platformCore:transactionNumber,omitempty"`
	TransferLocation               SearchMultiSelectField     `xml:"platformCore:transferLocation,omitempty"`
	TransferOrderQuantityCommitted SearchDoubleField          `xml:"platformCore:transferOrderQuantityCommitted,omitempty"`
	TransferOrderQuantityPacked    SearchDoubleField          `xml:"platformCore:transferOrderQuantityPacked,omitempty"`
	TransferOrderQuantityPicked    SearchDoubleField          `xml:"platformCore:transferOrderQuantityPicked,omitempty"`
	TransferOrderQuantityReceived  SearchDoubleField          `xml:"platformCore:transferOrderQuantityReceived,omitempty"`
	TransferOrderQuantityShipped   SearchDoubleField          `xml:"platformCore:transferOrderQuantityShipped,omitempty"`
	// type SearchEnumMultiSelectField `xml:"platformCore:type,omitempty"`
	Unit               SearchMultiSelectField     `xml:"platformCore:unit,omitempty"`
	UnitCostOverride   SearchDoubleField          `xml:"platformCore:unitCostOverride,omitempty"`
	UnitsType          SearchMultiSelectField     `xml:"platformCore:unitsType,omitempty"`
	VendType           SearchMultiSelectField     `xml:"platformCore:vendType,omitempty"`
	VisibleToCustomer  SearchBooleanField         `xml:"platformCore:visibleToCustomer,omitempty"`
	Voided             SearchBooleanField         `xml:"platformCore:voided,omitempty"`
	VsoeAllocation     SearchDoubleField          `xml:"platformCore:vsoeAllocation,omitempty"`
	VsoeAmount         SearchDoubleField          `xml:"platformCore:vsoeAmount,omitempty"`
	VsoeDeferral       SearchEnumMultiSelectField `xml:"platformCore:vsoeDeferral,omitempty"`
	VsoeDelivered      SearchBooleanField         `xml:"platformCore:vsoeDelivered,omitempty"`
	VsoePermitDiscount SearchEnumMultiSelectField `xml:"platformCore:vsoePermitDiscount,omitempty"`
	VsoePrice          SearchDoubleField          `xml:"platformCore:vsoePrice,omitempty"`
	WebSite            SearchMultiSelectField     `xml:"platformCore:webSite,omitempty"`
	CustomFieldList    SearchCustomFieldList      `xml:"platformCommon:customFieldList,omitempty"`
}

func (c TransactionSearchBasic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c TransactionSearchBasic) IsEmpty() bool {
	return zero.IsZero(c)
}
