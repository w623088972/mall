package conf

import (
	"github.com/beijibeijing/viper"
)

//ErrCode 返回结构体
type ErrCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"` //en英语，chs简中，cht繁中
}

//GetECode 从配置文件获取错误码
func GetECode(code, language string) *ErrCode {
	return &ErrCode{
		Code:    viper.GetInt("ErrCode." + code + ".code"),
		Message: viper.GetString("ErrCode." + code + ".message." + language),
	}
}

/* var (
	OK = &ErrCode{Code: 10000, Message: map[string]string{"en": viper.GetString("ErrCode.OK.en"), "chs": viper.GetString("ErrCode.OK.chs"), "cht": viper.GetString("ErrCode.OK.cht")}}

	InternalServerError = &ErrCode{Code: 10001, Message: map[string]string{"en": viper.GetString("ErrCode.InternalServerError.en"), "chs": viper.GetString("ErrCode.InternalServerError.chs"), "cht": viper.GetString("ErrCode.InternalServerError.cht")}}
	ErrDatabase         = &ErrCode{Code: 10002, Message: map[string]string{"en": viper.GetString("ErrCode.ErrDatabase.en"), "chs": viper.GetString("ErrCode.ErrDatabase.chs"), "cht": viper.GetString("ErrCode.ErrDatabase.cht")}}
	ErrRedis            = &ErrCode{Code: 10003, Message: map[string]string{"en": viper.GetString("ErrCode.ErrRedis.en"), "chs": viper.GetString("ErrCode.ErrRedis.chs"), "cht": viper.GetString("ErrCode.ErrRedis.cht")}}
	ErrParameterWrong   = &ErrCode{Code: 10006, Message: map[string]string{"en": viper.GetString("ErrCode.ErrParameterWrong.en"), "chs": viper.GetString("ErrCode.ErrParameterWrong.chs"), "cht": viper.GetString("ErrCode.ErrParameterWrong.cht")}}

	ErrTokenCreateFailed = &ErrCode{Code: 10101, Message: map[string]string{"en": viper.GetString("ErrCode.ErrTokenCreateFailed.en"), "chs": viper.GetString("ErrCode.ErrTokenCreateFailed.chs"), "cht": viper.GetString("ErrCode.ErrTokenCreateFailed.cht")}}
	ErrTokenInvalid      = &ErrCode{Code: 10105, Message: map[string]string{"en": viper.GetString("ErrCode.ErrTokenInvalid.en"), "chs": viper.GetString("ErrCode.ErrTokenInvalid.chs"), "cht": viper.GetString("ErrCode.ErrTokenInvalid.cht")}}

	ErrSmsOneMinute  = &ErrCode{Code: 10112, Message: map[string]string{"en": viper.GetString("ErrCode.ErrSmsOneMinute.en"), "chs": viper.GetString("ErrCode.ErrSmsOneMinute.chs"), "cht": viper.GetString("ErrCode.ErrSmsOneMinute.cht")}}
	ErrSms24Hours    = &ErrCode{Code: 10113, Message: map[string]string{"en": viper.GetString("ErrCode.ErrSms24Hours.en"), "chs": viper.GetString("ErrCode.ErrSms24Hours.chs"), "cht": viper.GetString("ErrCode.ErrSms24Hours.cht")}}
	ErrSmsSendFailed = &ErrCode{Code: 10114, Message: map[string]string{"en": viper.GetString("ErrCode.ErrSmsSendFailed.en"), "chs": viper.GetString("ErrCode.ErrSmsSendFailed.chs"), "cht": viper.GetString("ErrCode.ErrSmsSendFailed.cht")}}
	ErrSmsVCodeWrong = &ErrCode{Code: 10115, Message: map[string]string{"en": viper.GetString("ErrCode.ErrSmsVCodeWrong.en"), "chs": viper.GetString("ErrCode.ErrSmsVCodeWrong.chs"), "cht": viper.GetString("ErrCode.ErrSmsVCodeWrong.cht")}}

	ErrUserBan = &ErrCode{Code: 40004, Message: map[string]string{"en": viper.GetString("ErrCode.ErrUserBan.en"), "chs": viper.GetString("ErrCode.ErrUserBan.chs"), "cht": viper.GetString("ErrCode.ErrUserBan.cht")}}
) */
