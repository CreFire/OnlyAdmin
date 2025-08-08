package repo

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

// GetMongoClient 获取 MongoDB 客户端
func GetMongoClient() (*mongo.Client, error) {
	var err error
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		mongoURI := viper.GetString("db_url")
		client, e := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		if e != nil {
			err = e
			return
		}
		// 测试连接
		if e = client.Ping(ctx, nil); e != nil {
			err = e
			return
		}
		mongoClient = client
		logrus.Infof("MongoDB连接成功: %s", mongoURI)
	})
	return mongoClient, err
}

// 获取 Database
func GetDatabase() (*mongo.Database, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	dbName := viper.GetString("db_name")
	if dbName == "" {
		dbName = "tenantdb"
	}
	return client.Database(dbName), nil
}

// 获取 Collection
func GetCollection(name string) (*mongo.Collection, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	return db.Collection(name), nil
}
