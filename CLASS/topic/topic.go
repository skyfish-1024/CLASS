package topic

import (
	"CLASS/Gorm"
	"CLASS/account/login"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

//话题
type topic struct {
	ID       int       `gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	Poster   string    `gorm:"column:Poster" form:"Poster" json:"Poster,omitempty"`       //发布者
	Title    string    `gorm:"column:Title" form:"Title" json:"Title,omitempty"`          //话题标题
	Content  string    `gorm:"column:Content" form:"Content" json:"Content,omitempty"`    //话题内容
	PostTime time.Time `gorm:"column:PostTime" form:"PostTime" json:"PostTime"`           //发布时间
	PointNum int       `gorm:"column:PointNum" form:"PointNum" json:"PointNum,omitempty"` //点赞
}

//新增话题
func CreateTopic(T *topic) (err error) {
	err = Gorm.Db.Save(T).Error
	if err != nil {
		fmt.Println("新增话题失败err：", err)
		return err
	}
	return nil
}

//修改话题&暂未使用
//func EditTopic(T *topic) (err error) {
//	data := make(map[string]interface{})
//	data["ID"] = T.ID
//	data["Poster"] = T.Poster
//	data["content"] = T.Content
//	data["Title"] = T.Title
//	data["PointNum"] = T.PointNum
//	data["PostTime"] = T.PostTime
//	err = Gorm.Db.Model(&topic{}).Where("ID=?", T.ID).Updates(data).Error
//	if err != nil {
//		fmt.Println("修改失败err：", err)
//		return err
//	}
//	return nil
//}

//删除话题
func deleteTopic(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&topic{}).Error
	if err != nil {
		fmt.Println("删除失败err：", err)
		return err
	}
	return nil
}

// PointTopic 话题点赞
func PointTopic(ID int) (err error) {
	err = Gorm.Db.Model(&topic{}).Where("ID=?", ID).Update("PointNum", gorm.Expr("PointNum + 1")).Error
	if err != nil {
		fmt.Println("点赞失败err", err)
		return err
	}
	return nil
}

// GetTopic 获取话题
func GetTopic() (topics []topic, err error) {
	err = Gorm.Db.Find(&topics).Error
	if err != nil {
		fmt.Println("话题查询失败err：", err)
		return nil, err
	}
	return topics, nil
}

// GETTopic 发布话题的GET请求
func GETTopic(c *gin.Context) {
	c.HTML(http.StatusOK, "topic.html", gin.H{"username": login.User.StuName})
}

//POSTTopic 发布话题的POST请求
func POSTTopic(c *gin.Context) {
	T := new(topic)
	err := c.Bind(T)
	if err != nil {
		fmt.Println("话题数据绑定失败err：", err)
		return
	}
	T.Poster = login.User.StuName
	T.PostTime = time.Now()
	err = CreateTopic(T)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"结果：": "发布失败！"})
	}
	//发布成功，返回主页面
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")
}

//ONETopic 单个话题的详情页面
func ONETopic(c *gin.Context) {
	Topic := new(topic)
	ID := c.Param("ID") //获取话题ID
	id, _ := strconv.Atoi(ID)
	err := Gorm.Db.Where("ID=?", id).First(Topic).Error
	if err == nil {
		Comments, err := GetComment(id)
		if err != nil {
			fmt.Println("评论查询失败err", err)
		}
		c.HTML(200, "onetopic.html", gin.H{
			"Title":    Topic.Title,
			"Content":  Topic.Content,
			"PostTime": Topic.PostTime,
			"PointNum": Topic.PointNum,
			"username": login.User.StuName,
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
	err = deleteTopic(id)
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
	err = PointTopic(id)
	if err != nil {
		fmt.Println("点赞失败 MySQL err:", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home") //点赞成功，返回主页面
}
