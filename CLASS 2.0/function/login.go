package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"errors"
	"fmt"
)

//初始化user
func CreateUser() {
	model.User = new(model.Student)
	model.User.StuName = "游客"
}

//LoginCheck 登录账号密码验证
func LoginCheck(Phone string, password string) (err error) {
	err = Gorm.Db.Where("Phone=?", Phone).Take(model.User).Error
	if err != nil {
		err = errors.New("该用户不存在！")
		fmt.Println(err)
		return err
	} else if model.User.StuPassword != password {
		err = errors.New("密码错误！")
		fmt.Println(err)
		return err
	} else {
		//token,_:=SetToken(username)
		//fmt.Println("token:",token)
		fmt.Println("登录成功！")
	}
	return
}
