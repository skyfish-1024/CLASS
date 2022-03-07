package function

import (
	"CLASS/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义中间件，用cookie记住登录状态
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验
		cookie, err := c.Cookie("key_cookie")
		if err == nil {
			if _, ok := model.Users[cookie]; ok {
				c.Next()
				return
			}
		}
		//fmt.Println("cookie err:",err)
		//异常，重新登录
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/login")
		c.Abort() //若验证不通过，不再调用后续函数处理
		return
	}
}
