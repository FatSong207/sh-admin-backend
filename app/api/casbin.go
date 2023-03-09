package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CasbinApi struct {
	_userService interface_services.IUserService
	_logService  interface_services.ILogService
	iService     interface_services.ICasbinService
}

func NewCasbinApi() *CasbinApi {
	ins := &CasbinApi{
		iService:     services.NewCasbinService(),
		_userService: services.NewUserService(),
		_logService:  services.NewLogService(),
	}
	return ins
}

func (c *CasbinApi) GetAccessApiPathByRoleId(ctx *gin.Context) {
	param := ctx.Param("roleid")
	p := c.iService.GetAccessApiByRoleId(param)
	result := make([]string, 0)
	for _, v := range p {
		result = append(result, fmt.Sprintf("p:%sm:%s", v[1], v[2]))
	}
	common.Result(common.ErrCodeSuccess, result, ctx)
}

func (c *CasbinApi) UpdateCasbin(ctx *gin.Context) {
	t := new(models.UpdateCasbinParam)
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	err = c.iService.UpdateCasbin(t)
	if err != nil {
		common.ResultWithMessage(common.ErrCodeFailed, err.Error(), nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, nil, ctx)
}
