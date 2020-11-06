package util

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"path"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	ACLPrivate = 1
	ACLPublic  = 2
)

func UploadImages(projectTag, filePath string, c *gin.Context, keys []string) (map[string]string, error) {
	tmp := make(map[string]string, 0)
	for _, key := range keys {
		name, data, err := readFormFile(c, key)
		if err != nil {
			return nil, err
		}
		downloadUrl, err := uploadFile(projectTag, filePath, name, data)
		if err != nil {
			return nil, err
		}
		tmp[key] = downloadUrl
	}
	return tmp, nil
}

func readFormFile(c *gin.Context, key string) (string, []byte, error) {
	formFile, err := c.FormFile(key)
	if err != nil {
		return "", nil, err
	}

	buf := make([]byte, formFile.Size)
	file, err := formFile.Open()
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	n, err := file.Read(buf)
	if err != nil {
		return "", nil, err
	}
	log.Printf("readFormFile name:%s key:%s size:%d read size:%d\n", formFile.Filename, key, formFile.Size, n)

	return formFile.Filename, buf, nil
}

func genPath(filePath string, name string) string {
	now := time.Now()
	y, m, _ := now.Date()
	items := strings.Split(name, ".")
	n := len(items)
	r := rand.Uint32()
	strNow := now.Format("20060102150405") + fmt.Sprint(r)

	return path.Join(filePath, fmt.Sprint(y), fmt.Sprintf("%02d", m), strNow+"."+items[n-1])
}

//上传文件
func UploadFile(projectTag, path string, videoFile *multipart.FileHeader, acl int) (string, error) {
	buf := make([]byte, videoFile.Size)
	file, err := videoFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	n, err := file.Read(buf)
	if err != nil {
		return "", err
	}
	log.Printf("UploadVideo name:%s size:%d read size:%d\n", videoFile.Filename, videoFile.Size, n)

	ossAcl := oss.ObjectACL(oss.ACLPublicRead)
	if acl == ACLPrivate {
		ossAcl = oss.ObjectACL(oss.ACLPrivate)
	}

	return uploadFile(projectTag, path, videoFile.Filename, buf, ossAcl)
}

//上传文件
func uploadFile(projectTag, filePath string, name string, data []byte, options ...oss.Option) (string, error) {
	endPoint := viper.GetString("oss.endPoint")
	accessKey := viper.GetString("oss.accessKey")
	secret := viper.GetString("oss.secret")
	bucketName := viper.GetString("oss.bucket")

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

	filePath = genPath(filePath, name)
	log.Printf("uploadFile filePath:%s\n", filePath)

	//上传字符串
	if err := bucket.PutObject(filePath, bytes.NewReader(data), options...); err != nil {
		return "", err
	}

	return viper.GetString("oss.resourceUrl") + "/" + filePath, nil
}

//上传视频
func UploadVideo(projectTag, path string, videoFile *multipart.FileHeader, acl int) (string, error) {
	buf := make([]byte, videoFile.Size)
	file, err := videoFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	n, err := file.Read(buf)
	if err != nil {
		return "", err
	}
	log.Printf("UploadVideo name:%s size:%d read size:%d\n", videoFile.Filename, videoFile.Size, n)

	ossAcl := oss.ObjectACL(oss.ACLPublicRead)
	if acl == ACLPrivate {
		ossAcl = oss.ObjectACL(oss.ACLPrivate)
	}

	return uploadFile(projectTag, path, videoFile.Filename, buf, ossAcl)
}

//上传图片
func UploadImage(projectTag, path string, imgFile *multipart.FileHeader) (string, error) {
	buf := make([]byte, imgFile.Size)
	file, err := imgFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	n, err := file.Read(buf)
	if err != nil {
		return "", err
	}
	log.Printf("UploadImage name:%s size:%d read size:%d\n", imgFile.Filename, imgFile.Size, n)

	return uploadFile(projectTag, path, imgFile.Filename, buf)
}

func UpdateFileACL(projectTag, objKey string, acl int) error {
	endPoint := viper.GetString("oss.endPoint")
	accessKey := viper.GetString("oss.accessKey")
	secret := viper.GetString("oss.secret")
	bucketName := viper.GetString("oss.bucket")

	//创建OSSClient实例
	client, err := oss.New(endPoint, accessKey, secret)
	if err != nil {
		return err
	}

	//获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	objAcl := oss.ACLPrivate
	if acl == ACLPublic {
		objAcl = oss.ACLPublicRead
	}

	return bucket.SetObjectACL(objKey, objAcl)
}

//上传本地文件到OSS
func UploadLocalFile(projectTag, ossImgName, fileUrl string) (string, error) {
	endPoint := viper.GetString("oss.endPoint")
	accessKey := viper.GetString("oss.accessKey")
	secret := viper.GetString("oss.secret")
	bucketName := viper.GetString("oss.bucket")
	imgUrl := viper.GetString("oss.resourceUrl")
	path := viper.GetString("ossPath.invitationPath")

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
	err = bucket.PutObjectFromFile(path+"/"+ossImgName, fileUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	//oss文件路径
	ossImgUrl := imgUrl + "/" + path + "/" + ossImgName

	return ossImgUrl, nil
}

func UploadBinaryImg(projectTag, path, name string, data []byte) (string, error) {
	return uploadFile(projectTag, path, name, data)
}

//批量上传图片
func BatchUploadImage(projectTag, filePath, name string, data []byte, options ...oss.Option) (string, error) {
	endPoint := viper.GetString("oss.endPoint")
	accessKey := viper.GetString("oss.accessKey")
	secret := viper.GetString("oss.secret")
	bucketName := viper.GetString("oss.bucket")

	client, err := oss.New(endPoint, accessKey, secret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	filePath = genPath(filePath, name)
	log.Printf("BatchUploadImage filePath:%s\n", filePath)
	if err := bucket.PutObject(filePath, bytes.NewReader(data), options...); err != nil {
		return "", err
	}

	return viper.GetString("oss.resourceUrl") + "/" + filePath, nil
}
