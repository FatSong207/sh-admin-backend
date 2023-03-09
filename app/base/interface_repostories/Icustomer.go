package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type ICustomerRepostory interface {
	core.IBaseRepostory[models.Customer, models.CustomerOutDto]
	GetByEmail(email string) (*models.Customer, error)
}
