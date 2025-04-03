package crm

type Paysystem struct {
	Id           int    `json:"ID"`
	Name         string `json:"NAME"`
	Active       string `json:"ACTIVE"`
	Sort         int    `json:"SORT"`
	Description  string `json:"DESCRIPTION"`
	PersonTypeId int    `json:"PERSON_TYPE_ID"`
	ActionFile   string `json:"ACTION_FILE"`
	Handler      string `json:"HANDLER"`
	HandlerCode  string `json:"HANDLER_CODE"`
	HandlerName  string `json:"HANDLER_NAME"`
}
