package crm

import (
	"time"
)

type Product struct {
	Id              int       `json:"ID"`
	CatalogId       int       `json:"CATALOG_ID"`
	Price           float64   `json:"PRICE"`
	CurrencyId      string    `json:"CURRENCY_ID"`
	Name            string    `json:"NAME"`
	Code            string    `json:"CODE"`
	Description     string    `json:"DESCRIPTION"`
	DescriptionType string    `json:"DESCRIPTION_TYPE"`
	Active          string    `json:"ACTIVE"`
	SectionId       int       `json:"SECTION_ID"`
	Sort            int       `json:"SORT"`
	VatId           int       `json:"VAT_ID"`
	VatIncluded     string    `json:"VAT_INCLUDED"`
	Measure         int       `json:"MEASURE"`
	XmlId           string    `json:"XML_ID"`
	PreviewPicture  string    `json:"PREVIEW_PICTURE"`
	DetailPicture   string    `json:"DETAIL_PICTURE"`
	DateCreate      time.Time `json:"DATE_CREATE"`
	TimestampX      time.Time `json:"TIMESTAMP_X"`
	ModifiedBy      int       `json:"MODIFIED_BY"`
	CreatedBy       int       `json:"CREATED_BY"`
	Property768     string    `json:"PROPERTY_768"`
	Property769     string    `json:"PROPERTY_769"`
	Property770     string    `json:"PROPERTY_770"`
	Property771     string    `json:"PROPERTY_771"`
}
