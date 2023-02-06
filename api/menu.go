package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	response "SH-admin/models/common"
	"SH-admin/utils"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MenuApi struct {
	*BaseApi[models.Menu, models.MenuOutDto2]
	IServices.IMenuService
	IServices.IUserService
}

func NewMenuApi() *MenuApi {
	ins := &MenuApi{
		BaseApi:      NewBaseApi[models.Menu, models.MenuOutDto2](),
		IMenuService: Services.NewMenuService(),
		IUserService: Services.NewUserService(),
	}
	return ins
}

func (m *MenuApi) GetByIdApi(ctx *gin.Context) {
	id := ctx.Param("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	//getById, err := b.baseSvc.GetByIdApi(i)
	getById, err := m.IMenuService.GetOutDtoById(i)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	var ps = ""
	var p = getById.ParentId
	for true {
		if p == "0" {
			getById.ParentIds = ps
			break
		} else {
			parseInt, _ := strconv.ParseInt(p, 10, 64)
			whereStruct, _ := m.IMenuService.GetByWhereStruct(&models.Menu{
				Id: parseInt,
			})
			ps += fmt.Sprintf("%v,", whereStruct.Id)
			p = fmt.Sprintf("%v", whereStruct.ParentId)
		}
	}
	getById.ParentIds = ps
	response.Result(response.ErrCodeSuccess, getById, ctx)
}

func (m *MenuApi) InsertApi(ctx *gin.Context) {
	var menu = new(models.Menu)
	err := ctx.ShouldBind(menu)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}

	/***額外處理***/
	if menu.ParentId == "" {
		menu.ParentId = "0"
	}
	menu.ChName = menu.Meta.Title
	if menu.MenuType == 1 {
		menu.Component = ""
	}

	err, i := m.IMenuService.Insert(menu, false)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, i, ctx)
}

func (m *MenuApi) UpdateApi(ctx *gin.Context) {
	menu := new(models.MenuInDto)
	err := ctx.ShouldBind(menu)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}

	/***額外處理(把meta拆解出來)***/
	//metaMap := structs.Map(&menu.Meta)
	//mm := structs.Map(menu)
	//delete(mm, "Meta")
	//for k, v := range metaMap {
	//	mm[k] = v
	//}
	mm := structs.Map(menu)
	if mm["ParentId"] == "" {
		mm["ParentId"] = 0
	}
	mm["ChName"] = mm["Title"]
	if mm["MenuType"] == 1 {
		mm["Component"] = ""
	}

	update, err := m.IMenuService.Update(&models.Menu{Id: menu.Id}, mm, false)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if update == 0 {
		response.Result(response.ErrCodeUpdateFailed, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, update, ctx)

}

// GetMenuTreeApi 根據Token獲取使用者的功能權限
func (m *MenuApi) GetMenuTreeApi(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		response.Result(response.ErrCodeTokenError, nil, ctx)
		return
	}
	user, err := m.IUserService.GetById(claims.Uid)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}

	treeMap, err := m.IMenuService.GetMenuTree(user.RoleId)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}

	response.Result(response.ErrCodeSuccess, treeMap, ctx)
}

// GetAllMenuTreeApi 獲取所有功能模塊，用於功能模塊下的樹狀table顯示
func (m *MenuApi) GetAllMenuTreeApi(ctx *gin.Context) {
	tree, err := m.IMenuService.GetAllMenuTree(false)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, tree, ctx)
}

// GetAllMenuTreeCasApi 獲取所有功能模塊，用於功能模塊下的Cascader
func (m *MenuApi) GetAllMenuTreeCasApi(ctx *gin.Context) {
	tree, err := m.IMenuService.GetAllMenuTree(true)
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, tree, ctx)
}
