package main

import (
	"CLASS/Gorm"
	"CLASS/api"
	"CLASS/chat"
	"CLASS/function"
	"CLASS/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建用户
	function.CreateUser()
	model.Users = make(map[string]model.Student, 0)

	go chat.CHAT()
	//连接数据库
	err := Gorm.NewDB()
	if err != nil {
		panic(err)
	}
	defer Gorm.Db.Close()
	r := gin.Default()
	//加载静态文件
	r.LoadHTMLGlob("./statics/templates/*")
	r.StaticFS("/statics", http.Dir("./statics"))

	//登录
	r.GET("/CLASS/login", api.GETLogin)
	r.POST("/CLASS/login", api.POSTLogin)

	//注册
	r.GET("/CLASS/register", api.GETRegister)
	r.POST("/CLASS/register", api.POSTRegister)

	//找回密码
	r.GET("/CLASS/back", api.GETback)
	r.POST("/CLASS/back1", api.POSTback1)
	r.POST("/CLASS/back2", api.POSTback2)

	r.Use(function.AuthMiddleWare())
	{
		//注销
		r.GET("/CLASS/unsubscribe", api.GETUnsubscribe)
		r.POST("/CLASS/unsubscribe", api.POSTUnsubscribe)

		//主页
		r.GET("/CLASS/home", api.GETHome)
		//退出登录
		r.POST("/CLASS/out", api.OUT)

		//话题
		r.GET("/CLASS/publishtopic", api.GETTopic)      //发布话题
		r.POST("/CLASS/publishtopic", api.POSTTopic)    //提交发布
		r.GET("/CLASS/onetopic/:ID", api.ONETopic)      //话题详情页面
		r.POST("/CLASS/onetopic/:ID", api.DeleteTopic)  //删除话题
		r.POST("/CLASS/onetopic/:ID/point", api.PointT) //话题点赞

		//评论
		r.POST("/CLASS/onetopic/:ID/comment", api.POSTComment)              //提交发布
		r.POST("/CLASS/onetopic/:ID/comment/:commentID", api.DeleteComment) //删除评论
		r.POST("/CLASS/onetopic/:ID/comment/:commentID/point", api.PointC)  //评论点赞

		//文件
		//限制上传最大尺寸
		r.MaxMultipartMemory = 8 << 20
		r.POST("/CLASS/files/upload", api.UploadFile) //上传文件
		r.POST("/CLASS/files/:file", api.Download)    //下载文件

		//成绩
		r.GET("/CLASS/Addscore", api.AddScore)
		r.POST("/CLASS/Addscore", api.POStScore)               //添加
		r.POST("/CLASS/Deletescore/:ScoreId", api.Deletescore) //删除

		//课堂
		r.GET("/CLASS/classroom/:ClassMark", api.GETClassroom) //进入课堂
		r.GET("/CLASS/classroom/create", api.Create)           //创建课堂
		r.POST("/CLASS/classroom/:ClassMark", api.CreateClassroom)
		r.POST("/CLASS/classroom/enter", api.EnterClassroom)     //主页面进入课堂
		r.POST("/CLASS/classroom/:ClassMark/signIn", api.SignIn) //课堂签到

		//作业
		r.GET("/CLASS/classroom/:ClassMark/publishhomework", api.GETHomework)     //发布作业
		r.POST("/CLASS/classroom/:ClassMark/publishhomework", api.POSTHomework)   //提交发布
		r.GET("/CLASS/classroom/:ClassMark/onehomework/:ID", api.ONEHomework)     //作业详情
		r.POST("/CLASS/classroom/:ClassMark/onehomework/:ID", api.DeleteHomework) //删除作业

		//回答
		r.POST("/CLASS/classroom/onehomework/:ID/answer", api.POSTAnswer)             //POST请求回答
		r.POST("/CLASS/classroom/onehomework/:ID/answer/:answerID", api.DeleteAnswer) //删除回答

		r.Run(":9090")

	}
}
