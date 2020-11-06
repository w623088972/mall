package errno

import (
	"log"
)

type Errno struct {
	Code    int               `json:"code"`
	Message map[string]string `json:"message"` //en英语，chs简中，cht繁中
}

func (err Errno) Error() string {
	return err.Message["en"]
}

func DecodeErr(err error, language string) (int, string, string) {
	if err == nil {
		return OK.Code, OK.Message[language], OK.Message["en"]
	}

	switch typed := err.(type) {
	case *Errno:
		log.Println("DecodeErr language is ", language)
		log.Println("DecodeErr typed.Message[language] is ", typed.Message[language])
		return typed.Code, typed.Message[language], typed.Message["en"]
	default:
	}

	return InternalServerError.Code, err.Error(), err.Error()
}
