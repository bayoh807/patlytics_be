package resource

type ReportReq struct {
	Company *string `form:"company" validate:"required,min=3,max=100"`
	Patent  *string `form:"patent" validate:"required,min=3,max=20"`
}

type ReportRes struct {
	Patent  string      `json:"patent_id"`
	Company string      `json:"company_name"`
	Date    string      `json:"analysis_date"`
	Analyze interface{} `json:"analyze"`
}

type SearchReq struct {
	Keyword string `form:"keyword" validate:"required,min=3,max=100"`
	Type    string `form:"type" validate:"required,oneof=company patent"`
}
