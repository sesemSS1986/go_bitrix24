package crm

type Productsection struct {
	Id        int    `json:"ID"`
	CatalogId int    `json:"CATALOG_ID"`
	SectionId int    `json:"SECTION_ID"`
	Name      string `json:"NAME"`
	XmlId     string `json:"XML_ID"`
}
