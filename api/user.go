package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	response "SH-admin/models/common"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	*BaseApi[models.User, models.UserOutDto]
	IServices.IUserService
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
