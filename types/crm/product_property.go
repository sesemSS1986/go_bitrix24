package crm

type ProductProperty struct {
	Id               int    `json:"ID"`
	IblockId         int    `json:"IBLOCK_ID"`
	XmlId            string `json:"XML_ID"`
	Name             string `json:"NAME"`
	Active           string `json:"ACTIVE"`
	IsRequired       string `json:"IS_REQUIRED"`
	Sort             int    `json:"SORT"`
	PropertyType     string `json:"PROPERTY_TYPE"`
	Multiple         string `json:"MULTIPLE"`
	DefaultValue     string `json:"DEFAULT_VALUE"`
	RowCount         int    `json:"ROW_COUNT"`
	ColCount         int    `json:"COL_COUNT"`
	FileType         string `json:"FILE_TYPE"`
	LinkIblockId     int    `json:"LINK_IBLOCK_ID"`
	UserType         string `json:"USER_TYPE"`
	UserTypeSettings string `json:"USER_TYPE_SETTINGS"`
	Values           string `json:"VALUES"`
}
