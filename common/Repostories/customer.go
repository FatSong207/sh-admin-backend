package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/global"
	"SH-admin/models"
)

type CustomerRepostory struct {
	Core.BaseRepostory[models.Customer, models.CustomerOutDto]
}

// NewCustomerRepostory CTOR
func NewCustomerRepostory() IRepostories.ICustomerRepostory {
	ins := &CustomerRepostory{}
	return ins
}

// GetByEmail 根據Email獲取客戶
func (c *CustomerRepostory) GetByEmail(email string) (*models.Customer, error) {
	t := new(models.Customer)
	err := global.Db.Where("email like ?", "%"+email+"%").First(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
