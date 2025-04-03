package crm

type Address struct {
	TypeId       int    `json:"TYPE_ID"`
	EntityTypeId int    `json:"ENTITY_TYPE_ID"`
	EntityId     int    `json:"ENTITY_ID"`
	Address1     string `json:"ADDRESS_1"`
	Address2     string `json:"ADDRESS_2"`
	City         string `json:"CITY"`
	PostalCode   string `json:"POSTAL_CODE"`
	Region       string `json:"REGION"`
	Province     string `json:"PROVINCE"`
	Country      string `json:"COUNTRY"`
	CountryCode  string `json:"COUNTRY_CODE"`
	AnchorTypeId int    `json:"ANCHOR_TYPE_ID"`
	AnchorId     int    `json:"ANCHOR_ID"`
}
