package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
)

//GetStudents 查询所有学生
func GetStudents() (students []model.Student, err error) {
	err = Gorm.Db.Find(&students).Error
	if err != nil {
		fmt.Println("用户查询失败err：", err)
		return nil, err
	}
	return students, nil
}
