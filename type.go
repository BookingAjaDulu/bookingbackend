package bookingbackend

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username         string             `bson:"username," json:"username,"`
	Password         string             `bson:"password," json:"password,"`
	Confirm_Password string             `bson:"confirm_password," json:"confirm_password,"`
}

type Credential struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Lapangan struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Lapangan string             `bson:"nama_lapangan,omitempty" json:"nama_lapangan,omitempty"`
	Harga         string             `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi     string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Gambar        string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
}
