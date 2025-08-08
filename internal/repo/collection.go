package repo

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"server/internal/model"
	"time"
)

// 自动创建集合及索引（可选）
func EnsureCollections() error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	dbName := "tenantdb" // 可用 viper 读取
	db := client.Database(dbName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. 创建 users 集合（可选，不写也没事，插入时会自动创建）
	if err := db.CreateCollection(ctx, "users"); err != nil {
		// 已存在会报错，可以忽略
		logrus.Warn("users 集合可能已存在: ", err)
	} else {
		logrus.Info("users 集合已创建")
	}

	// 2. 创建索引（例如用户名和租户ID唯一索引）
	userColl := db.Collection("users")
	idx := mongo.IndexModel{
		Keys:    map[string]interface{}{"username": 1, "tenant_id": 1},
		Options: options.Index().SetUnique(true),
	}
	if _, err := userColl.Indexes().CreateOne(ctx, idx); err != nil {
		logrus.Warn("users 唯一索引创建失败: ", err)
	} else {
		logrus.Info("users 唯一索引创建成功")
	}

	// 3. 可继续为其他集合创建和建索引

	return nil
}

func EnsureAdminUser() error {
	admin, err := FindAdmin()
	if err != nil {
		return err
	}
	if admin != nil {
		logrus.Info("管理员已存在，无需初始化。")
		return nil
	}
	username := viper.GetString("admin.username")
	password := viper.GetString("admin.password")
	if username == "" || password == "" {
		return fmt.Errorf("管理员用户名和密码配置缺失")
	}
	// 加密存储
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminUser := &model.User{
		Username: username,
		Password: string(hashed),
		Role:     "admin",
		Active:   true,
	}
	err = InsertUser(adminUser)
	if err != nil {
		return err
	}
	logrus.Warnf("已自动创建超级管理员账号：%s，初始密码请尽快修改！", username)
	return nil
}
