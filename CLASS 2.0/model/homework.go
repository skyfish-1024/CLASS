package model

import "time"

type Homework struct {
	ID        int       `form:"ID" gorm:"column:ID" json:"ID,omitempty"`
	ClassMark string    `gorm:"column:ClassMark" json:"ClassMark,omitempty" form:"ClassMark"`
	Poster    string    `form:"Poster" gorm:"column:Poster" json:"Poster,omitempty"`
	Title     string    `form:"Title" gorm:"column:Title" json:"Title,omitempty"`
	Content   string    `form:"Content" gorm:"column:Content" json:"Content,omitempty"`
	Answer    string    `form:"Answer" gorm:"column:Answer" json:"Answer,omitempty"`
	PostTime  time.Time `form:"PostTime" gorm:"column:PostTime" json:"PostTime"`
}
