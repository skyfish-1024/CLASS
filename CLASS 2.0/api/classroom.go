package api

import (
	"CLASS/Gorm"
	"CLASS/function"
	"CLASS/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

//GETClassroom 单个课堂页面
func GETClassroom(c *gin.Context) {
	ClassMark := c.Param("ClassMark")
	classroom, err := function.Oneclass(ClassMark)
	if err != nil {
		c.JSON(404, gin.H{"结果:单个获取失败！err:": err})
	}
	homeworks, err := function.GetHomework(ClassMark)
	if err != nil {
		c.JSON(404, gin.H{"结果:作业获取失败！err:": err})
	}
	c.HTML(200, "classroom.html", gin.H{
		"homeworks": homeworks,
		"Teacher":   classroom.Teacher,
		"Describe":  classroom.Describe,
		"ClassMark": classroom.ClassMark,
		"SignNum":   classroom.SignNum,
		"user":      model.User.StuName,
	})
}

//Create 创建课堂的GET请求
func Create(c *gin.Context) {
	c.HTML(200, "createclassroom.html", nil)
}

//CreateClassroom 创建课堂的POST请求
func CreateClassroom(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	classRoom := new(model.Classroom)
	err := c.Bind(classRoom)
	if err != nil {
		fmt.Println("课堂数据绑定失败err：", err)
		return
	}
	classRoom.ClassMark = function.RandomString(5)
	classRoom.CreateTime = time.Now()
	err = Gorm.Db.Save(classRoom).Error
	if err != nil {
		fmt.Println("新增课堂失败err：", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+classRoom.ClassMark)

}

//EnterClassroom 加入课堂
func EnterClassroom(c *gin.Context) {
	ClassMark := c.PostForm("ClassMark")
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark)
}

//SignIn 课堂签到
func SignIn(c *gin.Context) {
	ClassMark := c.Param("ClassMark")
	err := function.Signin(ClassMark)
	if err != nil {
		fmt.Println("签到失败MySQL err:", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark)
}
