package api

import (
	"CLASS/Gorm"
	"CLASS/function"
	"CLASS/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//GETHomework GETTopic
func GETHomework(c *gin.Context) {
	cookie, _ := c.Cookie("key_cookie")
	c.HTML(http.StatusOK, "homework.html", gin.H{
		"username":  model.Users[cookie],
		"ClassMark": c.Param("ClassMark"),
	})
}

//POSTHomework 发布作业
func POSTHomework(c *gin.Context) {
	Poster := model.User.StuName
	H := new(model.Homework)
	err := c.Bind(H)
	if err != nil {
		fmt.Println("文章数据绑定失败err：", err)
		return
	}
	H.Poster = Poster
	H.PostTime = time.Now()
	H.ClassMark = c.Param("ClassMark")
	err = function.CreateHomework(H)
	if err == nil {
		//c.HTML(http.StatusOK, "home.html", nil)
		//c.JSON(http.StatusOK, gin.H{"outcome:": "ok"})
		ClassMark := c.Param("ClassMark")
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark)
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"结果：": "发布失败！"})
	}
}

//ONEHomework 单个作业详情
func ONEHomework(c *gin.Context) {
	homework := new(model.Homework)
	ID := c.Param("ID")
	id, _ := strconv.Atoi(ID)
	err := Gorm.Db.Where("ID=?", id).First(homework).Error
	if err == nil {
		Answers, err := function.GetAnswer(id)
		if err != nil {
			fmt.Println("评论回答失败err", err)
		}
		c.HTML(200, "onehomework.html", gin.H{
			"Title":     homework.Title,
			"Content":   homework.Content,
			"PostTime":  homework.PostTime,
			"username":  model.User.StuName,
			"ID":        homework.ID,
			"Answers":   Answers,
			"ClassMark": homework.ClassMark,
		})
	} else {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"结果：": "编辑失败！"})
	}
}

//DeleteHomework 删除作业
func DeleteHomework(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("转换IDerr：", err)
	}
	err = function.DdeleteHomework(id)
	if err != nil {
		fmt.Println("删除回答失败err:", err)
		return
	} else {
		//c.JSON(http.StatusOK, gin.H{"结果：": "删除成功！"})
		ClassMark := c.Param("ClassMark")
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark)
	}
}
