package Repostories

import (
	"SH-admin/common/IRepostories"
	"SH-admin/global"
	"SH-admin/models"
)

type Product2Repostory struct {
	//Core.BaseRepostory[models.Product]
}

// NewProduct2Repostory CTOR
func NewProduct2Repostory() IRepostories.IProductRepostory {
	ins := &ProductRepostory{}
	return ins
}

func (p *Product2Repostory) GetByCode(code string) *models.Product {
	t := new(models.Product)
	err := global.Db.Where("code=?", code).First(t).Error
	if err != nil {
		panic(err)
	}
	return t
}
