package crm

type Userfield struct {
	Id              string `json:"ID"`
	EntityId        string `json:"ENTITY_ID"`
	FieldName       string `json:"FIELD_NAME"`
	UserTypeId      string `json:"USER_TYPE_ID"`
	XmlId           string `json:"XML_ID"`
	Sort            string `json:"SORT"`
	Multiple        string `json:"MULTIPLE"`
	Mandatory       string `json:"MANDATORY"`
	ShowFilter      string `json:"SHOW_FILTER"`
	ShowInList      string `json:"SHOW_IN_LIST"`
	EditInList      string `json:"EDIT_IN_LIST"`
	IsSearchable    string `json:"IS_SEARCHABLE"`
	EditFormLabel   string `json:"EDIT_FORM_LABEL"`
	ListColumnLabel string `json:"LIST_COLUMN_LABEL"`
	ListFilterLabel string `json:"LIST_FILTER_LABEL"`
	ErrorMessage    string `json:"ERROR_MESSAGE"`
	HelpMessage     string `json:"HELP_MESSAGE"`
	List            string `json:"LIST"`
	Settings        string `json:"SETTINGS"`
}
