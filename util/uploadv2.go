package util

import (
	"bytes"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

type OssUploader struct {
	endPoint   string
	bucket     string
	key        string
	secret     string
	projectTag string
}

func NewOssUploader(projectTag string) *OssUploader {
	up := &OssUploader{}
	up.endPoint = viper.GetString("OssAdminEndPoint")
	up.bucket = viper.GetString(projectTag + ".OssAdminBucket")
	up.key = viper.GetString("OssAdminAccessKey")
	up.secret = viper.GetString("OssAdminSecret")
	up.projectTag = projectTag
	return up
}

func (up *OssUploader) UploadData(data []byte, filePath, name string) (string, error) {
	client, err := oss.New(up.endPoint, up.key, up.secret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(up.bucket)
	if err != nil {
		return "", err
	}
	filePath = genPath(filePath, name)
	log.Printf("UploadFile filePath:%s\n", filePath)
	if err := bucket.PutObject(filePath, bytes.NewReader(data)); err != nil {
		return "", err
	}
	return viper.GetString(up.projectTag+".OssResourceUrl") + "/" + filePath, nil
}
