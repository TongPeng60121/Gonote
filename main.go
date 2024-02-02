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
		clientID := c.Param("clientID")

		TrustWeb := usecases.SearchTrust(db, clientID)
		if len(TrustWeb) == 0 {
			c.JSON(404, gin.H{"error": "TrustWeb not found"})
			return
		}
		c.JSON(200, TrustWeb)
	})

	// 統計被使用的信任網站數量 - GET
	r.GET("/api/GetSessionCount", func(c *gin.Context) {
		TrustCount := usecases.SearchTrustCount(db)
		if len(TrustCount) == 0 {
			c.JSON(404, gin.H{"error": "TrustWeb not found"})
			return
		}
		c.JSON(200, TrustCount)
	})

	// 新增
	r.GET("/api/session", func(c *gin.Context) {
		var session models.Session

		if err := c.ShouldBindJSON(&session); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 將 Session 資料插入資料庫
		if err := usecases.InsertSessionToDB(db, session); err != nil {
			// 处理插入数据库时的错误
			c.JSON(500, gin.H{"error": "Failed to insert data into database"})
			return
		}

		c.JSON(201, session)
	})

	// 創建新城市 - POST
	r.POST("/api/cities", func(c *gin.Context) {
		var newCities []models.City
		if err := c.ShouldBindJSON(&newCities); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		usecases.CreateCities(db, newCities)
		c.JSON(201, newCities)
	})

	// 取得所有城市 - GET
	r.GET("/api/cities/all", func(c *gin.Context) {
		allCities := usecases.GetAllCities(db)
		c.JSON(200, allCities)
	})

	// 取得多個城市 - GET
	r.GET("/api/cities", func(c *gin.Context) {
		var query models.Query

		if err := c.ShouldBindJSON(&query); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		cities, err := usecases.GetCities(db, query.IDs, query.Names)
		if err != nil {
			c.JSON(404, gin.H{"error": "Cities not found"})
			return
		}
		c.JSON(200, cities)
	})

	// 更新城市信息 - PUT
	r.PUT("/api/cities", func(c *gin.Context) {
		var updatedCities []models.City
		if err := c.ShouldBindJSON(&updatedCities); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		usecases.UpdateCities(db, updatedCities)
		c.JSON(200, gin.H{"message": "Cities updated successfully"})
	})

	// 刪除城市 - DELETE
	r.DELETE("/api/cities", func(c *gin.Context) {
		var ids []uint
		if err := c.ShouldBindJSON(&ids); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		usecases.DeleteCities(db, ids)
		c.JSON(200, gin.H{"message": "Cities deleted successfully"})
	})
}
