package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	response "SH-admin/models/common"
	"SH-admin/utils"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
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

// GetVerifyCode 根據email發送驗證信並將對應的kv存進去redis
func (u *UserService) GetVerifyCode(email string) int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999998-100000+1) + 100000
	key := fmt.Sprintf("user:%s:code", email)
	if err := global.Rdb.SetEx(context.Background(), key, strconv.Itoa(code), 5*time.Minute).Err(); err != nil {
		return response.ErrCodeVerityCodeSendFailed
	}
	err := utils.SendMail("SHAdmin-驗證信件", fmt.Sprintf("<h2>您的驗證碼如下所示，請勿向他人透露</h2><h2>驗證碼有效期間5分鐘:</h2><h4 style='color:red'>%v</h4>", code), email)
	if err != nil {
		return response.ErrCodeVerityCodeSendFailed
	}
	return response.ErrCodeSuccess
}

func (u *UserService) Register(reg models.UserRegisterReq) int {
	//檢查是否存在
	whereStruct, err := u.userRepo.GetByWhereStruct(&models.User{
		Email: reg.Email,
	})
	if err != gorm.ErrRecordNotFound && err != nil {
		return response.ErrCodeInsertFailed
	}
	if whereStruct != nil {
		return response.ErrCodeUserHasExist
	}

	//檢查驗證碼
	key := "user:" + reg.Email + ":code"
	result, err := global.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return response.ErrCodeVerityCodeInvalid
	}
	if result != reg.Code {
		return response.ErrCodeVerityCodeInvalid
	}

	//新增至db
	user := &models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: reg.Password,
		RoleId:   12,
	}
	err, _ = u.userRepo.Insert(user, false)
	if err != nil {
		return response.ErrCodeInsertFailed
	}
	return response.ErrCodeSuccess
}
