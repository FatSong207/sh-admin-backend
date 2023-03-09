package interface_services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
)

type IUserService interface {
	core.IBaseService[models.User, models.UserOutDto]
	Login(login *models.UserLoginReq) (*models.UserLoginRes, error)
	GetVerifyCode(email string) int
	Register(reg models.UserRegisterReq) int
	FindWithPager(searchDto common.SearchDto[models.User]) (*[]*models.UserOutDto, int64, error)
}
