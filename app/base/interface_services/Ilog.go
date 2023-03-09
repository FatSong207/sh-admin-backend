package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
)

type ILogService interface {
	core.IBaseService[models.Log, models.LogOutDto]
	FindWithPager(searchDto common.SearchDto[models.Log]) (*[]*models.LogOutDto, int64, error)
	FindLoginlogWithPager(searchDto common.SearchDto[models.Log]) (*[]*models.LoginlogOutDto, int64, error)
}
