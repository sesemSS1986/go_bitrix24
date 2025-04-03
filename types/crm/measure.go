package crm

type Measure struct {
	Id               int    `json:"ID"`
	Code             int    `json:"CODE"`
	MeasureTitle     string `json:"MEASURE_TITLE"`
	SymbolRus        string `json:"SYMBOL_RUS"`
	SymbolIntl       string `json:"SYMBOL_INTL"`
	SymbolLetterIntl string `json:"SYMBOL_LETTER_INTL"`
	IsDefault        string `json:"IS_DEFAULT"`
}
