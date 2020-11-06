package util

import (
	"fmt"
	"os"
)

//文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	// 或者
	//return err == nil || !os.IsNotExist(err)
	// 或者
	//return !os.IsNotExist(err)
}

//删除文件
func DeleteFile(fileName string) error {
	err := os.Remove(fileName) //删除文件
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("DeleteFile remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
		return err
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("DeleteFile remove OK!")
		return nil
	}
}
