package function

import (
	"fmt"
	"io/ioutil"
)

// GetAllFile 获取文件信息
func GetAllFile(pathname string) (s []string, err error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
