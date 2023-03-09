package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type ICustomerService interface {
	core.IBaseService[models.Customer, models.CustomerOutDto]
	GetByEmail(email string) (*models.Customer, error)
}
