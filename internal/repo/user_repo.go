package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"server/internal/model"
	"time"
)

// 查询用户
func FindUserByEmail(email string) (*model.User, error) {
	coll, err := GetCollection("user")
	if err != nil {
		return nil, err
	}
	var user model.User
	err = coll.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, nil // 用户不存在
	}
	return &user, nil
}

// 修改密码
func UpdateUserPasswordByEmail(email, password string) error {
	coll, err := GetCollection("user")
	if err != nil {
		return err
	}
	// password请做加密存储
	_, err = coll.UpdateOne(context.TODO(),
		bson.M{"email": email},
		bson.M{"$set": bson.M{"password": password}},
	)
	return err
}

func FindUserByNameAndTenant(username, tenantID string) (*model.User, error) {
	coll, err := GetCollection("users")
	if err != nil {
		return nil, err
	}
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = coll.FindOne(ctx, bson.M{
		"username":  username,
		"tenant_id": tenantID,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InsertUser(user *model.User) error {
	coll, err := GetCollection("users")
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = coll.InsertOne(ctx, user)
	return err
}

// 查找管理员
func FindAdmin() (*model.User, error) {
	coll, err := GetCollection("user")
	if err != nil {
		return nil, err
	}
	var admin model.User
	err = coll.FindOne(context.TODO(), bson.M{"role": "admin"}).Decode(&admin)
	if err != nil {
		return nil, nil // 不存在
	}
	return &admin, nil
}
