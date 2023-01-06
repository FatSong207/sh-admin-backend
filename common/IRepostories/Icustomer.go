package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type ICustomerRepostory interface {
	Core.IRepostory[models.Customer, models.CustomerOutDto]
	GetByEmail(email string) (*models.Customer, error)
}
