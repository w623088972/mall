package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"log"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Pagination struct {
	PageNum int `json:"pageNum"`
	Total   int `json:"total"`
}

func NewPagination(page int, total int) *Pagination {
	return &Pagination{
		PageNum: page,
		Total:   total,
	}
}
func ComputeOffsetAndLimit(page, limit string) (int, int) {
	nPage := 0
	nLimit := 10
	if limit != "" {
		nLimit, _ = strconv.Atoi(limit)
	}
	if page != "" {
		nPage, _ = strconv.Atoi(page)
		if nPage > 0 {
			nPage = (nPage - 1) * nLimit
		}
	}
	return nPage, nLimit
}

func EncryptPassword(password string) string {
	data := md5.Sum([]byte(password))
	encryptPassword := make([]byte, 16)
	for i, v := range data {
		encryptPassword[i] = v
	}
	return strings.ToUpper(hex.EncodeToString(encryptPassword))
}

//BindFormV2 绑定表单, all若为false则忽略值为空的键
//PUT方法会提交一些值为空的键，可能会将某些记录字段置空
func BindFormV2(c *gin.Context, param interface{}, all bool) map[string]interface{} {
	typeInfo := reflect.TypeOf(param).Elem()
	value := reflect.ValueOf(param).Elem()
	kv := make(map[string]interface{}, 0)
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("form")
		formV := c.PostForm(tag)
		//是否忽略值为空的字段
		if formV == "" && !all {
			continue
		}
		fieldType := field.Type
		fieldV := value.Field(i)
		switch fieldType.Kind() {
		case reflect.Int:
			v, err := strconv.Atoi(formV)
			if err != nil {
				log.Printf("BindForm failed. param:%v key:%s\n", param, formV)
				continue
			}
			kv[tag] = v
			fieldV.SetInt(int64(v))
		case reflect.String:
			if formV == "undefined" {
				continue
			}
			fieldV.SetString(formV)
			kv[tag] = formV
		case reflect.Float64:
			d, _ := decimal.NewFromString(formV)
			floatV, _ := d.Float64()
			fieldV.SetFloat(floatV)
			kv[tag] = d
		}

	}
	return kv
}

//BindForm 绑定form表单。有些结构字段类型跟表单字段类型不匹配。要特殊处理
//BindForm 会忽略表单值为空的键
func BindForm(c *gin.Context, param interface{}) map[string]interface{} {
	typeInfo := reflect.TypeOf(param).Elem()
	value := reflect.ValueOf(param).Elem()
	kv := make(map[string]interface{}, 0)
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("form")
		formV := c.PostForm(tag)
		if formV == "" {
			continue
		}
		fieldType := field.Type
		fieldV := value.Field(i)
		switch fieldType.Kind() {
		case reflect.Int:
			v, err := strconv.Atoi(formV)
			if err != nil {
				log.Printf("BindForm failed. param:%v key:%s\n", param, formV)
				continue
			}
			kv[tag] = v
			fieldV.SetInt(int64(v))
		case reflect.String:
			if formV == "undefined" {
				continue
			}
			fieldV.SetString(formV)
			kv[tag] = formV
		case reflect.Float64:
			d, _ := decimal.NewFromString(formV)
			floatV, _ := d.Float64()
			fieldV.SetFloat(floatV)
			kv[tag] = d
		}

	}
	return kv
}

func BindQuery(c *gin.Context, param interface{}) {
	typeInfo := reflect.TypeOf(param).Elem()
	value := reflect.ValueOf(param).Elem()
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("form")
		param := c.Query(tag)
		if param == "" {
			continue
		}
		fieldType := field.Type
		fieldV := value.Field(i)
		switch fieldType.Kind() {
		case reflect.Int:
			v, _ := strconv.Atoi(param)
			fieldV.SetInt(int64(v))
		case reflect.String:
			fieldV.SetString(param)
		}
	}
}

func AssignByBitOrAfter(ids string) int {
	id := 0
	idArr := strings.Split(strings.Trim(ids, ","), ",")
	for _, val := range idArr {
		tmp, _ := strconv.Atoi(val)
		id |= tmp
	}

	return id
}

func ParseByBitOrAfter(id int) string {
	arr := []string{}
	for i := 0; i < id; i++ {
		tmp := pow(2, i)
		if id&tmp == tmp {
			arr = append(arr, strconv.Itoa(tmp))
		}
	}

	str := strings.Join(arr, ",")

	return str
}

func pow(x float64, n int) int {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}

	return int(result)
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	//向右移动一位
	result := calPow(x, n>>1)
	result *= result

	//如果n是奇数
	if n&1 == 1 {
		result *= x
	}

	return result
}

func Ceil2(f float64, n int) float64 {
	unit := float64(pow(10, n))
	return math.Ceil(f*unit) / unit
}

//生成随机字符串
func GetRandomString(num int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
