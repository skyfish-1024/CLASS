package api

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GETUnsubscribe 注销的get请求
func GETUnsubscribe(c *gin.Context) {
	c.HTML(http.StatusOK, "unsubscribe.html", gin.H{"SecQue": model.User.SecQue})
}

//POSTUnsubscribe 注销的POST请求
func POSTUnsubscribe(c *gin.Context) {
	StuName := c.PostForm("StuName")
	StuPassword := c.PostForm("StuPassword")
	Phone := c.PostForm("Phone")
	SecAns := c.PostForm("SecAns")
	if StuName != model.User.StuName {
		fmt.Println("姓名错误！")
	} else if StuPassword != model.User.StuPassword {
		fmt.Println("密码错误！")
	} else if Phone != model.User.Phone {
		fmt.Println("电话号码有误！")
	} else if SecAns != model.User.SecAns {
		fmt.Println("密保错误！")
	} else {
		U := new(model.Student)
		err := Gorm.Db.Where("Phone=?", Phone).Delete(U)
		fmt.Println("注销err", err)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"结果：": "注销失败！"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"结果：": "感谢您的使用！"})
	}
}
