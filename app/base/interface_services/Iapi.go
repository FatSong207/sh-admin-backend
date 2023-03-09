package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
)

type IApiService interface {
	core.IBaseService[models.Api, models.ApiOutDto]
	GetAllApiTree() (result []models.ApiForTree, err error)
	FindWithPager(searchDto common.SearchDto[models.Api]) (*[]*models.ApiOutDto, int64, error)
}
