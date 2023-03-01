package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/global"
	"SH-admin/models"
	"SH-admin/models/common"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type RoleApi struct {
	*BaseApi[models.Role, models.RoleOutDto]
	iService IServices.IRoleService
}

func NewRoleApi() *RoleApi {
	ins := &RoleApi{
		NewBaseApi[models.Role, models.RoleOutDto](),
		Services.NewRoleService(),
	}
	return ins
}

func (r *RoleApi) InsertApi(ctx *gin.Context) {
	var t = new(models.Role)
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	err, i := r.baseSvc.Insert(t, false)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	//新增casbin基礎規則
	rules := [][]string{{strconv.FormatInt(t.Id, 10), "/api/menu/tree", "GET"}, {strconv.FormatInt(t.Id, 10), "/api/system/dashboard", "GET"}, {strconv.FormatInt(t.Id, 10), "/api/user/info", "GET"}}
	success, err := global.CachedEnforcer.AddPolicies(rules)
	if err != nil {
		common.ResultWithMessage(common.ErrCodeFailed, err.Error(), nil, ctx)
		return
	}
	if !success {
		common.ResultWithMessage(common.ErrCodeFailed, "已存在相同api", nil, ctx)
		return
	}
	_ = global.CachedEnforcer.InvalidateCache()

	common.Result(common.ErrCodeSuccess, i, ctx)
}

func (r *RoleApi) UpdateApi(ctx *gin.Context) {
	t := new(models.RoleInDto)
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	/**CHECK**/
	sql, err := r.iService.GetBySQL(fmt.Sprintf("SELECT * FROM role WHERE name ='%s' AND en_name ='%s' ", t.Name, t.EnName))
	if err != nil && err != gorm.ErrRecordNotFound {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if sql != nil {
		common.ResultWithMessage(common.ErrCodeInsertFailed, "已存在相同名稱的角色", nil, ctx)
		return
	}

	rm := structs.Map(t)
	i, err := r.iService.Update(&models.Role{Id: t.Id}, rm, false)
	if err != nil {
		common.Result(common.ErrCodeInsertFailed, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, i, ctx)
}
