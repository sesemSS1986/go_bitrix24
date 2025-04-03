package crm

import (
	"time"
)

type Quote struct {
	Id                 int       `json:"ID"`
	QuoteNumber        string    `json:"QUOTE_NUMBER"`
	Title              string    `json:"TITLE"`
	StatusId           string    `json:"STATUS_ID"`
	CurrencyId         string    `json:"CURRENCY_ID"`
	Opportunity        float64   `json:"OPPORTUNITY"`
	TaxValue           float64   `json:"TAX_VALUE"`
	ExchRate           float64   `json:"EXCH_RATE"`
	AccountCurrencyId  string    `json:"ACCOUNT_CURRENCY_ID"`
	OpportunityAccount float64   `json:"OPPORTUNITY_ACCOUNT"`
	TaxValueAccount    float64   `json:"TAX_VALUE_ACCOUNT"`
	CompanyId          string    `json:"COMPANY_ID"`
	MycompanyId        string    `json:"MYCOMPANY_ID"`
	ContactId          string    `json:"CONTACT_ID"`
	ContactIds         string    `json:"CONTACT_IDS"`
	Begindate          time.Time `json:"BEGINDATE"`
	Closedate          time.Time `json:"CLOSEDATE"`
	Opened             string    `json:"OPENED"`
	Closed             string    `json:"CLOSED"`
	Comments           string    `json:"COMMENTS"`
	Content            string    `json:"CONTENT"`
	Terms              string    `json:"TERMS"`
	ClientTitle        string    `json:"CLIENT_TITLE"`
	ClientAddr         string    `json:"CLIENT_ADDR"`
	ClientContact      string    `json:"CLIENT_CONTACT"`
	ClientEmail        string    `json:"CLIENT_EMAIL"`
	ClientPhone        string    `json:"CLIENT_PHONE"`
	ClientTpId         string    `json:"CLIENT_TP_ID"`
	ClientTpaId        string    `json:"CLIENT_TPA_ID"`
	AssignedById       string    `json:"ASSIGNED_BY_ID"`
	CreatedById        string    `json:"CREATED_BY_ID"`
	ModifyById         string    `json:"MODIFY_BY_ID"`
	DateCreate         time.Time `json:"DATE_CREATE"`
	DateModify         time.Time `json:"DATE_MODIFY"`
	LeadId             string    `json:"LEAD_ID"`
	DealId             string    `json:"DEAL_ID"`
	PersonTypeId       int       `json:"PERSON_TYPE_ID"`
	LocationId         string    `json:"LOCATION_ID"`
	UtmSource          string    `json:"UTM_SOURCE"`
	UtmMedium          string    `json:"UTM_MEDIUM"`
	UtmCampaign        string    `json:"UTM_CAMPAIGN"`
	UtmContent         string    `json:"UTM_CONTENT"`
	UtmTerm            string    `json:"UTM_TERM"`
}
