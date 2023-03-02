package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
	"SH-admin/models/common"
)

type IApiService interface {
	Core.IService[models.Api, models.ApiOutDto]
	GetAllApiTree() (result []models.ApiForTree, err error)
	FindWithPager(searchDto common.SearchDto[models.Api]) (*[]*models.ApiOutDto, int64, error)
}
