package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type ICustomerService interface {
	Core.IService[models.Customer, models.CustomerOutDto]
	GetByEmail(email string) (*models.Customer, error)
}
