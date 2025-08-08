package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"` // 建议存加密
	Role     string             `bson:"role" json:"role"`  // "admin" "user"
	Email    string             `bson:"email" json:"email,omitempty"`
	TenantID int32              `bson:"tenantID" json:"tenant_id,omitempty"`
	Active   bool               `bson:"active" json:"active"`
}
