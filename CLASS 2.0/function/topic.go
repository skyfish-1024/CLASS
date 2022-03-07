package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

//新增话题
func CreateTopic(T *model.Topic) (err error) {
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
func DeleteTopic(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&model.Topic{}).Error
	if err != nil {
		fmt.Println("删除失败err：", err)
		return err
	}
	return nil
}

// PointTopic 话题点赞
func PointTopic(ID int) (err error) {
	err = Gorm.Db.Model(&model.Topic{}).Where("ID=?", ID).Update("PointNum", gorm.Expr("PointNum + 1")).Error
	if err != nil {
		fmt.Println("点赞失败err", err)
		return err
	}
	return nil
}

// GetTopic 获取话题
func GetTopic() (topics []model.Topic, err error) {
	err = Gorm.Db.Find(&topics).Error
	if err != nil {
		fmt.Println("话题查询失败err：", err)
		return nil, err
	}
	return topics, nil
}
