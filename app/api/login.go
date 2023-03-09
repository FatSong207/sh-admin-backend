package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"

	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
	_userService interface_services.IUserService
}

func NewLoginApi() *LoginApi {
	ins := &LoginApi{
		_userService: services.NewUserService(),
	}
	return ins
}

// Login 登入
func (l *LoginApi) Login(ctx *gin.Context) {
	var param models.UserLoginReq
	err := ctx.ShouldBind(&param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	login, err := l._userService.Login(&param)
	if err != nil {
		common.ResultWithMessage(common.ErrCOdeUserEmailOrPass, err.Error(), nil, ctx)
		return
	}
	//utils.SendMail("成功登入", "<h1>login success</h1>", login.User.Email)
	common.Result(common.ErrCodeSuccess, login, ctx)
}

// GetUserInfoApi 登入後獲取用戶信息
func (l *LoginApi) GetUserInfoApi(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		common.Result(common.ErrCodeTokenError, nil, ctx)
		return
	}
	user, err := l._userService.GetById(claims.Uid)

	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, user, ctx)
}

// GetVerifyCode 根據email發送驗證碼
func (l *LoginApi) GetVerifyCode(ctx *gin.Context) {
	e := ctx.Query("email")
	errCode := l._userService.GetVerifyCode(e)
	common.Result(errCode, nil, ctx)
}

// Register 註冊
func (l *LoginApi) Register(ctx *gin.Context) {
	var param = models.UserRegisterReq{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	errCode := l._userService.Register(param)
	common.Result(errCode, nil, ctx)
}
