package back

import (
	"CLASS/Gorm"
	"CLASS/account/login"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//BackCheck 密保找回
func BackCheck(StuName string, Phone string) (err error) {
	err = Gorm.Db.Where("Phone=?", Phone).Take(login.User).Error
	if err != nil {
		fmt.Println(err)
		errors.New("手机号不存在！")
		return err
	} else if login.User.StuName != StuName {
		fmt.Println(login.User.StuName)
		err = errors.New("学生姓名错误！")
		fmt.Println(err)
		return err
	}
	return nil
}

//GETback 找回密码的GET请求
func GETback(c *gin.Context) {
	c.HTML(http.StatusOK, "back1.html", nil)
}

//POSTback1 找回密码的POST请求1
func POSTback1(c *gin.Context) {
	StuName := c.PostForm("StuName")
	Phone := c.PostForm("Phone")
	err := BackCheck(StuName, Phone)
	if err != nil {
		c.JSON(200, gin.H{"结果：": err})
		return
	} else {
		c.HTML(http.StatusOK, "back2.html", gin.H{"SecQue": login.User.SecQue})
	}
	return
}

//POSTback2 找回密码的POST请求2
func POSTback2(c *gin.Context) {
	SecAns := c.PostForm("SecAns")
	fmt.Println("SecAns:", SecAns)
	if SecAns != login.User.SecAns {
		c.JSON(200, gin.H{"结果：": errors.New("密保答案错误！")})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"所有信息": login.User})
	}
}
