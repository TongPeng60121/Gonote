package http

import (
	"gonote/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type httpDelivery struct {
	trustWebSiteUsecase usecase.TrustWebSiteUsecase
	db                  *gorm.DB
}

func NewHTTPDelivery(trustWebSiteUsecase usecase.TrustWebSiteUsecase, db *gorm.DB) *httpDelivery {
	return &httpDelivery{
		trustWebSiteUsecase: trustWebSiteUsecase,
		db:                  db,
	}
}

func SetupRoutes(r *gin.Engine, db *gorm.DB, trustWebSiteUsecase usecase.TrustWebSiteUsecase) {
	// 创建 httpDelivery 实例
	delivery := NewHTTPDelivery(trustWebSiteUsecase, db)

	// 使用 ClientID 進行獲取 - GET
	r.GET("/api/GetSession/:clientID", delivery.GetSession)

	// 统计被使用的信任网站数量 - GET
	//r.GET("/api/GetSessionCount", delivery.GetSessionCount)

	// 新增
	//r.POST("/api/session", delivery.InsertSession)
}

// GetSession 处理获取 TrustWeb 数据的请求
func (h *httpDelivery) GetSession(c *gin.Context) {
	clientID := c.Param("clientID")

	// 调用 usecase 的 SearchTrustWeb 方法
	trustweb, err := h.trustWebSiteUsecase.SearchTrustWeb(clientID)

	// 处理潜在的错误
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch TrustWeb data"})
		return
	}

	// 返回 TrustWeb 数据
	c.JSON(200, trustweb)
}

// GetSessionCount 处理获取信任网站数量的请求
/*func (h *httpDelivery) GetSessionCount(c *gin.Context) {
	// 调用 usecase 的 SearchTrustCount 方法
	trustcount, err := h.trustWebSiteUsecase.SearchTrustCount(db)

	// 处理潜在的错误
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch TrustWeb count"})
		return
	}

	// 返回 TrustWeb 数量
	c.JSON(200, trustcount)
}

// InsertSession 处理新增 Session 数据的请求
func (h *httpDelivery) InsertSession(c *gin.Context) {
	// 调用 usecase 的 InsertSessionToDB 方法
	err := h.trustWebSiteUsecase.InsertSessionToDB(db, c)

	// 处理潜在的错误
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert data into database"})
		return
	}

	// 返回成功的消息
	c.JSON(200, gin.H{"message": "Data inserted successfully"})
}*/
