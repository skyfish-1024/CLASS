package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
)

//CreateAnswer 新增回答
func CreateAnswer(A *model.Answer) (err error) {
	err = Gorm.Db.Save(A).Error
	if err != nil {
		fmt.Println("新增回答失败err：", err)
		return err
	}
	return nil
}

//deleteAnswer 删除回答
func DdeleteAnswer(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&model.Answer{}).Error
	if err != nil {
		fmt.Println("删除回答失败err：", err)
		return err
	}
	return nil
}

//GetAnswer 获取回答
func GetAnswer(id int) (answers []model.Answer, err error) {
	err = Gorm.Db.Where("HomeworkId = ? ", id).Find(&answers).Error
	if err != nil {
		fmt.Println("获取回答失败err：", err)
		return nil, err
	}
	return answers, nil
}
