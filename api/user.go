package api

import (
	"SH-admin/common/IServices"
	"SH-admin/common/Services"
	"SH-admin/models"
	"SH-admin/models/common"
	"SH-admin/utils"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct {
	*BaseApi[models.User, models.UserOutDto]
	iService       IServices.IUserService
	_casbinService IServices.ICasbinService
}

func NewUserApi() *UserApi {
	ins := &UserApi{
		BaseApi:        NewBaseApi[models.User, models.UserOutDto](),
		iService:       Services.NewUserService(),
		_casbinService: Services.NewCasbinService(),
	}
	return ins
}

func (u *UserApi) FindWithPagerApi(ctx *gin.Context) {
	var param = common.NewSearchDto[models.User]()
	//ShouldBindQuery：把query string binding到struct，struct裡面的tag要用form:"xxx"
	//ShouldBindJSON：把POST Body binding到struct，struct裡面的tag要用json:"xxx"
	err := ctx.ShouldBind(param) //ShouldBind必須在目標結構體給定form標籤
	//err := ctx.ShouldBindQuery(param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	withPager, i, err := u.iService.FindWithPager(*param)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	common.PageResult(common.ErrCodeSuccess, withPager, i, ctx)
}

func (u *UserApi) UpdateApi(ctx *gin.Context) {
	t := new(models.User)
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	user, err := u.iService.GetById(t.Id)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}

	//修改user
	temps := struct {
		RoleId int64 `json:"roleId"`
		Status int   `json:"status"`
	}{t.RoleId, t.Status}
	m := structs.Map(&temps)
	update, err := u.iService.Update(&models.User{Id: t.Id}, m, false)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	if update == 0 {
		common.Result(common.ErrCodeUpdateFailed, nil, ctx)
		return
	}

	//修改casbin的role
	err = u._casbinService.UpdateUserRole([]string{strconv.FormatInt(t.Id, 10), strconv.FormatInt(user.RoleId, 10)}, []string{strconv.FormatInt(t.Id, 10), strconv.FormatInt(temps.RoleId, 10)})
	if err != nil {
		common.ResultWithMessage(common.ErrCodeUpdateFailed, err.Error(), nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, update, ctx)
}

func (u *UserApi) SendMailToUserApi(ctx *gin.Context) {
	t := new(struct {
		Name    string `json:"name" form:"name"`
		Email   string `json:"email" form:"email"`
		Content string `json:"content" form:"content"`
	})
	err := ctx.ShouldBind(t)
	if err != nil {
		common.Result(common.ErrCodeParamInvalid, nil, ctx)
		return
	}
	err = utils.SendMail("SHAdmin-通知信件", fmt.Sprintf("<h2>您好，%s 先生/小姐：</h2>%s", t.Name, t.Content), t.Email)
	if err != nil {
		common.ResultWithMessage(common.ErrCodeFailed, "信件發送失敗:"+err.Error(), nil, ctx)
		return
	}
	common.Result(common.ErrCodeSuccess, nil, ctx)
}
