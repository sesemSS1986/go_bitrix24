package crm

import (
	"time"
)

type Currency struct {
	Currency     string    `json:"CURRENCY"`
	AmountCnt    string    `json:"AMOUNT_CNT"`
	Amount       float64   `json:"AMOUNT"`
	Base         string    `json:"BASE"`
	Sort         string    `json:"SORT"`
	DateUpdate   time.Time `json:"DATE_UPDATE"`
	Lid          string    `json:"LID"`
	FormatString string    `json:"FORMAT_STRING"`
	FullName     string    `json:"FULL_NAME"`
	DecPoint     string    `json:"DEC_POINT"`
	ThousandsSep string    `json:"THOUSANDS_SEP"`
	Decimals     string    `json:"DECIMALS"`
	Lang         string    `json:"LANG"`
}
