package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

//UploadFile 上传文件
func UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有文件
	files := form.File["files"]
	// 遍历所有文件
	for _, file := range files {
		// 逐个存
		if err := c.SaveUploadedFile(file, "./files/"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")
	//c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}

//Download 下载文件
func Download(c *gin.Context) {
	name := c.Param("file")
	//fmt.Println(name)
	filename := "./files/" + name
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	c.File("./" + file.Name())
}
