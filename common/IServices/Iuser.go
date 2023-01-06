package IServices

import (
	"SH-admin/common/Core"
	"SH-admin/models"
)

type IUserService interface {
	Core.IService[models.User, models.UserOutDto]
	Login(login *models.UserLoginReq) (*models.UserLoginRes, error)
}
