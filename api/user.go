package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	"SH-admin/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	*BaseApi[models.User, models.UserOutDto]
	iService IServices.IUserService
}

func NewUserApi() *UserApi {
	ins := &UserApi{
		NewBaseApi[models.User, models.UserOutDto](),
		Services.NewUserService(),
	}
	return ins
}

// Login 登入
func (u *UserApi) Login(ctx *gin.Context) {
	var param models.UserLoginReq
	err := ctx.ShouldBind(&param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	login, err := u.iService.Login(&param)
	if err != nil {
		common.Result(common.ErrCOdeUserEmailOrPass, nil, ctx)
		return
	}
	//utils.SendMail("成功登入", "<h1>login success</h1>", login.User.Email)
	common.Result(common.ErrCodeSuccess, login, ctx)
}

// GetUserInfoApi 獲取用戶信息
func (u *UserApi) GetUserInfoApi(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		common.Result(common.ErrCodeTokenError, nil, ctx)
		return
	}
	user, err := u.iService.GetById(claims.Uid)

	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, user, ctx)
}

// GetVerifyCode 根據email發送驗證碼
func (u *UserApi) GetVerifyCode(ctx *gin.Context) {
	e := ctx.Query("email")
	errCode := u.iService.GetVerifyCode(e)
	common.Result(errCode, nil, ctx)
}

// Register 註冊
func (u *UserApi) Register(ctx *gin.Context) {
	var param = models.UserRegisterReq{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := u.iService.Register(param)
	common.Result(errCode, nil, ctx)
}
