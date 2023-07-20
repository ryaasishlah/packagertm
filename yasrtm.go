package packagertm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDataPengaturan(db *mongo.Database, Nama_yangbuat, Email_yangbuat string, Katasandi_yangbuat string, Preferensi_yangbuat string) (InsertedID interface{}) {
	var yangbuat yang_buat
	yangbuat.Nama_yangbuat = Nama_yangbuat
	yangbuat.Email_yangbuat = Email_yangbuat
	yangbuat.Katasandi_yangbuat = Katasandi_yangbuat
	yangbuat.Preferensi_yangbuat = Preferensi_yangbuat
	return InsertOneDoc(db, "yang_buat", yangbuat)
}

func GetDataListFromPreferensi(Preferensi_yangbuat string, db *mongo.Database, col string) (data yang_buat) {
	pfr := db.Collection(col)
	filter := bson.M{"preferensiyangbuat": Preferensi_yangbuat}
	err := pfr.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataListbypfr: %v\n", err)
	}
	return data
}

func GetDataListFromKatasandi(Katasandi_yangbuat string, db *mongo.Database, col string) (data yang_buat) {
	kts := db.Collection(col)
	filter := bson.M{"katasandiyangbuat": Katasandi_yangbuat}
	err := kts.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataListbykts: %v\n", err)
	}
	return data
}

func GetDataPengaturan(Nama_yangbuat string, db *mongo.Database, col string) (data yang_buat) {
	ygb := db.Collection(col)
	filter := bson.M{"namanyayangbuat": Nama_yangbuat}
	err := ygb.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataListbyygb: %v\n", err)
	}
	return data
}

func DeleteDataPengaturan(Email_yangbuat string, db *mongo.Database, col string) (data yang_buat) {
	emlys := db.Collection(col)
	filter := bson.M{"emailygbyangbuat": Email_yangbuat}
	err, _ := emlys.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataEmailygb : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
