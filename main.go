package main

import (
	"fmt"
	"myYoku/conf"
	"myYoku/controllers"
	"myYoku/daos"
	"myYoku/services"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

var appConf = conf.NewAppConf()

func init() {
	err := ini.MapTo(appConf, "./conf/conf.ini")
	if err != nil {
		fmt.Printf("ini文件读取失败:%v", err)
		os.Exit(1)
	}
}

func InitDB() *gorm.DB {
	addr := appConf.Username + ":" + appConf.Password + "@" + "(" + appConf.DBConf.Address + ")/" + appConf.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(addr)
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		fmt.Printf("创建数据库连接失败:%v", err)
		os.Exit(1)
	}
	sqlDB := db.DB()
	if err != nil {
		fmt.Printf("数据库设置失败:%v", err)
		os.Exit(1)
	}
	sqlDB.SetMaxIdleConns(10)           //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          //最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //设置连接空闲超时
	return db
}

func main() {
	db := InitDB()
	fmt.Println("数据连接成功")

	/* 初始化工具类 */
	utils := services.NewUtils(appConf.Md5Key)

	/* 初始路由*/
	r := gin.Default()
	/* 初始化 User*/
	userDao := daos.NewUserDao(db)
	userService := services.NewUserService(userDao, utils)
	userController := controllers.NewUserController(r, "", userService)
	userController.InitUserController()

	/* 初始化 Advert*/
	advertDao := daos.NewAdvertDao(db)
	advertService := services.NewAdvertService(advertDao, utils)
	advertController := controllers.NewAdvertController(r, "", advertService)
	advertController.InitAdvertController()

	/* 初始化 Video*/
	videoDao := daos.NewVideoDao(db)
	videoService := services.NewVideoService(videoDao, utils)
	videoController := controllers.NewVideoController(r, "", videoService)
	videoController.InitVideoController()

	/*初始化channelBase*/
	channelBaseDao := daos.NewChannelBaseDao(db)
	channelBaseService := services.NewChannelBaseService(channelBaseDao, utils)
	channelBaseController := controllers.NewChannelBaseController(r, "", channelBaseService)
	channelBaseController.InitCannelBaseController()

	/*初始化comment*/
	commentDao := daos.NewCommentDao(db)
	commentService := services.NewCommentService(commentDao, utils)
	commentController := controllers.NewCommentController(r, "", commentService)
	commentController.InitCommentController()

	/*批量通知消息发送*/
	messageDao := daos.NewMessageDao(db)
	messageService := services.NewMessageService(messageDao, utils)
	messageController := controllers.NewMessageController(r, "", messageService)
	messageController.InitMessageController()

	/*弹幕*/
	barrageDao := daos.NewBarrageDao(db)
	barrageService := services.NewBarrageService(barrageDao, utils)
	barrageController := controllers.NewBarrageController(r, "", barrageService)
	barrageController.InitBarrageController()

	/*阿里云*/
	aliyunClient, _ := services.NewAliyunClient(appConf.AccessKeyId, appConf.AccessKeySecret)
	aliyunService := services.NewAliyunService(aliyunClient)
	aliyunController := controllers.NewAliyunController(r, "", aliyunService)
	aliyunController.InitAliyunController()

	/*运行*/
	r.Run(appConf.RunConf.Address)

}
