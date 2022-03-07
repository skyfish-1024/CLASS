package api

import (
	"CLASS/function"
	account2 "CLASS/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GETback 找回密码的GET请求
func GETback(c *gin.Context) {
	c.HTML(http.StatusOK, "back1.html", nil)
}

//POSTback1 找回密码的POST请求1
func POSTback1(c *gin.Context) {
	StuName := c.PostForm("StuName")
	Phone := c.PostForm("Phone")
	err := function.BackCheck(StuName, Phone)
	if err != nil {
		c.JSON(200, gin.H{"结果：": err})
		return
	} else {
		c.HTML(http.StatusOK, "back2.html", gin.H{"SecQue": account2.User.SecQue})
	}
	return
}

//POSTback2 找回密码的POST请求2
func POSTback2(c *gin.Context) {
	SecAns := c.PostForm("SecAns")
	fmt.Println("SecAns:", SecAns)
	if SecAns != account2.User.SecAns {
		c.JSON(200, gin.H{"结果：": errors.New("密保答案错误！")})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"所有信息": account2.User})
	}
}
