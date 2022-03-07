package api

import (
	"CLASS/function"
	account2 "CLASS/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GETRegister 注册GET
func GETRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

//POSTRegister 注册POST
func POSTRegister(c *gin.Context) {
	U := new(account2.Student)
	c.Bind(U)
	err := function.RegisterCheck(U)
	if err == nil {
		c.HTML(http.StatusOK, "login.html", nil)
	} else { //若存在，则不可以注册
		c.HTML(http.StatusOK, "register.html", gin.H{
			"outcome": err,
		})
	}
}
