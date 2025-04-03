package crm

import (
	"time"
)

type Dealcategory struct {
	Id          int       `json:"ID"`
	CreatedDate time.Time `json:"CREATED_DATE"`
	Name        string    `json:"NAME"`
	IsLocked    string    `json:"IS_LOCKED"`
	Sort        int       `json:"SORT"`
}
