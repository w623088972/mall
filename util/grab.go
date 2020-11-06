package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

const weChatDirPath = "storage/grab/weChat"

//微信公众号抓取
func WeChatSubscriptionGrab(netUrl string) (string, error) {
	//获取页面信息
	src, err := GetHtml(netUrl)
	if err != nil {
		return "", err
	}

	//获取图片信息
	re := regexp.MustCompile(`<img([^<].+?) />`)
	rc := re.FindAllStringSubmatch(src, -1)
	var imgSrcArr []string
	for _, val := range rc {
		re = regexp.MustCompile(`data-src="([^<].+?)"`)
		img := re.FindAllStringSubmatch(val[1], -1)
		for _, v := range img {
			imgSrcArr = append(imgSrcArr, v[1])
		}
	}

	//判断目录是否存在
	isDirExist := IsExist(weChatDirPath)
	if isDirExist == false {
		//创建目录
		err := os.MkdirAll(weChatDirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	if imgSrcArr != nil {
		for _, val := range imgSrcArr {
			ossImgUrl, err := HandleImg(val)
			if err != nil {
				return "", err
			}
			//替换页面图片地址
			src = strings.Replace(src, `data-src="`+val+`"`, `data-src="`+ossImgUrl+`"`, -1)
		}
	}

	//判断目录是否存在
	isDirExist = IsExist(weChatDirPath)
	if isDirExist {
		//删除本地目录
		err = os.RemoveAll(weChatDirPath)
		if err != nil {
			return "", err
		}
	}

	return src, nil
}

//获取页面信息
func GetHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.")
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error.")
		return "", err
	}

	src := string(body)

	//去除html
	re, _ := regexp.Compile("\\<!DOCTYPE html>|<html>|</html\\>")
	src = re.ReplaceAllString(src, "")

	//去除head
	re, _ = regexp.Compile("\\<head[\\S\\s]+?\\</head\\>")
	src = re.ReplaceAllString(src, "")

	//去除link
	re, _ = regexp.Compile("\\<link[\\S\\s]+?>")
	src = re.ReplaceAllString(src, "")

	//去除style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//替换style
	re, _ = regexp.Compile(`style="`)
	src = re.ReplaceAllString(src, `data-ytwhkj="`)

	return src, nil
}

//下载文件
//ty：1图片，2视频
func DownloadFile(netFileUrl, fileUrl string, ty int) error {
	httpClient := http.DefaultClient
	httpClient.Timeout = time.Second * 150 //设置超时时间
	resp, err := httpClient.Get(netFileUrl)
	if err != nil {
		panic(err)
		return err
	}
	if resp.ContentLength <= 0 {
		log.Println("Destination server does not support breakpoint download.")
	}

	defer resp.Body.Close()

	file, err := os.Create(fileUrl)
	if err != nil {
		panic(err)
		return err
	}
	defer func() {
		file.Close()
	}()

	//写入文件
	written, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\nTotal length: %d", written)

	return nil
}

//处理图片
func HandleImg(imgUrl string) (string, error) {
	//解析val
	u, err := url.Parse(imgUrl)
	if err != nil {
		return "", err
	}

	localFileUrl := ""
	if strings.Index(u.RequestURI(), "?wx_fmt=") > 0 {
		localFileUrl += strings.Replace(u.RequestURI(), "?wx_fmt=", ".", -1)
	} else {
		localFileUrl += strings.Replace(u.RequestURI(), "?", "-", -1)
		//去除特殊字符
		re, _ := regexp.Compile("\\W")
		localFileUrl = re.ReplaceAllString(localFileUrl, "-") + ".gif"
	}
	localFileUrl = weChatDirPath + "/" + strings.Replace(localFileUrl, "/", "-", -1)

	//判断文件是否已上传到oss
	isOssFileExist, ossImgUrl, err := IsOssFileExist(localFileUrl)
	if err != nil {
		return "", err
	}
	if isOssFileExist { //已上传到oss，替换
		return ossImgUrl, nil
	} else { //上传本地图片到oss，并替换
		//下载图片到本地
		err = DownloadFile(imgUrl, localFileUrl, 1)
		if err != nil {
			return "", err
		}
		//上传本地文件到oss
		ossImgUrl, err := UploadFileToOss(localFileUrl)
		if err != nil {
			return "", err
		} else {
			return ossImgUrl, nil
		}
	}
}

//判断文件是否存在OSS
func IsOssFileExist(fileUrl string) (bool, string, error) {
	endPoint := viper.GetString("oss.endPoint")
	accessKey := viper.GetString("oss.accessKey")
	secret := viper.GetString("oss.secret")
	bucketName := viper.GetString("oss.bucket")

	imgUrl := viper.GetString("oss.resourceUrl")
	grabFilePath := viper.GetString("ossPath.grabFilePath")

	//创建OSSClient实例
	client, err := oss.New(endPoint, accessKey, secret)
	if err != nil {
		return false, "", err
	}

	//获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return false, "", err
	}

	//判断文件是否存在
	isExist, err := bucket.IsObjectExist(grabFilePath + "/" + fileUrl)
	if err != nil {
		return false, "", err
	}

	//oss文件路径
	ossImgUrl := ""
	if isExist {
		ossImgUrl = imgUrl + "/" + grabFilePath + "/" + fileUrl
	}

	return isExist, ossImgUrl, nil
}

//上传本地文件到OSS
func UploadFileToOss(fileUrl string) (string, error) {
	endPoint := viper.GetString("oss.endPoint")
	accessKey := viper.GetString("oss.accessKey")
	secret := viper.GetString("oss.secret")
	bucketName := viper.GetString("oss.bucket")

	imgUrl := viper.GetString("oss.resourceUrl")
	grabFilePath := viper.GetString("ossPath.grabFilePath")

	//创建OSSClient实例
	client, err := oss.New(endPoint, accessKey, secret)
	if err != nil {
		return "", err
	}

	//获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	//上传本地文件
	err = bucket.PutObjectFromFile(grabFilePath+"/"+fileUrl, fileUrl)
	if err != nil {
		fmt.Println("UploadFileToOss Error:", err)
		return "", err
	}

	//oss文件路径
	ossImgUrl := imgUrl + "/" + grabFilePath + "/" + fileUrl

	return ossImgUrl, nil
}
