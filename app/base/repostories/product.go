package repostories

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
	"SH-admin/global"
)

type ProductRepostory struct {
	core.BaseRepostory[models.Product, models.ProductOutDto]
}

// NewProductRepostory CTOR
func NewProductRepostory() interface_repostories.IProductRepostory {
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
