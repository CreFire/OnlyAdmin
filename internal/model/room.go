package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Address  string             `bson:"address" json:"address"`
	Landlord string             `bson:"landlord" json:"landlord"`
	TenantID string             `bson:"tenant_id" json:"tenant_id"` // 租户归属
	Note     string             `bson:"note,omitempty" json:"note,omitempty"`
	Active   bool               `bson:"active" json:"active"`
}
