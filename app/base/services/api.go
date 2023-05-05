package services

import (
	core "SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"SH-admin/global"
)

type ApiService struct {
	core.IBaseService[models.Api, models.ApiOutDto]
	_apiRepo interface_repostories.IApiRepostory
}

// NewApiService 供api層調用
func NewApiService() interface_services.IApiService {
	ins := &ApiService{
		IBaseService: core.NewBaseService[models.Api, models.ApiOutDto](),
		_apiRepo:     repostories.NewApiRepostory(),
	}
	return ins
}

func (a *ApiService) GetAllApiTree() (result []models.ApiForTree, err error) {
	all, err := a._apiRepo.GetAll()
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

func (a *ApiService) FindWithPager(searchDto common.SearchDto[models.Api]) (*[]*models.ApiOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.ApiOutDto, 0)
	var bind = make([]*models.Api, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	//t := new(T)
	db := global.DB().Model(query)
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	if query.ApiGroup != "" {
		db = db.Where("api_group like ?", "%"+query.ApiGroup+"%")
	}

	t, err := a._apiRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	return &dest, t, nil
}
