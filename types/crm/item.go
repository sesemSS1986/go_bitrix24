package crm

import "time"

type Item struct {
	Items []struct {
		ID                  int       `json:"ID"`
		XMLID               string    `json:"XMLID"`
		Title               string    `json:"TITLE"`
		CreatedBy           int       `json:"CREATED_BY"`
		UpdatedBy           int       `json:"UPDATED_BY"`
		MovedBy             int       `json:"MOVED_BY"`
		CreatedTime         time.Time `json:"CREATED_TIME"`
		UpdatedTime         time.Time `json:"UPDATED_TIME"`
		MovedTime           time.Time `json:"MOVED_TIME"`
		CategoryID          int       `json:"CATEGORY_ID"`
		Opened              string    `json:"OPENED"`
		StageID             string    `json:"STAGE_ID"`
		PreviousStageID     string    `json:"PREVIOUS_STAGE_ID"`
		Begindate           time.Time `json:"BEGINDATE"`
		Closedate           time.Time `json:"CLOSEDATE"`
		CompanyID           int       `json:"COMPANY_ID"`
		ContactID           int       `json:"CONTACT_ID"`
		Opportunity         int       `json:"OPPORTUNITY"`
		IsManualOpportunity string    `json:"IS_MANUAL_OPPORTUNITY"`
		TaxValue            int       `json:"TAX_VALUE"`
		CurrencyID          string    `json:"CURRENCY_ID"`
		OpportunityAccount  int       `json:"OPPORTUNITY_ACCOUNT"`
		TaxValueAccount     int       `json:"TAX_VALUE_ACCOUNT"`
		AccountCurrencyID   string    `json:"ACCOUNT_CURRENCY_ID"`
		MycompanyID         int       `json:"MYCOMPANY_ID"`
		SourceID            string    `json:"SOURCE_ID"`
		SourceDescription   string    `json:"SOURCE_DESCRIPTION"`
		WebformID           int       `json:"WEBFORM_ID"`
		//UfCrm51662608713    int         `json:"UF_CRM_5_1662608713"`
		//UfCrm51662608774    string      `json:"UF_CRM_5_1662608774"`
		//UfCrm51662608789    string      `json:"UF_CRM_5_1662608789"`
		//UfCrm51662608801    string      `json:"UF_CRM_5_1662608801"`
		//UfCrm51662608822    string      `json:"UF_CRM_5_1662608822"`
		//UfCrm51662608832    string      `json:"UF_CRM_5_1662608832"`
		//UfCrm51662608840    time.Time   `json:"UF_CRM_5_1662608840"`
		//UfCrm51662608850    int         `json:"UF_CRM_5_1662608850"`
		//UfCrm51662608860    int         `json:"UF_CRM_5_1662608860"`
		AssignedByID string      `json:"ASSIGNED_BY_ID"`
		UtmSource    interface{} `json:"UTM_SOURCE"`
		UtmMedium    interface{} `json:"UTM_MEDIUM"`
		UtmCampaign  interface{} `json:"UTM_CAMPAIGN"`
		UtmContent   interface{} `json:"UTM_CONTENT"`
		UtmTerm      interface{} `json:"UTM_TERM"`
		EntityTypeID int         `json:"ENTITY_TYPE_ID"`
		//Userfields   map[string]interface{} `json:"-"`
	} `json:"items"`
}
