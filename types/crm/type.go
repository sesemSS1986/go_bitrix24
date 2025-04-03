package crm

type Type struct {
	Types []struct {
		ID                        int    `json:"id"`
		Title                     string `json:"title"`
		Code                      string `json:"code"`
		CreatedBy                 int    `json:"createdBy"`
		EntityTypeID              int    `json:"entityTypeId"`
		IsCategoriesEnabled       string `json:"isCategoriesEnabled"`
		IsStagesEnabled           string `json:"isStagesEnabled"`
		IsBeginCloseDatesEnabled  string `json:"isBeginCloseDatesEnabled"`
		IsClientEnabled           string `json:"isClientEnabled"`
		IsUseInUserfieldEnabled   string `json:"isUseInUserfieldEnabled"`
		IsLinkWithProductsEnabled string `json:"isLinkWithProductsEnabled"`
		IsMycompanyEnabled        string `json:"isMycompanyEnabled"`
		IsDocumentsEnabled        string `json:"isDocumentsEnabled"`
		IsSourceEnabled           string `json:"isSourceEnabled"`
		IsObserversEnabled        string `json:"isObserversEnabled"`
		IsRecyclebinEnabled       string `json:"isRecyclebinEnabled"`
		IsAutomationEnabled       string `json:"isAutomationEnabled"`
		IsBizProcEnabled          string `json:"isBizProcEnabled"`
		IsSetOpenPermissions      string `json:"isSetOpenPermissions"`
		IsPaymentsEnabled         string `json:"isPaymentsEnabled"`
	} `json:"types"`
}

//
//type Types struct {
//	ID                        int    `json:"ID"`
//	Title                     string `json:"TITLE"`
//	Code                      string `json:"CODE"`
//	CreatedBy                 int    `json:"CREATED_BY"`
//	EntityTypeID              int    `json:"ENTITY_TYPE_ID"`
//	IsCategoriesEnabled       string `json:"IS_CATEGORIES_ENABLED"`
//	IsStagesEnabled           string `json:"IS_STAGES_ENABLED"`
//	IsBeginCloseDatesEnabled  string `json:"IS_BEGIN_CLOSE_DATES_ENABLED"`
//	IsClientEnabled           string `json:"IS_CLIENT_ENABLED"`
//	IsUseInUserfieldEnabled   string `json:"IS_USE_IN_USERFIELD_ENABLED"`
//	IsLinkWithProductsEnabled string `json:"IS_LINK_WITH_PRODUCTS_ENABLED"`
//	IsMycompanyEnabled        string `json:"IS_MYCOMPANY_ENABLED"`
//	IsDocumentsEnabled        string `json:"IS_DOCUMENTS_ENABLED"`
//	IsSourceEnabled           string `json:"IS_SOURCE_ENABLED"`
//	IsObserversEnabled        string `json:"IS_OBSERVERS_ENABLED"`
//	IsRecyclebinEnabled       string `json:"IS_RECYCLEBIN_ENABLED"`
//	IsAutomationEnabled       string `json:"IS_AUTOMATION_ENABLED"`
//	IsBizProcEnabled          string `json:"IS_BIZPROC_ENABLED"`
//	IsSetOpenPermissions      string `json:"IS_SET_OPEN_PERMISSIONS"`
//	IsPaymentsEnabled         string `json:"IS_PAYMENTS_ENABLED"`
//}

//type crmTypeGetStruct struct {
//	Result struct {
//		Type struct {
//			Relations struct {
//				Parent []struct {
//					EntityTypeID          int    `json:"entityTypeId"`
//					IsChildrenListEnabled string `json:"isChildrenListEnabled"`
//					IsPredefined          string `json:"isPredefined"`
//				} `json:"parent"`
//				Child []struct {
//					EntityTypeID          int    `json:"entityTypeId"`
//					IsChildrenListEnabled string `json:"isChildrenListEnabled"`
//					IsPredefined          string `json:"isPredefined"`
//				} `json:"child"`
//			} `json:"relations"`
//			LinkedUserFields struct {
//				CALENDAREVENTUFCRMCALEVENT string `json:"CALENDAR_EVENT|UF_CRM_CAL_EVENT"`
//				TASKSTASKUFCRMTASK         string `json:"TASKS_TASK|UF_CRM_TASK"`
//				TASKSTASKTEMPLATEUFCRMTASK string `json:"TASKS_TASK_TEMPLATE|UF_CRM_TASK"`
//			} `json:"linkedUserFields"`
//			CustomSections []struct {
//				ID         int    `json:"id"`
//				Title      string `json:"title"`
//				IsSelected string `json:"isSelected"`
//			} `json:"customSections"`
//			CustomSectionID int `json:"customSectionId"`
//		} `json:"type"`
//	}
//}

//type CrmTypeGetStruct struct {
//	Types []struct {
//		ID           int    `json:"ID,string"`
//		Title        string `json:"TITLE"`
//		EntityTypeID int    `json:"ENTITY_TYPE_ID"`
//	} `json:"types"`
//}
