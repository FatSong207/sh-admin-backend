package api

import (
	response "SH-admin/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

type SystemApi struct {
}

func NewSystemApi() *SystemApi {
	ins := &SystemApi{}
	return ins
}

func (s *SystemApi) GetServerInfo(ctx *gin.Context) {
	si := new(utils.ServerInfo)
	si.InitOS()
	err := si.InitCpu()
	if err != nil {
		response.Result(response.ErrCodeFailed, nil, ctx)
		return
	}
	err = si.InitRam()
	if err != nil {
		response.Result(response.ErrCodeFailed, nil, ctx)
		return
	}
	err = si.InitDisk()
	if err != nil {
		response.Result(response.ErrCodeFailed, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, si, ctx)
}
