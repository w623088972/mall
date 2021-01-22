package errno

import (
	"log"
)

type Errno struct {
	Code    int               `json:"resultCode"`
	Message map[string]string `json:"showMessage"` //en英语，chs简中，cht繁中
}

func (err Errno) Error() string {
	return err.Message["en"]
}

func DecodeErr(err error, language string) (int, string, string) {
	if err == nil {
		return OK.Code, OK.Message[language], OK.Message["en"]
	}

	switch typed := err.(type) {
	//	case *Err:
	//		return typed.Code, typed.Message, ""
	case *Errno:
		log.Println("language is ", language)
		log.Println("typed.Message[language] is ", typed.Message[language])
		return typed.Code, typed.Message[language], typed.Message["en"]
	default:
	}

	return InternalServerError.Code, err.Error(), err.Error()
}
