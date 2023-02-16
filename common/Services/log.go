package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	response "SH-admin/models/common"
	"encoding/json"
)

type LogService struct {
	Core.IService[models.Log, models.LogOutDto]
	logRepo IRepostories.ILogRepostory
}

// NewLogService 供api層調用
func NewLogService() IServices.ILogService {
	ins := &LogService{
		IService: Core.NewBaseService[models.Log, models.LogOutDto](),
		logRepo:  Repostories.NewLogRepostory(),
	}
	return ins
}

func (l *LogService) FindWithPager(searchDto response.SearchDto[models.Log]) (*[]*models.Log, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.Log, 0)
	var bind = make([]*models.Log, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	db := global.Db.Model(&query)
	db = db.Where(" type = ? ", "normalOp")
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	t, err := l.logRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	return &dest, t, nil
}

func (l *LogService) FindLoginlogWithPager(searchDto response.SearchDto[models.Log]) (*[]*models.LoginlogOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.LoginlogOutDto, 0)
	var bind = make([]*models.LoginlogOutDto, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	db := global.Db.Model(&query)
	db = db.Where(" type = ? ", "login")
	if query.Response != "" {
		switch query.Response {
		case "success":
			db = db.Where("response LIKE ?", "%\"message\":\"success\"%")
		case "failed":
			db = db.Where("response LIKE ?", "%\"data\":null%")
		}
	}
	t, err := l.logRepo.FindLoginlogWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	var fdest = make([]*models.LoginlogOutDto, 0)
	for _, item := range dest {
		log, _ := l.logRepo.GetById(item.Id)
		bs := []byte(log.Body)
		m := make(map[string]any)
		json.Unmarshal(bs, &m)
		item.Email = m["email"].(string)
		fdest = append(fdest, item)
	}
	return &fdest, t, nil
}
