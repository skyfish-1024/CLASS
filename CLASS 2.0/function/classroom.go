package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
)

//GetClassRoom 获取课堂记录
func GetClassRoom() (classRooms []model.Classroom, err error) {
	err = Gorm.Db.Find(&classRooms).Error
	if err != nil {
		fmt.Println("课堂查询失败err：", err)
		return nil, err
	}
	return classRooms, nil
}

//oneclass 获取一个课堂
func Oneclass(ClassMark string) (calssroom model.Classroom, err error) {
	err = Gorm.Db.Where("ClassMark=?", ClassMark).First(&calssroom).Error
	if err != nil {
		fmt.Println("单个课堂获取失败err:", err)
		return calssroom, err
	}
	return calssroom, nil
}

//Deleteclassroom 删除课堂
func Deleteclassroom(ClassMark string) (err error) {
	err = Gorm.Db.Where("ClassMark=?", ClassMark).Delete(&model.Classroom{}).Error
	if err != nil {
		fmt.Println("删除失败err：", err)
		return err
	}
	return nil
}

//Signin 课堂签到
func Signin(ClassMark string) (err error) {
	err = Gorm.Db.Model(&model.Classroom{}).Where("ClassMark=?", ClassMark).Update("SignNum", gorm.Expr("SignNum + 1")).Error
	if err != nil {
		fmt.Println("签到失败err", err)
		return err
	}
	return nil
}

//生成随机字符串，作为课堂暗号
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}
