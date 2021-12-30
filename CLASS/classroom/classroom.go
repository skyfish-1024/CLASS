package classroom

import (
	"CLASS/Gorm"
	"CLASS/account/login"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"
)

//课堂
type Classroom struct {
	ID         int       `gorm:"column:ID" json:"ID,omitempty" form:"ID"`
	ClassMark  string    `gorm:"column:ClassMark" json:"ClassMark,omitempty" form:"ClassMark"`
	Teacher    string    `gorm:"column:Teacher" json:"Teacher,omitempty" form:"Teacher"`
	Describe   string    `gorm:"column:Describe" json:"Describe,omitempty" form:"Describe"`
	SignNum    int       `gorm:"column:SignNum" json:"SignNum,omitempty" form:"SignNum"`
	CreateTime time.Time `gorm:"column:CreateTime" json:"CreateTime,omitempty" form:"CreateTime"`
}

//GetClassRoom 获取课堂记录
func GetClassRoom() (classRooms []Classroom, err error) {
	err = Gorm.Db.Find(&classRooms).Error
	if err != nil {
		fmt.Println("课堂查询失败err：", err)
		return nil, err
	}
	return classRooms, nil
}

//oneclass 获取一个课堂
func oneclass(ClassMark string) (calssroom Classroom, err error) {
	err = Gorm.Db.Where("ClassMark=?", ClassMark).First(&calssroom).Error
	if err != nil {
		fmt.Println("单个课堂获取失败err:", err)
		return calssroom, err
	}
	return calssroom, nil
}

//deleteclassroom 删除课堂
func deleteclassroom(ClassMark string) (err error) {
	err = Gorm.Db.Where("ClassMark=?", ClassMark).Delete(&Classroom{}).Error
	if err != nil {
		fmt.Println("删除失败err：", err)
		return err
	}
	return nil
}

//Signin 课堂签到
func Signin(ClassMark string) (err error) {
	err = Gorm.Db.Model(&Classroom{}).Where("ClassMark=?", ClassMark).Update("SignNum", gorm.Expr("SignNum + 1")).Error
	if err != nil {
		fmt.Println("签到失败err", err)
		return err
	}
	return nil
}

//生成随机字符串，作为课堂暗号
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}

//GETClassroom 单个课堂页面
func GETClassroom(c *gin.Context) {
	ClassMark := c.Param("ClassMark")
	classroom, err := oneclass(ClassMark)
	if err != nil {
		c.JSON(404, gin.H{"结果:单个获取失败！err:": err})
	}
	homeworks, err := GetHomework(ClassMark)
	if err != nil {
		c.JSON(404, gin.H{"结果:作业获取失败！err:": err})
	}
	c.HTML(200, "classroom.html", gin.H{
		"homeworks": homeworks,
		"Teacher":   classroom.Teacher,
		"Describe":  classroom.Describe,
		"ClassMark": classroom.ClassMark,
		"SignNum":   classroom.SignNum,
		"user":      login.User.StuName,
	})
}

//Create 创建课堂的GET请求
func Create(c *gin.Context) {
	c.HTML(200, "createclassroom.html", nil)
}

//CreateClassroom 创建课堂的POST请求
func CreateClassroom(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	classRoom := new(Classroom)
	err := c.Bind(classRoom)
	if err != nil {
		fmt.Println("课堂数据绑定失败err：", err)
		return
	}
	classRoom.ClassMark = randomString(5)
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
	err := Signin(ClassMark)
	if err != nil {
		fmt.Println("签到失败MySQL err:", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/classroom/"+ClassMark)
}
