package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/Repostories"
	"SH-admin/models"
	"fmt"
)

type CustomerService struct {
	Core.IService[models.Customer, models.CustomerOutDto]
	customerRepo IRepostories.ICustomerRepostory
	productRepo  IRepostories.IProductRepostory
}

// NewCustomerService CTOR
func NewCustomerService() *CustomerService {
	ins := &CustomerService{
		IService:     Core.NewBaseService[models.Customer, models.CustomerOutDto](),
		customerRepo: Repostories.NewCustomerRepostory(),
		productRepo:  Repostories.NewProductRepostory(),
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
