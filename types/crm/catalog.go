package crm

type Catalog struct {
	Id           int    `json:"ID"`
	Name         string `json:"NAME"`
	OriginatorId string `json:"ORIGINATOR_ID"`
	OriginId     string `json:"ORIGIN_ID"`
	XmlId        string `json:"XML_ID"`
}
