package models

type CasbinInfo struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}

type UpdateCasbinParam struct {
	RoleId     string        `json:"roleId" form:"roleId"`
	CasbinInfo []*CasbinInfo `json:"casbinInfo" form:"casbinInfo"`
}
