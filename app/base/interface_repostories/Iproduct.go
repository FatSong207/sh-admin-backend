package interface_repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
)

type IProductRepostory interface {
	core.IBaseRepostory[models.Product, models.ProductOutDto]
	GetByCode(code string) (*models.Product, error)
}
