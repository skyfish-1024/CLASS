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

type Homework struct {
	ID        int       `form:"ID" gorm:"column:ID" json:"ID,omitempty"`
	ClassMark string    `gorm:"column:ClassMark" json:"ClassMark,omitempty" form:"ClassMark"`
	Poster    string    `form:"Poster" gorm:"column:Poster" json:"Poster,omitempty"`
	Title     string    `form:"Title" gorm:"column:Title" json:"Title,omitempty"`
	Content   string    `form:"Content" gorm:"column:Content" json:"Content,omitempty"`
	Answer    string    `form:"Answer" gorm:"column:Answer" json:"Answer,omitempty"`
	PostTime  time.Time `form:"PostTime" gorm:"column:PostTime" json:"PostTime"`
}

//CreateHomework 发布作业
func CreateHomework(H *Homework) (err error) {
	err = Gorm.Db.Save(H).Error
	if err != nil {
		fmt.Println("发布作业失败err：", err)
		return err
	}
	return nil
}

//deleteHomework 删除作业
func deleteHomework(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&Homework{}).Error
	if err != nil {
		fmt.Println("作业删除失败err：", err)
		return err
	}
	return nil
}

//GetHomework 获取作业
func GetHomework(ClassMark string) (HomeworkS []Homework, err error) {
	err = Gorm.Db.Where("ClassMark=?", ClassMark).Find(&HomeworkS).Error
	if err != nil {
		fmt.Println("作业查询失败err：", err)
		return nil, err
	}
	return HomeworkS, nil
}

//GETHomework GETTopic
func GETHomework(c *gin.Context) {
	c.HTML(http.StatusOK, "homework.html", gin.H{
		"username":  login.User.StuName,
		"ClassMark": c.Param("ClassMark"),
	})
}

//POSTHomework 发布作业
func POSTHomework(c *gin.Context) {
	Poster := login.User.StuName
	H := new(Homework)
	err := c.Bind(H)
	if err != nil {
		fmt.Println("文章数据绑定失败err：", err)
		return
	}
	H.Poster = Poster
	H.PostTime = time.Now()
	H.ClassMark = c.Param("ClassMark")
	err = CreateHomework(H)
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
	homework := new(Homework)
	ID := c.Param("ID")
	id, _ := strconv.Atoi(ID)
	err := Gorm.Db.Where("ID=?", id).First(homework).Error
	if err == nil {
		Answers, err := GetAnswer(id)
		if err != nil {
			fmt.Println("评论回答失败err", err)
		}
		c.HTML(200, "onehomework.html", gin.H{
			"Title":     homework.Title,
			"Content":   homework.Content,
			"PostTime":  homework.PostTime,
			"username":  login.User.StuName,
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
	err = deleteHomework(id)
	if err != nil {
		fmt.Println("删除回答失败err:", err)
		return
	} else {
		//c.JSON(http.StatusOK, gin.H{"结果：": "删除成功！"})
		ClassMark := c.Param("ClassMark")
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark)
	}
}
