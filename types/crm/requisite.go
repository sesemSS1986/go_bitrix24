package crm

import (
	"time"
)

type Requisite struct {
	Id                 int       `json:"ID"`
	EntityTypeId       int       `json:"ENTITY_TYPE_ID"`
	EntityId           int       `json:"ENTITY_ID"`
	PresetId           int       `json:"PRESET_ID"`
	DateCreate         time.Time `json:"DATE_CREATE"`
	DateModify         time.Time `json:"DATE_MODIFY"`
	CreatedById        string    `json:"CREATED_BY_ID"`
	ModifyById         string    `json:"MODIFY_BY_ID"`
	Name               string    `json:"NAME"`
	Code               string    `json:"CODE"`
	XmlId              string    `json:"XML_ID"`
	OriginatorId       string    `json:"ORIGINATOR_ID"`
	Active             string    `json:"ACTIVE"`
	Sort               int       `json:"SORT"`
	RqName             string    `json:"RQ_NAME"`
	RqFirstName        string    `json:"RQ_FIRST_NAME"`
	RqLastName         string    `json:"RQ_LAST_NAME"`
	RqSecondName       string    `json:"RQ_SECOND_NAME"`
	RqCompanyName      string    `json:"RQ_COMPANY_NAME"`
	RqCompanyFullName  string    `json:"RQ_COMPANY_FULL_NAME"`
	RqCompanyRegDate   string    `json:"RQ_COMPANY_REG_DATE"`
	RqDirector         string    `json:"RQ_DIRECTOR"`
	RqAccountant       string    `json:"RQ_ACCOUNTANT"`
	RqCeoName          string    `json:"RQ_CEO_NAME"`
	RqCeoWorkPos       string    `json:"RQ_CEO_WORK_POS"`
	RqContact          string    `json:"RQ_CONTACT"`
	RqEmail            string    `json:"RQ_EMAIL"`
	RqPhone            string    `json:"RQ_PHONE"`
	RqFax              string    `json:"RQ_FAX"`
	RqIdentDoc         string    `json:"RQ_IDENT_DOC"`
	RqIdentDocSer      string    `json:"RQ_IDENT_DOC_SER"`
	RqIdentDocNum      string    `json:"RQ_IDENT_DOC_NUM"`
	RqIdentDocPersNum  string    `json:"RQ_IDENT_DOC_PERS_NUM"`
	RqIdentDocDate     string    `json:"RQ_IDENT_DOC_DATE"`
	RqIdentDocIssuedBy string    `json:"RQ_IDENT_DOC_ISSUED_BY"`
	RqIdentDocDepCode  string    `json:"RQ_IDENT_DOC_DEP_CODE"`
	RqInn              string    `json:"RQ_INN"`
	RqKpp              string    `json:"RQ_KPP"`
	RqUsrle            string    `json:"RQ_USRLE"`
	RqIfns             string    `json:"RQ_IFNS"`
	RqOgrn             string    `json:"RQ_OGRN"`
	RqOgrnip           string    `json:"RQ_OGRNIP"`
	RqOkpo             string    `json:"RQ_OKPO"`
	RqOktmo            string    `json:"RQ_OKTMO"`
	RqOkved            string    `json:"RQ_OKVED"`
	RqEdrpou           string    `json:"RQ_EDRPOU"`
	RqDrfo             string    `json:"RQ_DRFO"`
	RqKbe              string    `json:"RQ_KBE"`
	RqIin              string    `json:"RQ_IIN"`
	RqBin              string    `json:"RQ_BIN"`
	RqStCertSer        string    `json:"RQ_ST_CERT_SER"`
	RqStCertNum        string    `json:"RQ_ST_CERT_NUM"`
	RqStCertDate       string    `json:"RQ_ST_CERT_DATE"`
	RqVatPayer         string    `json:"RQ_VAT_PAYER"`
	RqVatId            string    `json:"RQ_VAT_ID"`
	RqVatCertSer       string    `json:"RQ_VAT_CERT_SER"`
	RqVatCertNum       string    `json:"RQ_VAT_CERT_NUM"`
	RqVatCertDate      string    `json:"RQ_VAT_CERT_DATE"`
	RqResidenceCountry string    `json:"RQ_RESIDENCE_COUNTRY"`
	RqBaseDoc          string    `json:"RQ_BASE_DOC"`
}
