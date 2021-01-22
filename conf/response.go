package conf

import (
	"github.com/gin-gonic/gin"
	"log"
	"myself/mall/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   int         `json:"count,omitempty"`
}

func SendResponse(c *gin.Context, err error, info string, data interface{}, language string) {
	code, message, eMessage := errno.DecodeErr(err, language)
	if info != "" && eMessage != "OK." {
		log.Println(info + eMessage)
	}
	if data == nil {
		var obj struct{}
		c.JSON(http.StatusOK, Response{
			Code:    code,
			Message: message,
			Data:    obj,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code:    code,
			Message: message,
			Data:    data,
		})
	}
}

func SendResponseWithCount(c *gin.Context, count int, data interface{}, language string) {
	if data == nil {
		var obj struct{}
		c.JSON(http.StatusOK, Response{
			Code:    errno.OK.Code,
			Message: "",
			Data:    obj,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code:    errno.OK.Code,
			Message: "",
			Data:    data,
			Count:   count,
		})
	}
}
