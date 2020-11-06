package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WeChatMiniProgramAccessToken struct {
	AccessToken string `json:"access_token"` //获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //凭证有效时间，单位：秒。目前是7200秒之内的值。
	WeChatError
}

//获取微信小程序 accessToken
func GetWeChatMiniProgramAccessToken(appId, secret string) (string, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=%s"
	grantType := "client_credential"
	path := fmt.Sprintf(url, appId, secret, grantType)
	log.Println("WeChatMiniProgram", path)

	resp, err := http.Get(path)
	if err != nil {
		log.Println("GetWeChatMiniProgramAccessToken http.Get error is " + err.Error())
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("GetWeChatMiniProgramAccessToken ioutil.ReadAll error is " + err.Error())
		return "", err
	}

	defer resp.Body.Close()

	result := &WeChatMiniProgramAccessToken{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("GetWeChatMiniProgramAccessToken json.Unmarshal error is " + err.Error())
		return "", err
	}
	if result.Errcode != 0 {
		err := fmt.Errorf("GetWeChatMiniProgramAccessToken error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return "", err
	}

	return result.AccessToken, nil
}

//获取微信小程序二维码
//通过该接口生成的小程序码，永久有效，数量暂无限制
func GetWeChatMiniProgramQrCode(accessToken, qrCodeName string, param []byte) error {
	path := "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=" + accessToken
	contentType := "application/x-www-form-urlencoded"

	req := bytes.NewBuffer([]byte(param))
	resp, err := http.Post(path, contentType, req)
	if err != nil {
		log.Println("GetWeChatMiniProgramQrCode http.Post error is " + err.Error())
		return err
	}

	file, err := os.Create(qrCodeName)
	if err != nil {
		log.Println("GetWeChatMiniProgramQrCode os.Create error is " + err.Error())
		panic(err)
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("GetWeChatMiniProgramQrCode file.Close error is " + err.Error())
			return
		}
	}()

	//写入文件
	written, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Println("GetWeChatMiniProgramQrCode io.Copy error is " + err.Error())
		return err
	}
	log.Println("GetWeChatMiniProgramQrCode Total length:", written)

	defer resp.Body.Close()

	return err
}

type WeChatAccessToken struct {
	AccessToken  string `json:"access_token"`  //接口调用凭证
	ExpiresIn    int    `json:"expires_in"`    //access_token 接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` //用户刷新 access_token
	Openid       string `json:"openid"`        //授权用户唯一标识
	Scope        string `json:"scope"`         //用户授权的作用域，使用逗号（,）分隔
	Unionid      string `json:"unionid"`       //当且仅当该移动应用已获得该用户的 userinfo 授权时，才会出现该字段
}
type WeChatError struct {
	Errcode int64  `json:"errcode"` //错误码
	Errmsg  string `json:"errmsg"`  //错误信息
}
type WeChatAccessTokenRes struct {
	WeChatAccessToken
	WeChatError
}

//获取微信access_token
func GetWeChatAppAccessToken(code, appId, appSecret string) (*WeChatAccessTokenRes, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=%s"
	grantType := "authorization_code"
	path := fmt.Sprintf(url, appId, appSecret, code, grantType)
	response, err := http.Get(path)
	if err != nil {
		log.Println("GetWeChatAppAccessToken http.Get error is " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("GetWeChatAppAccessToken ioutil.ReadAll error is " + err.Error())
		return nil, err
	}

	result := new(WeChatAccessTokenRes)
	err = json.Unmarshal(body, &result)
	if result.Errcode != 0 {
		err := fmt.Errorf("GetWeChatAppAccessToken error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return nil, err
	}

	return result, nil
}

type WeChatChatUserInfo struct {
	Openid     string      `json:"openid"`     //普通用户的标识，对当前开发者帐号唯一
	Nickname   string      `json:"nickname"`   //普通用户昵称
	Sex        int         `json:"sex"`        //普通用户性别，1 为男性，2 为女性
	Province   string      `json:"province"`   //普通用户个人资料填写的省份
	City       string      `json:"city"`       //普通用户个人资料填写的城市
	Country    string      `json:"country"`    //国家，如中国为 CN
	Headimgurl string      `json:"headimgurl"` //用户头像，最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640*640 正方形头像），用户没有头像时该项为空
	Privilege  interface{} `json:"privilege"`  //用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Unionid    string      `json:"unionid"`    //用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的 unionid 是唯一的
}
type WeChatChatUserInfoRes struct {
	WeChatChatUserInfo
	WeChatError
}

//lang 国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语，默认为 zh-CN
func GetWaChatUserInfo(accessToken, openId, lang string) (*WeChatChatUserInfoRes, error) {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
	path := fmt.Sprintf(url, accessToken, openId, lang)
	response, err := http.Get(path)
	if err != nil {
		log.Println("GetWaChatUserInfo http.Get error is " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("GetWaChatUserInfo ioutil.ReadAll error is " + err.Error())
		return nil, err
	}

	result := new(WeChatChatUserInfoRes)
	err = json.Unmarshal(body, &result)
	if result.Errcode != 0 {
		err := fmt.Errorf("GetWaChatUserInfo error: errcode=%v, errmsg=%v", result.Errcode, result.Errmsg)
		return nil, err
	}

	return result, nil
}

var (
	templateIdArr = map[int]string{
		1: "q54tw8FaEumFQal64qgOFl7DdyHhnrdTRISvCCzn9dU", //下单成功通知
		2: "J19zI3L2KtQfCgCMvXn-wDH31I6QCHpxk7MHqs0HrSQ", //取餐提醒
		3: "r3_JsFzASkZPO9vkVSQudgFIDka8kOzJ_9MPaYgR_qs", //订单完成
	}
)

type TemplateMsg struct {
	Touser           string        `json:"touser"`            //接收者（用户）的 openid
	TemplateId       string        `json:"template_id"`       //所需下发的订阅模板id
	Page             string        `json:"page"`              //点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             *TemplateData `json:"data"`              //模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	MiniprogramState string        `json:"miniprogram_state"` //跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             Miniprogram   `json:"lang"`              //进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}
type Miniprogram struct {
	AppId    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}
type KeyWordData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}
type TemplateData struct {
	Amount12          KeyWordData `json:"amount12,omitempty"`
	Date4             KeyWordData `json:"date4,omitempty"`
	Number1           KeyWordData `json:"number1,omitempty"`
	Phrase11          KeyWordData `json:"phrase11,omitempty"`
	Thing2            KeyWordData `json:"thing2,omitempty"`
	Thing5            KeyWordData `json:"thing5,omitempty"`
	Thing6            KeyWordData `json:"thing6,omitempty"`
	Thing7            KeyWordData `json:"thing7,omitempty"`
	Thing17           KeyWordData `json:"thing17,omitempty"`
	Thing22           KeyWordData `json:"thing22,omitempty"`
	Time8             KeyWordData `json:"time8,omitempty"`
	Time12            KeyWordData `json:"time12,omitempty"`
	CharacterString19 KeyWordData `json:"character_string19,omitempty"`
}
type SendTemplateResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	MsgId   string `json:"msgid"`
}

//发送订阅消息
func SendSubscribeMessage(accessToken string, templateType int, msg *TemplateMsg) (*SendTemplateResponse, error) {
	url := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
	path := fmt.Sprintf(url, accessToken)
	log.Println("SendUniformMessage path is ", path)

	msg.TemplateId = templateIdArr[templateType]
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("SendSubscribeMessage json.Marshal failed. err is " + err.Error())
		return nil, err
	}

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("SendSubscribeMessage http.Post failed. err is " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("SendSubscribeMessage ioutil.ReadAll failed. err is " + err.Error())
		return nil, err
	}

	var templateResponse SendTemplateResponse
	err = json.Unmarshal(body, &templateResponse)
	if err != nil {
		log.Println("SendSubscribeMessage json.Unmarshal failed. err is " + err.Error())
		return nil, err
	}

	return &templateResponse, nil
}
