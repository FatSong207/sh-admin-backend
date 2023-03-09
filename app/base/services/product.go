package services

import (
	core2 "SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
)

type ProductService struct {
	core2.IBaseService[models.Product, models.ProductOutDto]
	_productRepo interface_repostories.IProductRepostory
}

// NewProductService 供api層調用
func NewProductService() interface_services.IProductService {
	ins := &ProductService{
		IBaseService: core2.NewBaseService[models.Product, models.ProductOutDto](),
		_productRepo: repostories.NewProductRepostory(), //根據所需呼叫建構子(NewProduct2Repostory()也可以，因爲也實作介面了)
	}
	return ins
}

// GetByCode 根據code欄位獲取實體
func (p *ProductService) GetByCode(code string) (*models.Product, error) {
	t, err := p._productRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}
	return t, nil
}
