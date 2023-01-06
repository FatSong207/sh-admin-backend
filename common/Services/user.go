package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/models"
	"SH-admin/utils"
	"errors"
)

type UserService struct {
	Core.IService[models.User, models.UserOutDto]
	userRepo IRepostories.IUserRepostory
}

// NewUserService 供api層調用
func NewUserService() IServices.IUserService {
	ins := &UserService{
		IService: Core.NewBaseService[models.User, models.UserOutDto](),
		userRepo: Repostories.NewUserRepostory(),
	}
	return ins
}

func (u *UserService) Login(login *models.UserLoginReq) (*models.UserLoginRes, error) {
	user := &models.User{
		Email: login.Email,
	}
	whereStruct, err := u.userRepo.GetByWhereStruct(user)
	if err != nil {
		return nil, errors.New("查無此信箱！")
	}
	pwMatch := login.Password == whereStruct.Password
	if !pwMatch {
		return nil, errors.New("密碼錯誤！")
	}
	claims := utils.CreateClaims(whereStruct.Id)
	token, err := utils.CreateToken(claims)
	if err != nil {
		return nil, err
	}
	result := &models.UserLoginRes{
		Token: token,
		User:  *whereStruct,
	}
	return result, nil
}
