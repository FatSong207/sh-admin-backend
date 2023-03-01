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

func (a *ApiService) GetAllApiTree() (result []models.ApiForTree, err error) {
	all, err := a.apiRepo.GetAll()
	if err != nil {
		return nil, err
	}
	m := make(map[string][]models.Api)
	for _, api := range all {
		m[api.ApiGroup] = append(m[api.ApiGroup], api)
	}
	for k, apis := range m {
		at := models.ApiForTree{
			Description: k,
			Path:        k,
			Id:          k,
			Children:    apis,
		}
		result = append(result, at)
	}
	return result, nil
}
