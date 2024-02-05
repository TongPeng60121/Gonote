package main

import (
	"gonote/connectsql"
	models "gonote/models"
	usecases "gonote/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	Dbname := "city" //資料庫
	db := connectsql.InitDB(Dbname)
	defer func() {
		// 通過 db.DB() 獲取底層的 *sql.DB 對象，然後調用其 Close 方法來關閉連接
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// 建立對應的資料表
	db.AutoMigrate(&models.City{})

	// 創建 GIN 引擎
	r := gin.Default()

	// 設定路由
	setupRoutes(r, db)

	// 啟動 GIN 服務
	r.Run(":8080")
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {

	// 使用 ClientID 進行獲取 - GET
	r.GET("/api/GetSession/:clientID", func(c *gin.Context) {
		usecases.SearchTrust(db, c)
	})

	// 統計被使用的信任網站數量 - GET
	r.GET("/api/GetSessionCount", func(c *gin.Context) {
		usecases.SearchTrustCount(db, c)
	})

	// 新增
	r.GET("/api/session", func(c *gin.Context) {
		usecases.InsertSessionToDB(db, c)
	})

	// 創建新城市 - POST
	r.POST("/api/cities", func(c *gin.Context) {
		usecases.CreateCities(db, c)
	})

	// 取得所有城市 - GET
	r.GET("/api/cities/all", func(c *gin.Context) {
		usecases.GetAllCities(db, c)
	})

	// 取得多個城市 - GET
	r.GET("/api/cities", func(c *gin.Context) {
		usecases.GetCities(db, c)
	})

	// 更新城市信息 - PUT
	r.PUT("/api/cities", func(c *gin.Context) {
		usecases.UpdateCities(db, c)
	})

	// 刪除城市 - DELETE
	r.DELETE("/api/cities", func(c *gin.Context) {
		usecases.DeleteCities(db, c)
	})
}
