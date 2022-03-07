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

// GETTopic 发布话题的GET请求
func GETTopic(c *gin.Context) {
	c.HTML(http.StatusOK, "topic.html", gin.H{"username": model.User.StuName})
}

//POSTTopic 发布话题的POST请求
func POSTTopic(c *gin.Context) {
	T := new(model.Topic)
	err := c.Bind(T)
	if err != nil {
		fmt.Println("话题数据绑定失败err：", err)
		return
	}
	T.Poster = model.User.StuName
	T.PostTime = time.Now()
	err = function.CreateTopic(T)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"结果：": "发布失败！"})
	}
	//发布成功，返回主页面
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")
}

//ONETopic 单个话题的详情页面
func ONETopic(c *gin.Context) {
	Topic := new(model.Topic)
	ID := c.Param("ID") //获取话题ID
	id, _ := strconv.Atoi(ID)
	err := Gorm.Db.Where("ID=?", id).First(Topic).Error
	if err == nil {
		Comments, err := function.GetComment(id)
		if err != nil {
			fmt.Println("评论查询失败err", err)
		}
		c.HTML(200, "onetopic.html", gin.H{
			"Title":    Topic.Title,
			"Content":  Topic.Content,
			"Poster":   Topic.Poster,
			"PostTime": Topic.PostTime,
			"PointNum": Topic.PointNum,
			"username": model.User.StuName,
			"ID":       Topic.ID,
			"Comments": Comments,
		})
	} else {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"结果：": "查询失败！"})
	}
}

//DeleteTopic 删除单个话题
func DeleteTopic(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("转换 ID err：", err)
	}
	err = function.DeleteTopic(id)
	if err != nil {
		fmt.Println("删除话题失败err:", err)
		return
	} else {
		//c.JSON(http.StatusOK, gin.H{"结果：": "删除成功！"})
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home") //删除成功，返回主页面
	}
}

//PointT 话题点赞
func PointT(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("点赞 ID 转换 err:", err)
		return
	}
	err = function.PointTopic(id)
	if err != nil {
		fmt.Println("点赞失败 MySQL err:", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home") //点赞成功，返回主页面
}
