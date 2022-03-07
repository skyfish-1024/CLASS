package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
)

//CreateHomework 发布作业
func CreateHomework(H *model.Homework) (err error) {
	err = Gorm.Db.Save(H).Error
	if err != nil {
		fmt.Println("发布作业失败err：", err)
		return err
	}
	return nil
}

//deleteHomework 删除作业
func DdeleteHomework(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&model.Homework{}).Error
	if err != nil {
		fmt.Println("作业删除失败err：", err)
		return err
	}
	return nil
}

//GetHomework 获取作业
func GetHomework(ClassMark string) (HomeworkS []model.Homework, err error) {
	err = Gorm.Db.Where("ClassMark=?", ClassMark).Find(&HomeworkS).Error
	if err != nil {
		fmt.Println("作业查询失败err：", err)
		return nil, err
	}
	return HomeworkS, nil
}
