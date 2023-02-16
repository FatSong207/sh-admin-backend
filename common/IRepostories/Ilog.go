package IRepostories

import (
	"SH-admin/common/Core"
	"SH-admin/models"
	"SH-admin/models/common"
	"gorm.io/gorm"
)

type ILogRepostory interface {
	Core.IRepostory[models.Log, models.LogOutDto]
	FindLoginlogWithPager(searchDto common.PageInfo, db *gorm.DB, order string, dest *[]*models.LoginlogOutDto, bind *[]*models.LoginlogOutDto) (int64, error)
}
