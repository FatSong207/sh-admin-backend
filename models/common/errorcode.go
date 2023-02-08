package common

const (
	ErrCodeSuccess      = 0 // 成功
	ErrCodeFailed       = 1 // 失败
	ErrCodeParamInvalid = 2 // 參數錯誤
	ErrCodeNoLogin      = 3 // 尚未登入或非法訪問
	ErrCodeTokenExpire  = 4 // token已過期
	ErrCodeInsertFailed = 5 //新增失敗
	ErrCodeUpdateFailed = 6 //更新失敗

	ErrCode403 = 403 //權限不足

	ErrCodeUserHasExist         = 10001 // 用户已经存在
	ErrCodeUserNotExist         = 10002 // 用户不存在
	ErrCOdeUserEmailOrPass      = 10003 // 帳號密碼錯誤
	ErrCodeTokenError           = 10041 // Token錯誤
	ErrCodeVerityCodeSendFailed = 10004 // 驗證碼發送失敗
	ErrCodeVerityCodeInvalid    = 10005 // 驗證碼無效
	//ErrCodeCompanyCreateFailed  = 10006 // 企业创建失败
	//ErrCodeCompanyIdNotExist    = 10007 // 企业编号不存在
	//ErrCodeEmailFormatInvalid   = 10008 // 邮箱格式无效
	//ErrCodeUserPassResetFailed  = 10009 // 用户密码重置失败
)

var msg = map[int]string{
	ErrCodeSuccess:              "success",
	ErrCodeFailed:               "failed",
	ErrCodeParamInvalid:         "參數錯誤",
	ErrCodeNoLogin:              "尚未登入或非法訪問",
	ErrCodeTokenExpire:          "Token已過期",
	ErrCodeInsertFailed:         "新增失敗",
	ErrCodeUpdateFailed:         "更新失敗",
	ErrCode403:                  "權限不足",
	ErrCodeUserHasExist:         "用户已經存在",
	ErrCodeUserNotExist:         "用戶不存在",
	ErrCOdeUserEmailOrPass:      "帳號密碼錯誤",
	ErrCodeTokenError:           "Token錯誤",
	ErrCodeVerityCodeSendFailed: "驗證碼發送失敗",
	ErrCodeVerityCodeInvalid:    "驗證碼無效",
	//ErrCodeCompanyCreateFailed:  "company create failed",
	//ErrCodeEmailFormatInvalid:   "email format invalid",
	//ErrCodeUserPassResetFailed:  "user password reset failed",
}
