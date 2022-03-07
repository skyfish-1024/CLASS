package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"errors"
	"fmt"
)

//BackCheck 密保找回
func BackCheck(StuName string, Phone string) (err error) {
	err = Gorm.Db.Where("Phone=?", Phone).Take(model.User).Error
	if err != nil {
		fmt.Println(err)
		errors.New("手机号不存在！")
		return err
	} else if model.User.StuName != StuName {
		fmt.Println(model.User.StuName)
		err = errors.New("学生姓名错误！")
		fmt.Println(err)
		return err
	}
	return nil
}
