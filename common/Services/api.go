package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/models"
)

type ApiService struct {
	Core.IService[models.Api, models.ApiOutDto]
	apiRepo IRepostories.IApiRepostory
}

// NewApiService 供api層調用
func NewApiService() IServices.IApiService {
	ins := &ApiService{
		IService: Core.NewBaseService[models.Api, models.ApiOutDto](),
		apiRepo:  Repostories.NewApiRepostory(),
	}
	return ins
}
