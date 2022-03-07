package model

//学生
type Student struct {
	ID          uint   `gorm:"column:ID" form:"ID" json:"ID,omitempty"`                            //id
	StuName     string `gorm:"column:StuName" form:"StuName" json:"StuName,omitempty"`             //学生姓名&账号
	StuPassword string `gorm:"column:StuPassword" form:"StuPassword" json:"StuPassword,omitempty"` //学生密码
	Phone       string `gorm:"column:Phone" form:"Phone" json:"Phone,omitempty"`                   //电话
	Age         int    `gorm:"column:Age" form:"Age" json:"Age,omitempty"`                         //年龄
	Gender      string `gorm:"column:Gender" form:"Gender" json:"Gender,omitempty"`                //性别
	Class       string `gorm:"column:Class" form:"Class" json:"Class,omitempty"`                   //班级
	SecQue      string `gorm:"column:SecQue" form:"SecQue" json:"SecQue,omitempty"`                //security question 密保问题
	SecAns      string `gorm:"column:SecAns" form:"SecAns" json:"SecAns,omitempty"`                //security answer   密保答案
}

var User *Student
