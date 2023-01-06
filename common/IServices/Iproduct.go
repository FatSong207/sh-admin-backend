package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IProductService interface {
	Core.IService[models.Product, models.ProductOutDto]
	GetByCode(code string) (*models.Product, error)
}
