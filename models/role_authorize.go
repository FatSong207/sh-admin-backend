package models

type RoleAuthorize struct {
	RoleId      int64 `json:"roleId"`
	AuthorizeId int64 `json:"-"`
}

func (RoleAuthorize) TableName() string {
	return "role_authorize"
}
