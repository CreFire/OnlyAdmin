package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/model"
)

type RuleRepo struct{}

func (r *RuleRepo) Coll() (*mongo.Collection, error) {
	return GetCollection("rule")
}

// 查询列表
func (r *RuleRepo) List(ctx context.Context, tenantID string) ([]*model.Rule, error) {
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

	var rules []*model.Rule
	for cursor.Next(ctx) {
		var rule model.Rule
		if err := cursor.Decode(&rule); err == nil {
			rules = append(rules, &rule)
		}
	}
	return rules, nil
}

// 新增
func (r *RuleRepo) Create(ctx context.Context, rule *model.Rule) error {
	coll, err := r.Coll()
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(ctx, rule)
	return err
}

// 更新
func (r *RuleRepo) Update(ctx context.Context, id string, data bson.M) error {
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
func (r *RuleRepo) Delete(ctx context.Context, id string) error {
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
