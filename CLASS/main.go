package main

import (
	"CLASS/Gorm"
	"CLASS/account/back"
	"CLASS/account/cookie"
	"CLASS/account/login"
	"CLASS/account/register"
	"CLASS/account/unsubscribe"
	"CLASS/chat"
	"CLASS/classroom"
	"CLASS/file"
	"CLASS/home"
	"CLASS/score"
	"CLASS/topic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建用户
	login.CreateUser()

	go chat.CHAT()
	//连接数据库
	err := Gorm.NewDB()
	if err != nil {
		panic(err)
	}
	defer Gorm.Db.Close()
	r := gin.Default()
	//加载静态文件
	r.LoadHTMLGlob("./account/templates/*")
	r.StaticFS("/account", http.Dir("./account"))

	//登录
	r.GET("/CLASS/login", login.GETLogin)
	r.POST("/CLASS/login", login.POSTLogin)

	//注册
	r.GET("/CLASS/register", register.GETRegister)
	r.POST("/CLASS/register", register.POSTRegister)

	//找回密码
	r.GET("/CLASS/back", back.GETback)
	r.POST("/CLASS/back1", back.POSTback1)
	r.POST("/CLASS/back2", back.POSTback2)

	r.Use(cookie.AuthMiddleWare())
	{
		//注销
		r.GET("/CLASS/unsubscribe", unsubscribe.GETUnsubscribe)
		r.POST("/CLASS/unsubscribe", unsubscribe.POSTUnsubscribe)

		//主页
		r.GET("/CLASS/home", home.GETHome)
		//退出登录
		r.POST("/CLASS/out", home.OUT)

		//话题
		r.GET("/CLASS/publishtopic", topic.GETTopic)      //发布话题
		r.POST("/CLASS/publishtopic", topic.POSTTopic)    //提交发布
		r.GET("/CLASS/onetopic/:ID", topic.ONETopic)      //话题详情页面
		r.POST("/CLASS/onetopic/:ID", topic.DeleteTopic)  //删除话题
		r.POST("/CLASS/onetopic/:ID/point", topic.PointT) //话题点赞

		//评论
		r.POST("/CLASS/onetopic/:ID/comment", topic.POSTComment)              //提交发布
		r.POST("/CLASS/onetopic/:ID/comment/:commentID", topic.DeleteComment) //删除评论
		r.POST("/CLASS/onetopic/:ID/comment/:commentID/point", topic.PointC)  //评论点赞

		//文件
		//限制上传最大尺寸
		r.MaxMultipartMemory = 8 << 20
		r.POST("/CLASS/files/upload", file.UploadFile) //上传文件
		r.POST("/CLASS/files/:file", file.Download)    //下载文件

		//成绩
		r.GET("/CLASS/Addscore", score.AddScore)
		r.POST("/CLASS/Addscore", score.POStScore)               //添加
		r.POST("/CLASS/Deletescore/:ScoreId", score.Deletescore) //删除

		//课堂
		r.GET("/CLASS/classroom/:ClassMark", classroom.GETClassroom) //进入课堂
		r.GET("/CLASS/classroom/create", classroom.Create)           //创建课堂
		r.POST("/CLASS/classroom/:ClassMark", classroom.CreateClassroom)
		r.POST("/CLASS/classroom/enter", classroom.EnterClassroom)     //主页面进入课堂
		r.POST("/CLASS/classroom/:ClassMark/signIn", classroom.SignIn) //课堂签到

		//作业
		r.GET("/CLASS/classroom/:ClassMark/publishhomework", classroom.GETHomework)     //发布作业
		r.POST("/CLASS/classroom/:ClassMark/publishhomework", classroom.POSTHomework)   //提交发布
		r.GET("/CLASS/classroom/:ClassMark/onehomework/:ID", classroom.ONEHomework)     //作业详情
		r.POST("/CLASS/classroom/:ClassMark/onehomework/:ID", classroom.DeleteHomework) //删除作业

		//回答
		r.POST("/CLASS/classroom/onehomework/:ID/answer", classroom.POSTAnswer)             //POST请求回答
		r.POST("/CLASS/classroom/onehomework/:ID/answer/:answerID", classroom.DeleteAnswer) //删除回答

		r.Run(":9090")

	}
}
