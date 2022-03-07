package api

import (
	"CLASS/function"
	"CLASS/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//POSTComment 发布评论的POST请求
func POSTComment(c *gin.Context) {
	Comment := new(model.Comment)
	err := c.Bind(Comment)
	if err != nil {
		fmt.Println("评论数据绑定失败err：", err)
		return
	}
	Comment.Observer = model.User.StuName
	ID := c.Param("ID") //获取Topic的ID
	id, _ := strconv.Atoi(ID)
	Comment.TopicId = id
	Comment.CommentTime = time.Now()
	err = function.CreateComment(Comment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"结果：": "发布失败！"})
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/onetopic/"+ID)
}

//DeleteComment 删除评论的POST请求
func DeleteComment(c *gin.Context) {
	ID := c.Param("commentID") //获取评论ID
	id, _ := strconv.Atoi(ID)
	err := function.DeleteComment(id)
	if err != nil {
		fmt.Println("删除话题失败err:", err)
		return
	} else {
		//c.JSON(http.StatusOK, gin.H{"结果：": "删除成功！"})
		TopicId := c.Param("ID")
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/onetopic/"+TopicId)
	}
}

//PointC 评论点赞的POST请求
func PointC(c *gin.Context) {
	ID := c.Param("commentID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("点赞 评论ID 转换失败 err:", err)
		return
	}
	err = function.PointComment(id)
	if err != nil {
		fmt.Println("点赞失败MySQL err:", err)
		return
	}
	TopicId := c.Param("ID")
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/onetopic/"+TopicId)
}
