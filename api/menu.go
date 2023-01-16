package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	response "SH-admin/models/common"
	"SH-admin/utils"
	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	*BaseApi[models.Menu, models.MenuOutDto]
	IServices.IMenuService
	IServices.IUserService
}

func NewMenuApi() *MenuApi {
	ins := &MenuApi{
		NewBaseApi[models.Menu, models.MenuOutDto](),
		Services.NewMenuService(),
		Services.NewUserService(),
	}
	return ins
}

// GetMenuTree 根據Token獲取使用者的功能權限
func (m *MenuApi) GetMenuTree(ctx *gin.Context) {
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

// GetAllMenuTree 獲取所有功能模塊，用於功能模塊下的樹狀table顯示
func (m *MenuApi) GetAllMenuTree(ctx *gin.Context) {
	tree, err := m.IMenuService.GetAllMenuTree()
	if err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, ctx)
		return
	}
	response.Result(response.ErrCodeSuccess, tree, ctx)
}
