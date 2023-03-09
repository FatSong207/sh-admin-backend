package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RoleAuthorizeApi struct {
	*BaseApi[models.RoleAuthorize, models.RoleAuthorize]
	iService     interface_services.IRoleAuthorizeService
	_userService interface_services.IUserService
}

func NewRoleAuthorizeApi() *RoleAuthorizeApi {
	ins := &RoleAuthorizeApi{
		NewBaseApi[models.RoleAuthorize, models.RoleAuthorize](),
		services.NewRoleAuthorizeService(),
		services.NewUserService(),
	}
	return ins
}

// GetAuthorizeIds 根據RoleId或取擁有的AuthorizeId
func (r *RoleAuthorizeApi) GetAuthorizeIds(ctx *gin.Context) {
	rid := ctx.Param("roleid")
	i, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	whereStruct, err := r.iService.GetListByWhereStruct(&models.RoleAuthorize{RoleId: i})
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	result := make([]int64, 0)
	for _, authorize := range whereStruct {
		result = append(result, authorize.AuthorizeId)
	}
	common.Result(common.ErrCodeSuccess, result, ctx)
}

func (r *RoleAuthorizeApi) UpdateApi(ctx *gin.Context) {
	t := new(models.RoleAuthorizeUpdateDto)
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	/**CHECK**/

	//rm := structs.Map(t)
	i, err := r.iService.UpdateBatchByRoleId(t.RoleId, t.AuthorizeIds)
	if err != nil {
		common.Result(common.ErrCodeInsertFailed, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, i, ctx)
}
