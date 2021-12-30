package score

import (
	"CLASS/Gorm"
	"CLASS/account/login"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//分数构成
type score struct {
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

//CreateScore 新增成绩
func CreateScore(S *score) (err error) {
	err = Gorm.Db.Save(S).Error
	if err != nil {
		fmt.Println("新增成绩失败err：", err)
		return err
	}
	return nil
}

//修改成绩&暂未实现
//func EditScore(S *score) (err error) {
//	data := make(map[string]interface{})
//	data["ID"] = S.ID
//	data["Phone"] = S.Phone
//	data["StuName"] = S.StuName
//	data["Class"] = S.Class
//	data["ChineseScore"] = S.ChineseScore
//	data["EnglishScore"] = S.EnglishScore
//	data["MathScore"] = S.MathScore
//	data["ExaminationName"] = S.ExaminationName
//	data["ExaminationNumber"] = S.ExaminationNumber
//	data["TotalScore"] = S.TotalScore
//	err = Gorm.Db.Model(&score{}).Where("ID=?", S.ID).Updates(data).Error
//	if err != nil {
//		fmt.Println("修改失败err：", err)
//		return err
//	}
//	return nil
//}

//deleteScore 删除成绩
func deleteScore(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&score{}).Error
	if err != nil {
		fmt.Println("成绩删除失败err：", err)
		return err
	}
	return nil
}

//GetScore 获取成绩
func GetScore() (scores []score, err error) {
	err = Gorm.Db.Find(&scores).Error
	if err != nil {
		fmt.Println("成绩查询失败err：", err)
		return nil, err
	}
	return scores, nil
}

//AddScore 成绩发布页面的GET请求
func AddScore(c *gin.Context) {
	c.HTML(200, "score.html", gin.H{"username": login.User.StuName})
}

//POStScore 成绩添加的POST请求
func POStScore(c *gin.Context) {
	Score := new(score)
	err := c.Bind(Score)
	if err != nil {
		fmt.Println("成绩绑定失败！err：", err)
		return
	}
	if Score.StuName == " " || Score.ExaminationName == " " || Score.Class == " " || Score.ExaminationNumber == 0 {
		fmt.Println("所填数据不合法")
		c.JSON(http.StatusOK, gin.H{"结果：": "添加失败！", "错误：": "所填数据不合法"})
		return
	}
	err = CreateScore(Score)
	if err == nil {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")
	} else {
		c.JSON(http.StatusOK, gin.H{"结果：": "添加失败！", "错误：": err})
	}

}

//Deletescore 删除分数的POST请求
func Deletescore(c *gin.Context) {
	ScoreId := c.Param("ScoreId")
	ID, _ := strconv.Atoi(ScoreId)
	err := deleteScore(ID)
	if err != nil {
		fmt.Println("成绩删除失败！err：", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")

}
