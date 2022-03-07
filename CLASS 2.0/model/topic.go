package model

import "time"

//话题
type Topic struct {
	ID       int       `gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	Poster   string    `gorm:"column:Poster" form:"Poster" json:"Poster,omitempty"`       //发布者
	Title    string    `gorm:"column:Title" form:"Title" json:"Title,omitempty"`          //话题标题
	Content  string    `gorm:"column:Content" form:"Content" json:"Content,omitempty"`    //话题内容
	PostTime time.Time `gorm:"column:PostTime" form:"PostTime" json:"PostTime"`           //发布时间
	PointNum int       `gorm:"column:PointNum" form:"PointNum" json:"PointNum,omitempty"` //点赞
}
