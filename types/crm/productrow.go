package crm

type Productrow struct {
	Id             int     `json:"ID"`
	OwnerId        int     `json:"OWNER_ID"`
	OwnerType      string  `json:"OWNER_TYPE"`
	ProductId      int     `json:"PRODUCT_ID"`
	ProductName    string  `json:"PRODUCT_NAME"`
	Price          float64 `json:"PRICE"`
	PriceExclusive float64 `json:"PRICE_EXCLUSIVE"`
	PriceNetto     float64 `json:"PRICE_NETTO"`
	PriceBrutto    float64 `json:"PRICE_BRUTTO"`
	Quantity       float64 `json:"QUANTITY"`
	DiscountTypeId int     `json:"DISCOUNT_TYPE_ID"`
	DiscountRate   float64 `json:"DISCOUNT_RATE"`
	DiscountSum    float64 `json:"DISCOUNT_SUM"`
	TaxRate        float64 `json:"TAX_RATE"`
	TaxIncluded    string  `json:"TAX_INCLUDED"`
	Customized     string  `json:"CUSTOMIZED"`
	MeasureCode    int     `json:"MEASURE_CODE"`
	MeasureName    string  `json:"MEASURE_NAME"`
	Sort           int     `json:"SORT"`
}
