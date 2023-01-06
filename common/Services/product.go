package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/models"
)

type ProductService struct {
	Core.IService[models.Product, models.ProductOutDto]
	productRepo IRepostories.IProductRepostory
}

// NewProductService 供api層調用
func NewProductService() IServices.IProductService {
	ins := &ProductService{
		IService:    Core.NewBaseService[models.Product, models.ProductOutDto](),
		productRepo: Repostories.NewProductRepostory(), //根據所需呼叫建構子(NewProduct2Repostory()也可以，因爲也實作介面了)
	}
	return ins
}

// GetByCode 根據code欄位獲取實體
func (p *ProductService) GetByCode(code string) (*models.Product, error) {
	t, err := p.productRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}
	return t, nil
}
