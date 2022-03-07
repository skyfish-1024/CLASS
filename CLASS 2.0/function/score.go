package function

import (
	"CLASS/Gorm"
	"CLASS/model"
	"fmt"
)

//CreateScore 新增成绩
func CreateScore(S *model.Score) (err error) {
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
func DeleteScore(ID int) (err error) {
	err = Gorm.Db.Where("ID=?", ID).Delete(&model.Score{}).Error
	if err != nil {
		fmt.Println("成绩删除失败err：", err)
		return err
	}
	return nil
}

//GetScore 获取成绩
func GetScore() (scores []model.Score, err error) {
	err = Gorm.Db.Find(&scores).Error
	if err != nil {
		fmt.Println("成绩查询失败err：", err)
		return nil, err
	}
	return scores, nil
}
