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

//评论
type comment struct {
	ID          int       ` gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	TopicId     int       `gorm:"column:TopicId" form:"TopicId" json:"TopicId,omitempty"`    //外键与Topic绑定标识
	Observer    string    `gorm:"column:Observer" form:"Observer" json:"Observer,omitempty"` //评论者
	Comment     string    `gorm:"column:Comment" form:"Comment" json:"Comment,omitempty"`    //评论
	CommentTime time.Time `gorm:"column:CommentTime" form:"CommentTime" json:"CommentTime"`  //评论时间
	PointNum    int       `gorm:"column:PointNum" form:"PointNum" json:"PointNum,omitempty"` //点赞数
}

// CreateComment 新增评论
func CreateComment(T *comment) (err error) {
	err = Gorm.Db.Save(T).Error
	if err != nil {
		fmt.Println("新增话题失败 err：", err)
		return err
	}
	return nil
}

//修改评论&暂未实现
//func EditComment(T *comment) (err error) {
//	data := make(map[string]interface{})
//	data["ID"] = T.ID
//	data["TopicId"] = T.TopicId
//	data["Observer"] = T.Observer
//	data["Comment"] = T.Comment
//	data["PointNum"] = T.PointNum
//	data["CommentTime"] = T.CommentTime
//	err = Gorm.Db.Model(&comment{}).Where("ID=?", T.ID).Updates(data).Error
//	if err != nil {
//		fmt.Println("修改失败err：", err)
//		return err
//	}
//	return nil
//}

//deleteComment 删除评论
func deleteComment(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&comment{}).Error
	if err != nil {
		fmt.Println("删除评论失败err：", err)
		return err
	}
	return nil
}

// PointComment 评论点赞
func PointComment(ID int) (err error) {
	err = Gorm.Db.Model(&comment{}).Where("ID=?", ID).Update("PointNum", gorm.Expr("PointNum + 1")).Error
	if err != nil {
		fmt.Println("评论点赞失败err", err)
		return err
	}
	return nil
}

// GetComment 获取话题对应评论
func GetComment(id int) (comments []comment, err error) {
	err = Gorm.Db.Where("TopicId = ? ", id).Find(&comments).Error
	if err != nil {
		fmt.Println("评论查询失败err：", err)
		return nil, err
	}
	return comments, nil
}

//POSTComment 发布评论的POST请求
func POSTComment(c *gin.Context) {
	Comment := new(comment)
	err := c.Bind(Comment)
	if err != nil {
		fmt.Println("评论数据绑定失败err：", err)
		return
	}
	Comment.Observer = login.User.StuName
	ID := c.Param("ID") //获取Topic的ID
	id, _ := strconv.Atoi(ID)
	Comment.TopicId = id
	Comment.CommentTime = time.Now()
	err = CreateComment(Comment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"结果：": "发布失败！"})
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/onetopic/"+ID)
}

//DeleteComment 删除评论的POST请求
func DeleteComment(c *gin.Context) {
	ID := c.Param("commentID") //获取评论ID
	id, _ := strconv.Atoi(ID)
	err := deleteComment(id)
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
	err = PointComment(id)
	if err != nil {
		fmt.Println("点赞失败MySQL err:", err)
		return
	}
	TopicId := c.Param("ID")
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/onetopic/"+TopicId)
}
