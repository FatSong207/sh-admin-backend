package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	"strconv"
)

type MenuService struct {
	Core.IService[models.Menu, models.MenuOutDto2]
	menuRepo     IRepostories.IMenuRepostory
	roleAuthRepo IRepostories.IRoleAuthorizeRepostory
}

// NewMenuService 供api層調用
func NewMenuService() IServices.IMenuService {
	ins := &MenuService{
		IService:     Core.NewBaseService[models.Menu, models.MenuOutDto2](),
		menuRepo:     Repostories.NewMenuRepostory(),
		roleAuthRepo: Repostories.NewRoleAuthorizeRepostory(),
	}
	return ins
}

func (m *MenuService) GetMenuTreeMap(roleId int64, isCas bool) (treeMap map[string][]models.MenuOutDto, err error) {
	var allMenus []models.MenuOutDto
	var baseMenu []models.Menu
	var db = global.Db.Model(&models.Menu{})
	treeMap = make(map[string][]models.MenuOutDto)

	if isCas {
		db = db.Where("menu_type = ?", 1)
	}

	if roleId == 12 {
		err := db.Order("sort").Find(&baseMenu).Error
		if err != nil {
			return nil, err
		}
	} else {
		roleAuths, err := m.roleAuthRepo.GetListByWhereStruct(&models.RoleAuthorize{
			RoleId: roleId,
		})
		if err != nil {
			return nil, err
		}

		var MenuIds []int64
		for i := range roleAuths {
			MenuIds = append(MenuIds, roleAuths[i].AuthorizeId)
		}

		err = db.Where("Id in (?)", MenuIds).Order("sort").Find(&baseMenu).Error
		if err != nil {
			return nil, err
		}
	}

	for i := range baseMenu {
		allMenus = append(allMenus, models.MenuOutDto{
			Menu:   baseMenu[i],
			RoleId: roleId,
			MenuId: baseMenu[i].Id,
		})
	}

	//err = global.GVA_DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	//if err != nil {
	//	return
	//}
	//var btnMap = make(map[uint]map[string]uint)
	//for _, v := range btns {
	//	if btnMap[v.SysMenuID] == nil {
	//		btnMap[v.SysMenuID] = make(map[string]uint)
	//	}
	//	btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityId
	//}
	for _, v := range allMenus {
		//v.Btns = btnMap[v.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

func (m *MenuService) GetMenuTree(roleId int64) (menus []models.MenuOutDto, err error) {
	//if roleId == 12 { //RoleId==12代表是超級管理員，GetMenuTreeMap(0)可查出所有
	//	roleId = 0
	//}
	menuTree, err := m.GetMenuTreeMap(roleId, false)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = m.GetChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (m *MenuService) GetAllMenuTree(isCas bool) (menus []models.MenuOutDto, err error) {
	menuTree, err := m.GetMenuTreeMap(12, isCas) //獲取全部等同於超級管理員可訪問的菜單所以參數帶12
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = m.GetChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (m *MenuService) GetChildrenList(menu *models.MenuOutDto, treeMap map[string][]models.MenuOutDto) (err error) {
	menu.Children = treeMap[strconv.FormatInt(menu.MenuId, 10)]
	for i := 0; i < len(menu.Children); i++ {
		err = m.GetChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
