package api

import (
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/services"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"SH-admin/global"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type RoleApi struct {
	*BaseApi[models.Role, models.RoleOutDto]
	iService              interface_services.IRoleService
	_roleAuthorizeService interface_services.IRoleAuthorizeService
}

func NewRoleApi() *RoleApi {
	ins := &RoleApi{
		BaseApi:               NewBaseApi[models.Role, models.RoleOutDto](),
		iService:              services.NewRoleService(),
		_roleAuthorizeService: services.NewRoleAuthorizeService(),
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
	//基礎menu(一定會有儀表盤)
	err, _ = r._roleAuthorizeService.Insert(&models.RoleAuthorize{RoleId: t.Id, AuthorizeId: 1}, false)
	if err != nil {
		common.Result(common.ErrCodeInsertFailed, nil, ctx)
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
		common.Result(common.ErrCodeUpdateFailed, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, i, ctx)
}

func (r *RoleApi) DeleteApi(ctx *gin.Context) {
	param := ctx.Param("ids")
	if param == "" {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	split := strings.Split(param, ",")
	ks := make([]int, 0)
	for _, s := range split {
		v, err := strconv.Atoi(s)
		if err != nil {
			common.Result(common.ErrCodeParamInvalid, nil, ctx)
			return
		}
		ks = append(ks, v)
	}

	rowsAffected, err := r.iService.DeleteByKeys(ks)
	if err != nil {
		common.Result(common.ErrCodeDeleteFailed, nil, ctx)
		return
	}
	if rowsAffected == 0 {
		common.Result(common.ErrCodeDeleteFailed, nil, ctx)
		return
	}

	//刪除casbin規則
	success, err := global.CachedEnforcer.RemoveFilteredPolicy(0, split...)
	if err != nil {
		common.ResultWithMessage(common.ErrCodeDeleteFailed, err.Error(), nil, ctx)
		return
	}
	if !success {
		common.Result(common.ErrCodeDeleteFailed, nil, ctx)
		return
	}
	_ = global.CachedEnforcer.InvalidateCache()

	common.Result(common.ErrCodeSuccess, rowsAffected, ctx)
}

func (r RoleApi) GetAllRoleListForSelect(ctx *gin.Context) {
	all, err := r.iService.GetAll()
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, all, ctx)
}

func (r *RoleApi) FindWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.Role]()
	//ShouldBindQuery：把query string binding到struct，struct裡面的tag要用form:"xxx"
	//ShouldBindJSON：把POST Body binding到struct，struct裡面的tag要用json:"xxx"
	err := ctx.ShouldBind(param) //ShouldBind必須在目標結構體給定form標籤
	//err := ctx.ShouldBindQuery(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := r.iService.FindWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}
