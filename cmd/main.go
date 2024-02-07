package main

import (
	"gonote/internal/connectsql"
	http "gonote/internal/delivery"
	"gonote/internal/repository"
	"gonote/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	Dbname := "trustwebsite" //資料庫
	db := connectsql.InitDB(Dbname)
	defer func() {
		// 通過 db.DB() 獲取底層的 *sql.DB 對象，然後調用其 Close 方法來關閉連接
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// 创建 TrustWebRepository 实例
	trustWebRepo := repository.NewTrustWebRepository(db)

	// 创建 TrustWebSiteUsecase 实例
	trustWebUsecase := usecase.NewTrustWebSiteUsecase(trustWebRepo)

	// 创建 GIN 引擎
	r := gin.Default()

	// 設定路由，提供 trustWebSiteUsecase 实例
	http.SetupRoutes(r, db, trustWebUsecase)

	// 啟動 GIN 服務
	r.Run(":8080")
}
