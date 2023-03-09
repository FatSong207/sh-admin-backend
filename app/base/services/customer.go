package services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
	"fmt"
)

type CustomerService struct {
	core.IBaseService[models.Customer, models.CustomerOutDto]
	customerRepo interface_repostories.ICustomerRepostory
	productRepo  interface_repostories.IProductRepostory
}

// NewCustomerService CTOR
func NewCustomerService() *CustomerService {
	ins := &CustomerService{
		IBaseService: core.NewBaseService[models.Customer, models.CustomerOutDto](),
		customerRepo: repostories.NewCustomerRepostory(),
		productRepo:  repostories.NewProductRepostory(),
	}
	return ins
}

// GetByEmail 根據Email獲取Customer
func (c *CustomerService) GetByEmail(email string) (*models.Customer, error) {
	pro, _ := c.productRepo.GetByCode("005")
	fmt.Println("product:", pro)
	t, err := c.customerRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return t, nil
}
