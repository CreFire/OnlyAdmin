package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/model"
)

type RoomRepo struct{}

func (r *RoomRepo) Coll() (*mongo.Collection, error) {
	return GetCollection("room")
}

// 查询列表
func (r *RoomRepo) List(ctx context.Context, tenantID string) ([]*model.Room, error) {
	coll, err := r.Coll()
	if err != nil {
		return nil, err
	}
	filter := bson.M{}
	if tenantID != "" {
		filter["tenant_id"] = tenantID
	}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []*model.Room
	for cursor.Next(ctx) {
		var room model.Room
		if err := cursor.Decode(&room); err == nil {
			rooms = append(rooms, &room)
		}
	}
	return rooms, nil
}

// 新增
func (r *RoomRepo) Create(ctx context.Context, room *model.Room) error {
	coll, err := r.Coll()
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(ctx, room)
	return err
}

// 更新
func (r *RoomRepo) Update(ctx context.Context, id string, data bson.M) error {
	coll, err := r.Coll()
	if err != nil {
		return err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	return err
}

// 删除
func (r *RoomRepo) Delete(ctx context.Context, id string) error {
	coll, err := r.Coll()
	if err != nil {
		return err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = coll.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
