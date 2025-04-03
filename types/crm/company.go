package crm

import (
	"time"
)

type Company struct {
	Id                    int       `json:"ID,string"`
	Title                 string    `json:"TITLE"`
	CompanyType           string    `json:"COMPANY_TYPE"`
	Logo                  int       `json:"LOGO"`
	Address               string    `json:"ADDRESS"`
	Address2              string    `json:"ADDRESS_2"`
	AddressCity           string    `json:"ADDRESS_CITY"`
	AddressPostalCode     string    `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion         string    `json:"ADDRESS_REGION"`
	AddressProvince       string    `json:"ADDRESS_PROVINCE"`
	AddressCountry        string    `json:"ADDRESS_COUNTRY"`
	AddressCountryCode    string    `json:"ADDRESS_COUNTRY_CODE"`
	AddressLegal          string    `json:"ADDRESS_LEGAL"`
	RegAddress            string    `json:"REG_ADDRESS"`
	RegAddress2           string    `json:"REG_ADDRESS_2"`
	RegAddressCity        string    `json:"REG_ADDRESS_CITY"`
	RegAddressPostalCode  string    `json:"REG_ADDRESS_POSTAL_CODE"`
	RegAddressRegion      string    `json:"REG_ADDRESS_REGION"`
	RegAddressProvince    string    `json:"REG_ADDRESS_PROVINCE"`
	RegAddressCountry     string    `json:"REG_ADDRESS_COUNTRY"`
	RegAddressCountryCode string    `json:"REG_ADDRESS_COUNTRY_CODE"`
	BankingDetails        string    `json:"BANKING_DETAILS"`
	Industry              string    `json:"INDUSTRY"`
	Employees             string    `json:"EMPLOYEES"`
	CurrencyId            string    `json:"CURRENCY_ID"`
	Revenue               float64   `json:"REVENUE"`
	Opened                string    `json:"OPENED"`
	Comments              string    `json:"COMMENTS"`
	HasPhone              string    `json:"HAS_PHONE"`
	HasEmail              string    `json:"HAS_EMAIL"`
	HasImol               string    `json:"HAS_IMOL"`
	IsMyCompany           string    `json:"IS_MY_COMPANY"`
	AssignedById          string    `json:"ASSIGNED_BY_ID"`
	CreatedById           string    `json:"CREATED_BY_ID"`
	ModifyById            string    `json:"MODIFY_BY_ID"`
	DateCreate            time.Time `json:"DATE_CREATE"`
	DateModify            time.Time `json:"DATE_MODIFY"`
	ContactId             string    `json:"CONTACT_ID"`
	LeadId                string    `json:"LEAD_ID"`
	OriginatorId          string    `json:"ORIGINATOR_ID"`
	OriginId              string    `json:"ORIGIN_ID"`
	OriginVersion         string    `json:"ORIGIN_VERSION"`
	UtmSource             string    `json:"UTM_SOURCE"`
	UtmMedium             string    `json:"UTM_MEDIUM"`
	UtmCampaign           string    `json:"UTM_CAMPAIGN"`
	UtmContent            string    `json:"UTM_CONTENT"`
	UtmTerm               string    `json:"UTM_TERM"`
	Phone                 string    `json:"PHONE"`
	Email                 string    `json:"EMAIL"`
	Web                   string    `json:"WEB"`
	Im                    string    `json:"IM"`
}
