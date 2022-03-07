package model

import "time"

//评论
type Comment struct {
	ID          int       ` gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	TopicId     int       `gorm:"column:TopicId" form:"TopicId" json:"TopicId,omitempty"`    //与Topic的绑定标识
	Observer    string    `gorm:"column:Observer" form:"Observer" json:"Observer,omitempty"` //评论者
	Comment     string    `gorm:"column:Comment" form:"Comment" json:"Comment,omitempty"`    //评论
	CommentTime time.Time `gorm:"column:CommentTime" form:"CommentTime" json:"CommentTime"`  //评论时间
	PointNum    int       `gorm:"column:PointNum" form:"PointNum" json:"PointNum,omitempty"` //点赞数
}
