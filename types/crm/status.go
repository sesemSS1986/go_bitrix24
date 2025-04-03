package crm

type Status struct {
	Id       int    `json:"ID"`
	EntityId string `json:"ENTITY_ID"`
	StatusId string `json:"STATUS_ID"`
	Sort     int    `json:"SORT"`
	Name     string `json:"NAME"`
	NameInit string `json:"NAME_INIT"`
	System   string `json:"SYSTEM"`
	Extra    string `json:"EXTRA"`
}
