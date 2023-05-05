package repostories

import (
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/models"
	"SH-admin/global"
)

type Product2Repostory struct {
	//core.BaseRepostory[models.Product]
}

// NewProduct2Repostory CTOR
func NewProduct2Repostory() interface_repostories.IProductRepostory {
	ins := &ProductRepostory{}
	return ins
}

func (p *Product2Repostory) GetByCode(code string) *models.Product {
	t := new(models.Product)
	err := global.DB().Where("code=?", code).First(t).Error
	if err != nil {
		panic(err)
	}
	return t
}
