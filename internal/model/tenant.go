// model/tenant.go
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tenant struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Phone    string             `bson:"phone" json:"phone"`
	WxOpenID string             `bson:"wx_openid" json:"wx_openid"`
	RoomID   primitive.ObjectID `bson:"room_id" json:"room_id"`
	Active   bool               `bson:"active" json:"active"`
}
