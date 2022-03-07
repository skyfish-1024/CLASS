package api

import (
	"CLASS/function"
	"CLASS/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GETLogin 登录的GET请求
func GETLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//POSTLogin 登录的POST请求
func POSTLogin(c *gin.Context) {
	//获取form参数
	Phone := c.PostForm("Phone")
	Password := c.PostForm("password")
	err := function.LoginCheck(Phone, Password) //检查用户名和密码是否正确
	if err != nil {
		c.String(200, "结果：", err)
	} else {
		model.Users[Phone] = *model.User
		//登录成功，设置中间件
		c.SetCookie("key_cookie", Phone, 3600, "/", "http://127.0.0.1:9090/CLASS", false, false)

		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")
	}
}
