package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/model"
)

type BillRepo struct{}

func (r *BillRepo) Coll() (*mongo.Collection, error) {
	return GetCollection("bill")
}

// 查询列表
func (r *BillRepo) List(ctx context.Context, tenantID string) ([]*model.Bill, error) {
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

	var bills []*model.Bill
	for cursor.Next(ctx) {
		var bill model.Bill
		if err := cursor.Decode(&bill); err == nil {
			bills = append(bills, &bill)
		}
	}
	return bills, nil
}

// 新增
func (r *BillRepo) Create(ctx context.Context, bill *model.Bill) error {
	coll, err := r.Coll()
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(ctx, bill)
	return err
}

// 更新
func (r *BillRepo) Update(ctx context.Context, id string, data bson.M) error {
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
func (r *BillRepo) Delete(ctx context.Context, id string) error {
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
