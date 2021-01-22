package conf

import (
	"bytes"
	"errors"
	"runtime"

	"github.com/sirupsen/logrus"
)

//CatchPanic 捕获异常并打印到日志中
func CatchPanic(callFrom string) {
	//defer func() {
	if p := recover(); p != nil {
		//fmt.Println("panic recover! p:", p)
		str, ok := p.(string)
		var err error
		if ok {
			err = errors.New(str)
		} else {
			err = errors.New("panic")
		}
		//debug.PrintStack()
		LOG.Self.WithFields(logrus.Fields{
			"err":        err.Error(),
			"p":          p,
			"PanicTrace": string(PanicTrace(4)),
		}).Debug("信息异常 CatchPanic callFrom : " + callFrom)

		/*
			if callFrom == "hub Run()" {
				LOG.Self.Info("信息异常 CatchPanic restart hub ReRun()")
				websocket.WsHub.ReRun() //重启webscoket主逻辑
			}

		*/
	} else {
		LOG.Self.WithFields(logrus.Fields{
			"p": p,
		}).Debug("正常异常 CatchPanic nothing callFrom : " + callFrom)
	}
	//}()
}

//PanicTrace 捕获异常
func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}
