package services

import (
	"SH-admin/app/base/core"
	"SH-admin/app/base/interface_repostories"
	"SH-admin/app/base/interface_services"
	"SH-admin/app/base/repostories"
	"SH-admin/app/models"
	"SH-admin/app/models/common"
	"SH-admin/global"
	"encoding/json"
)

type LogService struct {
	core.IBaseService[models.Log, models.LogOutDto]
	_logRepo  interface_repostories.ILogRepostory
	_userRepo interface_repostories.IUserRepostory
}

// NewLogService 供api層調用
func NewLogService() interface_services.ILogService {
	ins := &LogService{
		IBaseService: core.NewBaseService[models.Log, models.LogOutDto](),
		_logRepo:     repostories.NewLogRepostory(),
		_userRepo:    repostories.NewUserRepostory(),
	}
	return ins
}

func (l *LogService) FindWithPager(searchDto common.SearchDto[models.Log]) (*[]*models.LogOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.LogOutDto, 0)
	var bind = make([]*models.Log, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	db := global.DB().Model(&query)
	db = db.Where(" type = ? ", "normalOp")
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	t, err := l._logRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	fdest := make([]*models.LogOutDto, 0)
	for _, item := range dest {
		if item.UserID == 0 { //當UserID == 0 代表是不需要授權的訪問但是逼須記錄在log，例如註冊
			fdest = append(fdest, item)
			continue
		}
		u, err := l._userRepo.GetById(int64(item.UserID))
		if err != nil {
			return nil, 0, err
		}
		item.UserName = u.Name
		fdest = append(fdest, item)
	}
	return &fdest, t, nil
}

func (l *LogService) FindLoginlogWithPager(searchDto common.SearchDto[models.Log]) (*[]*models.LoginlogOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.LoginlogOutDto, 0)
	var bind = make([]*models.LoginlogOutDto, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}

	db := global.DB().Model(&query)
	db = db.Where(" type = ? ", "login")
	if query.Response != "" {
		switch query.Response {
		case "success":
			db = db.Where("response LIKE ?", "%\"message\":\"success\"%")
		case "failed":
			db = db.Where("response LIKE ?", "%\"data\":null%")
		}
	}

	t, err := l._logRepo.FindLoginlogWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	var fdest = make([]*models.LoginlogOutDto, 0)
	for _, item := range dest {
		log, _ := l._logRepo.GetById(item.Id)
		bs := []byte(log.Body)
		m := make(map[string]any)
		json.Unmarshal(bs, &m)
		item.Email = m["email"].(string)
		fdest = append(fdest, item)
	}
	return &fdest, t, nil
}
