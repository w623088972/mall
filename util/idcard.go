package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type idCardCertResp struct {
	Status     string `json:"status"`
	Msg        string `json:"msg"`
	IdCard     string `json:"idCard"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Area       string `json:"area"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Prefecture string `json:"prefecture"`
	Birthday   string `json:"birthday"`
	AddrCode   string `json:"addrCode"`
	LastCode   string `json:"lastCode"`
}

type CertInfo struct {
	Name   string `json:"name"`
	IDCard string `json:"idcard"`
	Code   int    `json:"code"`
}

//姓名身份证号认证
func VerifyIdCard(idCard string, name string) (bool, *CertInfo, error) {
	appCode := viper.GetString("idAppCode")
	headName := viper.GetString("idHeadName")
	apiUrl := viper.GetString("idApiUrl")

	param := url.Values{}

	//配置请求参数,方法内部已处理URLEncode问题,中文参数可以直接传参
	param.Set("idCard", idCard) //身份证号
	param.Set("name", name)     //姓名

	idCardResp := idCardCertResp{}
	data, err := GetWithHead(apiUrl, param, headName, appCode)
	if err != nil {
		return false, nil, err
	} else {
		json.Unmarshal(data, &idCardResp)
	}
	if idCardResp.Status == "01" {
		return true, nil, nil
	}
	if idCardResp.Status == "205" {
		return false, &CertInfo{
			Name:   name,
			IDCard: idCard,
			Code:   1,
		}, nil
	}

	return false, &CertInfo{
		Name:   name,
		IDCard: idCard,
		Code:   2,
	}, nil
}

//姓名身份证号认证
func IdCardCert(idCard string, name string) (bool, *idCardCertResp, error) {
	isRight := false

	appCode := viper.GetString("idAppCode")
	headName := viper.GetString("idHeadName")
	apiUrl := viper.GetString("idApiUrl")

	param := url.Values{}

	//配置请求参数,方法内部已处理URLEncode问题,中文参数可以直接传参
	param.Set("idCard", idCard) //身份证号
	param.Set("name", name)     //姓名

	var idCardResp idCardCertResp
	data, err := GetWithHead(apiUrl, param, headName, appCode)
	if err != nil {
		return isRight, nil, err
	} else {
		json.Unmarshal(data, &idCardResp)
		//		var netReturn map[string]interface{}
		//		json.Unmarshal(data, &netReturn)
		//		log.Println(netReturn)
		//		for k, v := range netReturn {
		//			log.Println(k, ":", v, "")
		//		}
	}
	if idCardResp.Status == "01" {
		isRight = true
	}

	return isRight, &idCardResp, err
}

func GetWithHead(apiUrl string, params url.Values, headName string, headValue string) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiUrl)
	if err != nil {
		log.Println("解析url错误:\r\n", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", Url.String(), nil)
	req.Header.Add(headName, headValue)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
