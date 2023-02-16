package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
	"SH-admin/models/common"
)

type ILogService interface {
	Core.IService[models.Log, models.LogOutDto]
	FindWithPager(searchDto common.SearchDto[models.Log]) (*[]*models.Log, int64, error)
	FindLoginlogWithPager(searchDto common.SearchDto[models.Log]) (*[]*models.LoginlogOutDto, int64, error)
}
