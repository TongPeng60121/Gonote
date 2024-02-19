package connectsql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB(Dbname string) *gorm.DB {
	//設置MySQL連線
	username := "root"     //帳號
	password := "abc12345" //密碼
	host := "127.0.0.1"    //IP
	port := 3377           //port

	// 設置 MariaDB 連接資訊
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	// 連接到 MariaDB
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
