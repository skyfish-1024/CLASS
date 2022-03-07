package model

//分数构成
type Score struct {
	ID                int    `gorm:"column:ID" form:"ID" json:"ID,omitempty"`
	ExaminationName   string `gorm:"column:ExaminationName" form:"ExaminationName" json:"ExaminationName,omitempty"`
	ExaminationNumber int    `gorm:"column:ExaminationNumber" form:"ExaminationNumber" json:"ExaminationNumber,omitempty"`
	StuName           string `gorm:"column:StuName" form:"StuName" json:"StuName,omitempty"`
	Class             string `gorm:"column:Class" form:"Class" json:"Class,omitempty"`
	Phone             string `gorm:"column:Phone" form:"Phone" json:"Phone,omitempty"`
	ChineseScore      int    `gorm:"column:ChineseScore" form:"ChineseScore" json:"ChineseScore,omitempty"`
	MathScore         int    `gorm:"column:MathScore" form:"MathScore" json:"MathScore,omitempty"`
	EnglishScore      int    `gorm:"column:EnglishScore" form:"EnglishScore" json:"EnglishScore,omitempty"`
	TotalScore        int    `gorm:"column:TotalScore" form:"TotalScore" json:"TotalScore,omitempty"`
}
