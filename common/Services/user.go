package Services

import (
	"SH-admin/common/Core"
	"SH-admin/common/IRepostories"
	"SH-admin/common/IServices"
	"SH-admin/common/Repostories"
	"SH-admin/global"
	"SH-admin/models"
	"SH-admin/models/common"
	"SH-admin/utils"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type UserService struct {
	Core.IService[models.User, models.UserOutDto]
	_userRepo IRepostories.IUserRepostory
	_roleRepo IRepostories.IRoleRepostory
}

// NewUserService 供api層調用
func NewUserService() IServices.IUserService {
	ins := &UserService{
		IService:  Core.NewBaseService[models.User, models.UserOutDto](),
		_userRepo: Repostories.NewUserRepostory(),
		_roleRepo: Repostories.NewRoleRepostory(),
	}
	return ins
}

func (u *UserService) Login(login *models.UserLoginReq) (*models.UserLoginRes, error) {
	user := &models.User{
		Email: login.Email,
	}

	//基本檢查
	whereStruct, err := u._userRepo.GetByWhereStruct(user)
	if err != nil {
		return nil, errors.New("查無此信箱！")
	}
	pwMatch := login.Password == whereStruct.Password
	if !pwMatch {
		return nil, errors.New("密碼錯誤！")
	}
	if whereStruct.Status == 0 {
		return nil, errors.New("帳號已被禁用！")
	}

	//產生token
	claims := utils.CreateClaims(whereStruct.Id, whereStruct.RoleId)
	token, err := utils.CreateToken(claims)
	if err != nil {
		return nil, err
	}
	result := &models.UserLoginRes{
		Token: token,
		User:  *whereStruct,
	}

	//記錄登入數量
	r, _ := global.Rdb.Exists(context.Background(), "visitCount").Result()
	if r == 1 {
		global.Rdb.Incr(context.Background(), "visitCount")
	} else {
		t := time.Now().AddDate(0, 0, 1).Format("2006/01/02")
		tt, _ := time.ParseInLocation("2006/01/02", t, time.Local)
		r := tt.Unix() - time.Now().Unix()

		global.Rdb.SetEx(context.Background(), "visitCount", 1, time.Duration(r)*time.Second)
	}

	return result, nil
}

// GetVerifyCode 根據email發送驗證信並將對應的kv存進去redis
func (u *UserService) GetVerifyCode(email string) int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999998-100000+1) + 100000
	key := fmt.Sprintf("user:%s:code", email)
	if err := global.Rdb.SetEx(context.Background(), key, strconv.Itoa(code), 5*time.Minute).Err(); err != nil {
		return common.ErrCodeVerityCodeSendFailed
	}
	err := utils.SendMail("SHAdmin-驗證信件", fmt.Sprintf("<h2>您的驗證碼如下所示，請勿向他人透露</h2><h2>驗證碼有效期間5分鐘:</h2><h4 style='color:red'>%v</h4>", code), email)
	if err != nil {
		return common.ErrCodeVerityCodeSendFailed
	}
	return common.ErrCodeSuccess
}

func (u *UserService) Register(reg models.UserRegisterReq) int {
	//檢查是否存在
	whereStruct, err := u._userRepo.GetByWhereStruct(&models.User{
		Email: reg.Email,
	})
	if err != gorm.ErrRecordNotFound && err != nil {
		return common.ErrCodeInsertFailed
	}
	if whereStruct != nil {
		return common.ErrCodeUserHasExist
	}

	//檢查驗證碼
	key := "user:" + reg.Email + ":code"
	result, err := global.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return common.ErrCodeVerityCodeInvalid
	}
	if result != reg.Code {
		return common.ErrCodeVerityCodeInvalid
	}

	//新增至db
	user := &models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: reg.Password,
		RoleId:   14,
		Status:   1,
	}
	err, _ = u._userRepo.Insert(user, false)
	if err != nil {
		return common.ErrCodeInsertFailed
	}
	//新增基礎casbin規則
	global.CachedEnforcer.AddRoleForUser(strconv.FormatInt(user.Id, 10), strconv.FormatInt(user.RoleId, 10))
	return common.ErrCodeSuccess
}

func (u *UserService) FindWithPager(searchDto common.SearchDto[models.User]) (*[]*models.UserOutDto, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*models.UserOutDto, 0)
	var bind = make([]*models.User, 0)
	var o = ""
	for k, i := range searchDto.OrderRule.OrderBy {
		o += k + " " + i
	}
	db := global.Db.Model(&query)

	//db = db.Where(" type = ? ", "normalOp")
	if query.Name != "" {
		db = db.Where("name like ?", "%"+query.Name+"%")
	}
	if query.RoleId != 0 {
		db = db.Where("role_id = ?", query.RoleId)
	}

	t, err := u._userRepo.FindWithPager(searchDto.PageInfo, db, o, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	fdest := make([]*models.UserOutDto, 0)
	for _, item := range dest {
		role, err := u._roleRepo.GetById(item.RoleId)
		if err != nil {
			item.RoleName = ""
		} else {
			item.RoleName = role.Name
		}
		fdest = append(fdest, item)
	}
	return &fdest, t, nil
}
