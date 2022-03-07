package model

import "time"

type Answer struct {
	ID         int       `gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	HomeworkId int       `gorm:"column:HomeworkId" form:"HomeworkId" json:"HomeworkId,omitempty"`
	TheAnswer  string    `gorm:"column:TheAnswer" form:"TheAnswer" json:"TheAnswer,omitempty"`
	Student    string    `gorm:"column:Student" form:"Student" json:"Student,omitempty"`
	AnswerTime time.Time `gorm:"column:AnswerTime" form:"AnswerTime" json:"AnswerTime"`
}
