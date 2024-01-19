package bookingbackend

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoConnect(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func InsertOneDoc(db *mongo.Database, col string, docs interface{}) (insertedID primitive.ObjectID, err error) {
	cols := db.Collection(col)
	result, err := cols.InsertOne(context.Background(), docs)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, err
}

func Register(db *mongo.Database, col string, userdata User) error {
	cols := db.Collection(col)

	hash, _ := HashPassword(userdata.Password)
	user := bson.D{
		{Key: "username", Value: userdata.Username},
		{Key: "password", Value: hash},
	}

	result, err := cols.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	userdata.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func Login(db *mongo.Database, col string, userdata User) (user User, status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"username": userdata.Username}

	err = cols.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		err = fmt.Errorf("username tidak ditemukan")
		return user, false, err
	}

	if !CheckPasswordHash(userdata.Password, user.Password) {
		err = fmt.Errorf("password salah")
		return user, false, err
	}

	return user, true, nil
}

// obat
func GetAllLapangan(db *mongo.Database, col string) (lapangan []Lapangan, err error) {
	cols := db.Collection(col)

	cursor, err := cols.Find(context.Background(), bson.M{})
	if err != nil {
		return lapangan, err
	}

	err = cursor.All(context.Background(), &lapangan)
	if err != nil {
		return lapangan, err
	}

	return lapangan, nil
}

func GetLapanganByID(db *mongo.Database, col string, _id primitive.ObjectID) (lapangan Lapangan, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	err = cols.FindOne(context.Background(), filter).Decode(&lapangan)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("data tidak di temukan dengan ID: ", _id)
		} else {
			fmt.Println("error retrieving data for ID", _id, ":", err.Error())
		}
	}

	return lapangan, nil
}

func InsertLapang(db *mongo.Database, col string, lapangan Lapangan) (docs Lapangan, err error) {
	objectID := primitive.NewObjectID()

	datalapangan := bson.D{
		{Key: "_id", Value: objectID},
		{Key: "nama_lapangan", Value: lapangan.Nama_Lapangan},
		{Key: "harga", Value: lapangan.Harga},
		{Key: "deskripsi", Value: lapangan.Deskripsi},
		{Key: "gambar", Value: lapangan.Gambar},
	}

	InsertedID, err := InsertOneDoc(db, col, datalapangan)
	if err != nil {
		fmt.Printf("InsertObat: %v\n", err)
		return docs, err
	}

	docs.ID = InsertedID

	return docs, nil
}

func UpdateLapangan(db *mongo.Database, col string, _id primitive.ObjectID, lapangan Lapangan) (docs Lapangan, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "nama_lapangan", Value: lapangan.Nama_Lapangan},
			{Key: "harga", Value: lapangan.Harga},
			{Key: "deskripsi", Value: lapangan.Deskripsi},
			{Key: "gambar", Value: lapangan.Gambar},
		}},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return docs, err
	}

	if result.MatchedCount == 0 {
		return docs, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return docs, nil
}

func DeleteLapangan(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return true, nil
}
