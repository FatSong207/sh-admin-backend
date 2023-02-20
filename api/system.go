package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/global"
	"SH-admin/models/common"
	"SH-admin/utils"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SystemApi struct {
	_userService IServices.IUserService
	_logService  IServices.ILogService
}

func NewSystemApi() *SystemApi {
	ins := &SystemApi{
		_userService: Services.NewUserService(),
		_logService:  Services.NewLogService(),
	}
	return ins
}

func (s *SystemApi) GetServerInfoApi(ctx *gin.Context) {
	si := new(utils.ServerInfo)
	si.InitOS()
	err := si.InitCpu()
	if err != nil {
		common.Result(common.ErrCodeFailed, nil, ctx)
		return
	}
	err = si.InitRam()
	if err != nil {
		common.Result(common.ErrCodeFailed, nil, ctx)
		return
	}
	err = si.InitDisk()
	if err != nil {
		common.Result(common.ErrCodeFailed, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, si, ctx)
}

func (s *SystemApi) GetDashboardApi(ctx *gin.Context) {
	d := &utils.Dashboard{}
	uc, err := s._userService.GetAll()
	if err != nil {
		common.Result(common.ErrCodeFailed, nil, ctx)
		return
	}
	d.UserCount = len(uc)

	lc, err := s._logService.GetAll()
	if err != nil {
		common.Result(common.ErrCodeFailed, nil, ctx)
		return
	}
	d.LogCount = len(lc)

	vcs, err := global.Rdb.Get(context.Background(), "visitCount").Result()
	if err != nil {
		common.Result(common.ErrCodeFailed, nil, ctx)
		return
	}
	vc, _ := strconv.Atoi(vcs)
	d.VisitCount = vc

	common.Result(common.ErrCodeSuccess, d, ctx)
}
