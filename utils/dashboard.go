package utils

type Dashboard struct {
	UserCount  int `json:"userCount" form:"userCount"`
	LogCount   int `json:"logCount" form:"logCount"`
	VisitCount int `json:"visitCount" form:"visitCount"`
}
