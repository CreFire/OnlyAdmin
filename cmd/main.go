package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"server/api"
	"server/internal/repo"
)

func initConfig() {

	viper.SetConfigName("config")   // 配置文件名，不带扩展名
	viper.SetConfigType("yaml")     // 配置文件类型
	viper.AddConfigPath("./config") // 查找路径

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %v", err)
	}

	logrus.Infof("配置文件加载成功: %s", viper.ConfigFileUsed())
}

func initLog() {
	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

func main() {
	initConfig()
	initLog()
	// 初始化MongoDB集合和索引
	if err := repo.EnsureCollections(); err != nil {
		panic("MongoDB集合初始化失败: " + err.Error())
	}
	if err := repo.EnsureAdminUser(); err != nil {
		panic("初始化管理员账号失败: " + err.Error())
	}
	mux := gin.Default()
	api.RegisterRoutes(mux)
	addr := viper.GetString("addr")
	logrus.Infof("服务启动：%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		logrus.Fatalf("服务异常退出：%v", err)
	}
}
