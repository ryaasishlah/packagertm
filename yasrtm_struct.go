package packagertm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type yang_buat struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_yangbuat       string             `bson:"nama_yangbuat,omitempty" json:"nama_yangbuat,omitempty"`
	Email_yangbuat      string             `bson:"email_yangbuat,omitempty" json:"email_yangbuat,omitempty"`
	Katasandi_yangbuat  string             `bson:"katasandi_yangbuat,omitempty" json:"katasandi_yangbuat,omitempty"`
	Preferensi_yangbuat string             `bson:"preferensi_yangbuat,omitempty" json:"preferensi_yangbuat,omitempty"`
}
