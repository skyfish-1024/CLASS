package model

import "time"

//课堂
type Classroom struct {
	ID         int       `gorm:"column:ID" json:"ID,omitempty" form:"ID"`
	ClassMark  string    `gorm:"column:ClassMark" json:"ClassMark,omitempty" form:"ClassMark"`
	Teacher    string    `gorm:"column:Teacher" json:"Teacher,omitempty" form:"Teacher"`
	Describe   string    `gorm:"column:Describe" json:"Describe,omitempty" form:"Describe"`
	SignNum    int       `gorm:"column:SignNum" json:"SignNum,omitempty" form:"SignNum"`
	CreateTime time.Time `gorm:"column:CreateTime" json:"CreateTime,omitempty" form:"CreateTime"`
}
