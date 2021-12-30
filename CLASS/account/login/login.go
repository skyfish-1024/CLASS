package login

import (
	"CLASS/Gorm"
	"CLASS/account/user/student"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var User *student.Student

//初始化user
func CreateUser() {
	User = new(student.Student)
	User.StuName = "游客"
}

//LoginCheck 登录账号密码验证
func LoginCheck(Phone string, password string) (err error) {
	err = Gorm.Db.Where("Phone=?", Phone).Take(User).Error
	if err != nil {
		err = errors.New("该用户不存在！")
		fmt.Println(err)
		return err
	} else if User.StuPassword != password {
		err = errors.New("密码错误！")
		fmt.Println(err)
		return err
	} else {
		//token,_:=SetToken(username)
		//fmt.Println("token:",token)
		fmt.Println("登录成功！")
	}
	return
}

//GETLogin 登录的GET请求
func GETLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//POSTLogin 登录的POST请求
func POSTLogin(c *gin.Context) {
	//获取form参数
	Phone := c.PostForm("Phone")
	Password := c.PostForm("password")
	err := LoginCheck(Phone, Password) //检查用户名和密码是否正确
	if err != nil {
		c.String(200, "结果：", err)
	} else {
		//登录成功，设置中间件
		c.SetCookie("key_cookie", Phone, 3600, "/", "http://127.0.0.1:9090/CLASS", false, false)
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")
	}
}
