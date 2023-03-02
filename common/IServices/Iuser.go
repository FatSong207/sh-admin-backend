package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
	"SH-admin/models/common"
)

type IUserService interface {
	Core.IService[models.User, models.UserOutDto]
	Login(login *models.UserLoginReq) (*models.UserLoginRes, error)
	GetVerifyCode(email string) int
	Register(reg models.UserRegisterReq) int
	FindWithPager(searchDto common.SearchDto[models.User]) (*[]*models.UserOutDto, int64, error)
}
