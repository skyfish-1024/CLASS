package api

import (
	"CLASS/function"
	"CLASS/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddScore 成绩发布页面的GET请求
func AddScore(c *gin.Context) {
	c.HTML(200, "score.html", gin.H{"username": model.User.StuName})
}

//POStScore 成绩添加的POST请求
func POStScore(c *gin.Context) {
	Score := new(model.Score)
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
	err = function.CreateScore(Score)
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
	err := function.DeleteScore(ID)
	if err != nil {
		fmt.Println("成绩删除失败！err：", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/CLASS/home")

}
