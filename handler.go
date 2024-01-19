package bookingbackend

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	user       User
	lapangan   Lapangan
	credential Credential
	response   Response
)

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func UserLogin(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	credential.Status = 400

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		credential.Message = "error parsing application/json: " + err.Error()
	}

	users, _, err := Login(mconn, collectionname, user)
	if err != nil {
		credential.Message = err.Error()
		return GCFReturnStruct(credential)
	}

	credential.Status = 200
	credential.Message = "Selamat Datang " + users.Username

	return GCFReturnStruct(credential)
}

// lapangan
func GetAllDataLapangan(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	data, err := GetAllLapangan(mconn, collectionname)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get All Lapangan Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func GetDataLapanganByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	data, err := GetLapanganByID(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get Lapangan By ID Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func InsertLapangan(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	err := json.NewDecoder(r.Body).Decode(&lapangan)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	data, err := InsertLapang(mconn, collectionname, lapangan)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Insert lapangan Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func UpdateDataLapangan(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	err = json.NewDecoder(r.Body).Decode(&lapangan)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	data, err := UpdateLapangan(mconn, collectionname, ID, lapangan)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Update lapangan Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func DeleteDataLapangan(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	_, err = DeleteLapangan(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Delete Obat Success " + lapangan.Nama_Lapangan

	return GCFReturnStruct(response)
}
