package crm

type RequisiteLink struct {
	EntityTypeId   int `json:"ENTITY_TYPE_ID"`
	EntityId       int `json:"ENTITY_ID"`
	RequisiteId    int `json:"REQUISITE_ID"`
	BankDetailId   int `json:"BANK_DETAIL_ID"`
	McRequisiteId  int `json:"MC_REQUISITE_ID"`
	McBankDetailId int `json:"MC_BANK_DETAIL_ID"`
}
