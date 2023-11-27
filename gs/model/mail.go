package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mail struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	MailId     uint32             `bson:"mail_id"`
	Title      string             `bson:"title"`
	Content    string             `bson:"content"`
	Sender     string             `bson:"sender"`
	SendTime   uint32             `bson:"send_time"`
	ExpireTime uint32             `bson:"expire_time"`
	IsRead     bool               `bson:"is_read"`
	IsStar     bool               `bson:"is_star"`
}
