package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	response "SH-admin/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	*BaseApi[models.User, models.UserOutDto]
	IServices.IUserService
	IServices.IMenuService
}

func NewUserApi() *UserApi {
	ins := &UserApi{
		NewBaseApi[models.User, models.UserOutDto](),
		Services.NewUserService(),
		Services.NewMenuService(),
	}
	return ins
}

// Login 登入
func (u *UserApi) Login(ctx *gin.Context) {
	var param models.UserLoginReq
	err := ctx.ShouldBind(&param)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	login, err := u.IUserService.Login(&param)
	if err != nil {
		response.Result(response.ErrCOdeUserEmailOrPass, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, login, ctx)
}

// GetUserInfoApi 獲取用戶信息
func (u *UserApi) GetUserInfoApi(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		response.Result(response.ErrCodeTokenError, nil, ctx)
		return
	}
	user, err := u.IUserService.GetById(claims.Uid)

	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, user, ctx)
}
