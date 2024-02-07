package connectsql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(Dbname string) *gorm.DB {
	//設置MySQL連線
	username := "root"  //帳號
	password := ""      //密碼
	host := "127.0.0.1" //IP
	port := 3306        //port

	// 設置 MariaDB 連接資訊
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	// 連接到 MariaDB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
