package util

import (
	"github.com/gomodule/redigo/redis"
	"github.com/goroom/aliyun_sms"
	rd "github.com/goroom/rand"
	"github.com/spf13/viper"
	"log"
	redisDb "myself/mall/redis"
)

type VerifyCodeResponse struct {
	Nation string `form:"nation" json:"nation"` //国家,86中国，其他非中国
	Mobile string `form:"mobile" json:"mobile"`
}

//SendVerifyCode 发送短信数字验证码
func SendVerifyCode(mobile, nation string) (string, error) {
	rand := rd.GetRand()
	randCode := rand.String(6, rd.RST_NUMBER)
	//发送手机验证码
	ParamString := `{"code":"` + randCode + `"}`
	if err := SendAliyunSms(mobile, nation, ParamString); err != nil {
		return "", err
	}

	return randCode, nil
}

//判断验证码是否正确
func CodeVerify(nation, mobile, verifyCode, registerSource string) (bool, error) {
	for _, v := range viper.GetStringSlice("testMobiles") {
		if v == nation+mobile && verifyCode == viper.GetString("testVerifyCode") {
			return true, nil
		}
	}

	//从池里获取连接
	rc := redisDb.RedisClient.Self.Get()
	//用完后将连接放回连接池
	defer func() {
		rc.Close()
		log.Println("CodeVerify end redis ActiveCount:", redisDb.RedisClient.Self.ActiveCount())
	}()
	log.Println("CodeVerify start redis ActiveCount:", redisDb.RedisClient.Self.ActiveCount())

	redisKey := viper.GetString("project.name") + ":cache:token:" + registerSource + ":mobile:" + nation + mobile

	codeRight := false
	var redisCode string
	var redisUse int
	redisInfo, err := redis.Values(rc.Do("HMGET", redisKey, "code", "use"))
	if err != nil {
		return codeRight, err
	}

	_, err = redis.Scan(redisInfo, &redisCode, &redisUse)
	if err != nil {
		return codeRight, err
	}
	log.Printf("redisVerifyCode:%v  verifycode:%v use:%v\n", redisCode, verifyCode, redisUse)
	if redisCode != verifyCode {
		return codeRight, err
	}
	codeRight = true

	//验证码已使用
	_, err = rc.Do("HSET", redisKey, "use", 1) //err未做处理
	log.Println("codeVerify HSET ", err)

	return codeRight, err
}

//阿里短信
func SendAliyunSms(mobile, nation, ParamString string) error {
	var err error
	numbers := mobile
	var aliyun_Sms *aliyun_sms.AliyunSms
	var SmsSignName, SmsTemplateCode string
	if nation != "86" {
		numbers = nation + numbers
		SmsSignName = viper.GetString("sms.smsSignNameNot86")
		SmsTemplateCode = viper.GetString("sms.smsTemplateCodeNot86")
	} else {
		SmsSignName = viper.GetString("sms.smsSignName86")
		SmsTemplateCode = viper.GetString("sms.smsTemplateCode86")
	}

	aliyun_AccessKeyId := viper.GetString("aliyunAccessKeyId")
	aliyun_AccessKeySecret := viper.GetString("aliyunAccessKeySecret")
	aliyun_Sms, err = aliyun_sms.NewAliyunSms(SmsSignName, SmsTemplateCode, aliyun_AccessKeyId, aliyun_AccessKeySecret)
	if err != nil {
		log.Println("SendAliyunSms 阿里云短信服务注册失败", err)
		return err
	}

	err = aliyun_Sms.Send(numbers, ParamString)
	if err != nil {
		log.Println("SendAliyunSms 阿里云短信服务发送失败", err)
		return err
	}

	return err
}
