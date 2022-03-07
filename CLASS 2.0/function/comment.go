package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

// CreateComment 新增评论
func CreateComment(T *model.Comment) (err error) {
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
func DeleteComment(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&model.Comment{}).Error
	if err != nil {
		fmt.Println("删除评论失败err：", err)
		return err
	}
	return nil
}

// PointComment 评论点赞
func PointComment(ID int) (err error) {
	err = Gorm.Db.Model(&model.Comment{}).Where("ID=?", ID).Update("PointNum", gorm.Expr("PointNum + 1")).Error
	if err != nil {
		fmt.Println("评论点赞失败err", err)
		return err
	}
	return nil
}

// GetComment 获取话题对应评论
func GetComment(id int) (comments []model.Comment, err error) {
	err = Gorm.Db.Where("TopicId = ? ", id).Find(&comments).Error
	if err != nil {
		fmt.Println("评论查询失败err：", err)
		return nil, err
	}
	return comments, nil
}
