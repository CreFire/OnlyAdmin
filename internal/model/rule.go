package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rule struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TenantID      string             `bson:"tenant_id" json:"tenant_id"`
	RoomID        string             `bson:"room_id" json:"room_id"`
	Type          string             `bson:"type" json:"type"`   // "rent" | "water" | "electricity" | "gas"
	Cycle         string             `bson:"cycle" json:"cycle"` // "monthly" | "quarterly" | "custom"
	Day           int                `bson:"day" json:"day"`     // 每月/周期的第几天
	Amount        float64            `bson:"amount" json:"amount"`
	PaymentMethod string             `bson:"payment_method" json:"payment_method"` // "wechat" | "alipay" | "cash"
	RemindBefore  int                `bson:"remind_before" json:"remind_before"`   // 提前几天提醒
	Active        bool               `bson:"active" json:"active"`
	Note          string             `bson:"note,omitempty" json:"note,omitempty"`
}
