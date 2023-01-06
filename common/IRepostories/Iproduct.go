package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IProductRepostory interface {
	Core.IRepostory[models.Product, models.ProductOutDto]
	GetByCode(code string) (*models.Product, error)
}
