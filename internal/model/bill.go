package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bill struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TenantID      string             `bson:"tenant_id" json:"tenant_id"`
	RoomID        string             `bson:"room_id" json:"room_id"`
	Type          string             `bson:"type" json:"type"` // "rent" | "water" | "electricity" | "gas"
	Amount        float64            `bson:"amount" json:"amount"`
	Period        string             `bson:"period" json:"period"`                             // 账单期，如 "2024-06"
	DueDate       string             `bson:"due_date" json:"due_date"`                         // 到期日 "2024-06-05"
	Status        string             `bson:"status" json:"status"`                             // "unpaid" | "paid" | "overdue"
	PayTime       string             `bson:"pay_time,omitempty" json:"pay_time,omitempty"`     // 实际支付时间
	PaymentMethod string             `bson:"payment_method" json:"payment_method"`             // "wechat" | "alipay" | "cash"
	Voucher       string             `bson:"voucher,omitempty" json:"voucher,omitempty"`       // 凭证URL
	RemindLog     []string           `bson:"remind_log,omitempty" json:"remind_log,omitempty"` // 提醒记录
	Note          string             `bson:"note,omitempty" json:"note,omitempty"`
}
