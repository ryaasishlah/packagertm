package packagertm

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "PengaturanAplikasi",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertDataPengaturan(t *testing.T) {
	Nama_yangbuat := "Ryaas Ishlah"
	Email_yangbuat := "ryaasganteng@gmail.com"
	Katasandi_yangbuat := "ryaasganteng"
	Preferensi_yangbuat := "Pengaturan Aplikasi"
	hasil := InsertDataPengaturan(MongoConn, Nama_yangbuat, Email_yangbuat, Katasandi_yangbuat, Preferensi_yangbuat)
	fmt.Println(hasil)
}

func TestGetDataPengaturan(t *testing.T) {
	Nama_yangbuat := "Ryaas Ishlah"
	hasil := GetDataPengaturan(Nama_yangbuat, MongoConn, "yang_buat")
	fmt.Println(hasil)
}

func TestGetDataListFromPreferensi(t *testing.T) {
	Preferensi_yangbuat := "Pengaturan Aplikasi"
	hasil := GetDataListFromPreferensi(Preferensi_yangbuat, MongoConn, "yang_buat")
	fmt.Println(hasil)
}

func TestDeleteDataPengaturan(t *testing.T) {
	Email_yangbuat := "ryaasganteng@gmail.com"
	hasil := DeleteDataPengaturan(Email_yangbuat, MongoConn, "yang_buat")
	fmt.Println(hasil)
}
