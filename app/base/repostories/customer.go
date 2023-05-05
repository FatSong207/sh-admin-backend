package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
	"SH-admin/global"
)

type CustomerRepostory struct {
	core.BaseRepostory[models.Customer, models.CustomerOutDto]
}

// NewCustomerRepostory CTOR
func NewCustomerRepostory() interface_repostories.ICustomerRepostory {
	ins := &CustomerRepostory{}
	return ins
}

// GetByEmail 根據Email獲取客戶
func (c *CustomerRepostory) GetByEmail(email string) (*models.Customer, error) {
	t := new(models.Customer)
	err := global.DB().Where("email like ?", "%"+email+"%").First(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
