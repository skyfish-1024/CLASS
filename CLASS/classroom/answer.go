package classroom

import (
	"CLASS/Gorm"
	"CLASS/account/login"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type answer struct {
	ID         int       `gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	HomeworkId int       `gorm:"column:HomeworkId" form:"HomeworkId" json:"HomeworkId,omitempty"`
	TheAnswer  string    `gorm:"column:TheAnswer" form:"TheAnswer" json:"TheAnswer,omitempty"`
	Student    string    `gorm:"column:Student" form:"Student" json:"Student,omitempty"`
	AnswerTime time.Time `gorm:"column:AnswerTime" form:"AnswerTime" json:"AnswerTime"`
}

//CreateAnswer 新增回答
func CreateAnswer(A *answer) (err error) {
	err = Gorm.Db.Save(A).Error
	if err != nil {
		fmt.Println("新增回答失败err：", err)
		return err
	}
	return nil
}

//deleteAnswer 删除回答
func deleteAnswer(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&answer{}).Error
	if err != nil {
		fmt.Println("删除回答失败err：", err)
		return err
	}
	return nil
}

//GetAnswer 获取回答
func GetAnswer(id int) (answers []answer, err error) {
	err = Gorm.Db.Where("HomeworkId = ? ", id).Find(&answers).Error
	if err != nil {
		fmt.Println("获取回答失败err：", err)
		return nil, err
	}
	return answers, nil
}

//POSTAnswer 发布回答的POST请求
func POSTAnswer(c *gin.Context) {
	Answer := new(answer)
	err := c.Bind(Answer)
	if err != nil {
		fmt.Println("评论数据绑定失败err：", err)
		return
	}
	Answer.Student = login.User.StuName
	ID := c.Param("ID")
	id, _ := strconv.Atoi(ID)
	Answer.HomeworkId = id
	Answer.AnswerTime = time.Now()
	err = CreateAnswer(Answer)
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
	err := deleteAnswer(id)
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
