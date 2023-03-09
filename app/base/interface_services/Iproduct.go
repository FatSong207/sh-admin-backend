package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IProductService interface {
	core.IBaseService[models.Product, models.ProductOutDto]
	GetByCode(code string) (*models.Product, error)
}
