package models

type User struct {
	Id       int64  `gorm:"primaryKey" form:"id"`
	Email    string `gorm:"email" form:"email"`
	Password string `gorm:"password" form:"password"`
	Name     string `gorm:"name" form:"name"`
	Version  int    `gorm:"version" form:"version"`
	Expired  int64  `gorm:"expired" form:"expired"`
	Status   int    `gorm:"status" form:"status"`
	Created  int64  `gorm:"created" form:"created"`
	Updated  int64  `gorm:"updated" form:"updated"`
}

type UserOutDto struct {
	Id    int64  `gorm:"primaryKey"`
	Email string `gorm:"email"`
	//Password string `gorm:"password"`
	Name    string `gorm:"name"`
	Version int    `gorm:"version"`
	Expired int64  `gorm:"expired"`
	Status  int    `gorm:"status"`
	Created int64  `gorm:"created"`
	//Updated  int64  `gorm:"updated"`
}

// UserLoginReq 登入參數
type UserLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserLoginRes 登入返回結果
type UserLoginRes struct {
	Token string `json:"token"`
	User  User   `json:"user"`
	//ExpiresAt *jwt.NumericDate `json:"expiresAt"`
}

func (u User) TableName() string {
	return "user"
}
