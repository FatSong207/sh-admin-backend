package common

type PageInfo struct {
	PageNum  int `json:"pageNum" form:"pageNum"`   //第幾頁
	PageSize int `json:"pageSize" form:"pageSize"` //每頁大小
}
