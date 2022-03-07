package api

import (
	"CLASS/function"
	"CLASS/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GETHome 主页面的GET请求
func GETHome(c *gin.Context) {
	classRooms, err := function.GetClassRoom()
	if err != nil {
		c.JSON(404, gin.H{"结果:历史课堂获取失败！err:": err})
	}
	Students, err := function.GetStudents()
	if err != nil {
		c.JSON(404, gin.H{"结果:用户获取失败！err:": err})
	}
	Scores, err := function.GetScore()
	if err != nil {
		c.JSON(404, gin.H{"结果:成绩获取失败！err:": err})
	}
	topics, err := function.GetTopic()
	if err != nil {
		c.JSON(404, gin.H{"结果:话题获取失败！err:": err})
	}
	files, err := function.GetAllFile("./files")
	if err != nil {
		fmt.Println("文件获取失败err：", err)
	}
	cookie, _ := c.Cookie("key_cookie")
	c.HTML(200, "home.html", gin.H{
		"classRooms": classRooms,
		"Students":   Students,
		"Scores":     Scores,
		"Topics":     topics,
		"username":   model.Users[cookie].StuName,
		"files":      files,
	})
}

func OUT(c *gin.Context) {
	model.User = new(model.Student)
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/login")
}
