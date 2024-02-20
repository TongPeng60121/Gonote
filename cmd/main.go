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
	defer db.Close()

	// 創建 TrustWebRepository
	trustWebRepo := repository.NewTrustWebRepository(db)

	// 創建 TrustWebSiteUsecase
	trustWebUsecase := usecase.NewTrustWebSiteUsecase(trustWebRepo)

	// 創建 GIN
	r := gin.Default()

	// 設定路由，提供 trustWebSiteUsecase
	http.SetupRoutes(r, db, trustWebUsecase)

	// 啟動 GIN 服務
	r.Run(":8080")
}
