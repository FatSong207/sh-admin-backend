package Repostories

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/global"
	"SH-admin/models"
)

type ProductRepostory struct {
	Core.BaseRepostory[models.Product, models.ProductOutDto]
}

// NewProductRepostory CTOR
func NewProductRepostory() IRepostories.IProductRepostory {
	ins := &ProductRepostory{}
	return ins
}

func (p *ProductRepostory) GetByCode(code string) (*models.Product, error) {
	t := new(models.Product)
	err := global.Db.Where("code=?", code).First(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
