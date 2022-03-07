package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//RegisterCheck 注册查验
func RegisterCheck(User *model.Student) (err error) {
	//检查输入是否合法
	if User.StuName == "" {
		fmt.Println("姓名不能为空")
		return errors.New("姓名不能为空")
	} else if User.Phone == "" {
		fmt.Println("电话号码不能为空")
		return errors.New("电话号码不能为空")
	} else if User.StuPassword == "" {
		fmt.Println("密码不能为空")
		return errors.New("密码不能为空")
	}
	//检查用户是否已经存在
	err = Gorm.Db.Where("Phone=?", User.Phone).Take(User).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = Gorm.Db.Save(User).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("电话号码已被使用")
		return errors.New("电话号码已被使用")
	}
	return
}
