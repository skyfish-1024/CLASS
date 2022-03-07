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

//POSTAnswer 发布回答的POST请求
func POSTAnswer(c *gin.Context) {
	Answer := new(model.Answer)
	err := c.Bind(Answer)
	if err != nil {
		fmt.Println("评论数据绑定失败err：", err)
		return
	}
	Answer.Student = model.User.StuName
	ID := c.Param("ID")
	id, _ := strconv.Atoi(ID)
	Answer.HomeworkId = id
	Answer.AnswerTime = time.Now()
	err = function.CreateAnswer(Answer)
	if err == nil {
		//c.HTML(http.StatusOK, "home.html", nil)
		//c.JSON(http.StatusOK, gin.H{"结果：": "评论成功！"})
		ClassMark := c.Param("ClassMark")
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark+"/onehomework/"+ID)
	} else {
		c.JSON(http.StatusOK, gin.H{"结果：": "提交失败！"})
	}
}

//DeleteAnswer 删除回答
func DeleteAnswer(c *gin.Context) {
	ID := c.Param("answerID")
	id, _ := strconv.Atoi(ID)
	err := function.DdeleteAnswer(id)
	if err != nil {
		fmt.Println("删除话题失败err:", err)
		return
	} else {
		HomeworkId := c.Param("ID")
		ClassMark := c.Param("ClassMark")
		//c.JSON(http.StatusOK, gin.H{"结果：": "删除成功！"})
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark+"/onehomework/"+HomeworkId)
	}
}
