package models

type Customer struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	Source   string `gorm:"source"`
	Phone    string `gorm:"phone"`
	Email    string `gorm:"email"`
	Industry string `gorm:"industry"`
	Level    string `gorm:"level"`
	Remarks  string `gorm:"remarks"`
	Region   string `gorm:"region"`
	Address  string `gorm:"address"`
	Status   int    `gorm:"status"`
	Creator  int64  `gorm:"creator"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

type CustomerOutDto struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	Source   string `gorm:"source"`
	Phone    string `gorm:"phone"`
	Email    string `gorm:"email"`
	Industry string `gorm:"industry"`
	Level    string `gorm:"level"`
	Remarks  string `gorm:"remarks"`
	Region   string `gorm:"region"`
	Address  string `gorm:"address"`
	Status   int    `gorm:"status"`
	//Creator  int64  `gorm:"creator"`
	Created int64 `gorm:"created"`
	//Updated  int64  `gorm:"updated"`
}

func (c Customer) TableName() string {
	return "customer"
}
