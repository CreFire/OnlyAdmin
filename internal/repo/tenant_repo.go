package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/model"
)

type TenantRepo struct{}

func (r *TenantRepo) Coll() *mongo.Collection {
	collection, err := GetCollection("tenant")
	if err != nil {
		return nil
	}
	return collection
}

// 列表
func (r *TenantRepo) List(ctx context.Context, tenantID string) ([]*model.Tenant, error) {
	var results []*model.Tenant
	cursor, err := r.Coll().Find(ctx, bson.M{
		// 可加租户隔离条件
		// "tenant_id": tenantID,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var t model.Tenant
		if err := cursor.Decode(&t); err == nil {
			results = append(results, &t)
		}
	}
	return results, nil
}

// 新增
func (r *TenantRepo) Create(ctx context.Context, t *model.Tenant) error {
	_, err := r.Coll().InsertOne(ctx, t)
	return err
}

// 更新
func (r *TenantRepo) Update(ctx context.Context, id string, data bson.M) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Coll().UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})
	return err
}

// 删除
func (r *TenantRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Coll().DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
