package crm

import "time"

type CrmMultifield struct {
	ID        string `json:"ID"`
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
	TypeID    string `json:"TYPE_ID"`
}

type photo struct {
	ID          int    `json:"id"`
	ShowURL     string `json:"showUrl"`
	DownloadURL string `json:"downloadUrl"`
}

type Contact struct {
	Id                int             `json:"ID,string"`
	Honorific         string          `json:"HONORIFIC"`
	Name              string          `json:"NAME"`
	SecondName        string          `json:"SECOND_NAME"`
	LastName          string          `json:"LAST_NAME"`
	FullName          string          `json:"FULL_NAME"`
	Photo             photo           `json:"PHOTO,omitempty"`
	Birthdate         string          `json:"BIRTHDATE"`
	TypeId            string          `json:"TYPE_ID"`
	SourceId          string          `json:"SOURCE_ID"`
	SourceDescription string          `json:"SOURCE_DESCRIPTION"`
	Post              string          `json:"POST"`
	Address           string          `json:"ADDRESS"`
	Address2          string          `json:"ADDRESS_2"`
	AddressCity       string          `json:"ADDRESS_CITY"`
	AddressPostalCode string          `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion     string          `json:"ADDRESS_REGION"`
	AddressProvince   string          `json:"ADDRESS_PROVINCE"`
	AddressCountry    string          `json:"ADDRESS_COUNTRY"`
	Comments          string          `json:"COMMENTS"`
	Opened            string          `json:"OPENED"`
	Export            string          `json:"EXPORT"`
	HasPhone          string          `json:"HAS_PHONE"`
	HasEmail          string          `json:"HAS_EMAIL"`
	HasImol           string          `json:"HAS_IMOL"`
	AssignedById      int             `json:"ASSIGNED_BY_ID,string"`
	CreatedById       int             `json:"CREATED_BY_ID,string"`
	ModifyById        int             `json:"MODIFY_BY_ID,string"`
	DateCreate        time.Time       `json:"DATE_CREATE"`
	DateModify        time.Time       `json:"DATE_MODIFY"`
	CompanyId         string          `json:"COMPANY_ID"`
	LeadId            int             `json:"LEAD_ID,string"`
	OriginatorId      string          `json:"ORIGINATOR_ID"`
	OriginId          string          `json:"ORIGIN_ID"`
	OriginVersion     string          `json:"ORIGIN_VERSION"`
	FaceId            int             `json:"FACE_ID,string"`
	UtmSource         string          `json:"UTM_SOURCE"`
	UtmMedium         string          `json:"UTM_MEDIUM"`
	UtmCampaign       string          `json:"UTM_CAMPAIGN"`
	UtmContent        string          `json:"UTM_CONTENT"`
	UtmTerm           string          `json:"UTM_TERM"`
	Phone             []CrmMultifield `json:"PHONE,omitempty"`
	Email             []CrmMultifield `json:"EMAIL,omitempty"`
	Web               []CrmMultifield `json:"WEB,omitempty"`

	Userfield map[string]interface{}
}
