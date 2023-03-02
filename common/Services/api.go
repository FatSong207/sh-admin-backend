package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	"SH-admin/models/common"
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

func (a *ApiService) FindWithPager(searchDto common.SearchDto[models.Api]) (*[]*models.ApiOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.ApiOutDto, 0)
	var bind = make([]*models.Api, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	//t := new(T)
	db := global.Db.Model(query)
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	if query.ApiGroup != "" {
		db = db.Where("api_group like ?", "%"+query.ApiGroup+"%")
	}

	t, err := a.apiRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	return &dest, t, nil
}
