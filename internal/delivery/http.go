package http

import (
	"gonote/internal/errors"
	"gonote/internal/repository"
	"gonote/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	// 創建 httpDelivery
	delivery := NewHTTPDelivery(trustWebSiteUsecase, db)

	// 使用 ClientID 進行抓取 - GET
	r.GET("/api/get-clientid/:clientID", delivery.GetSession)

	// 統計被使用的信任網站數量 - GET
	r.GET("/api/get-session-count/:limit", delivery.GetSessionCount)

	// 新增
	r.POST("/api/insert-session", delivery.InsertSession)
}

// GetSession 抓取 TrustWeb 資料的請求
func (h *httpDelivery) GetSession(c *gin.Context) {
	clientIDStr := c.Param("clientID")

	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		// 格式錯誤
		c.JSON(400, gin.H{"error": "Invalid clientID"})
		return
	}

	// 使用 usecase 的 SearchTrustWeb 方法
	trustweb, err := h.trustWebSiteUsecase.SearchTrustWeb(c.Request.Context(), clientID)

	// 處理錯誤
	if err != nil {
		IsErrorEx(c, err)
	} else {
		// 返回 TrustWeb 資料
		c.JSON(200, trustweb)
	}
}

// GetSessionCount 抓取信任網站的請求
func (h *httpDelivery) GetSessionCount(c *gin.Context) {
	limit := c.Param("limit")

	// 使用 usecase 的 SearchTrustCount 方法
	trustcount, err := h.trustWebSiteUsecase.SearchTrustCount(c.Request.Context(), limit)

	// 處理錯誤
	if err != nil {
		IsErrorEx(c, err)
	} else {
		// 返回 TrustWeb 數量
		c.JSON(200, trustcount)
	}
}

// InsertSession 新增 Session 資料的請求
func (h *httpDelivery) InsertSession(c *gin.Context) {
	var session []repository.TrustJson
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 使用 usecase 的 InsertSessionToDB 方法
	err := h.trustWebSiteUsecase.InsertSessionToDB(c.Request.Context(), session)

	// 處理錯誤
	if err != nil {
		IsErrorEx(c, err)
	} else {
		c.JSON(200, gin.H{"message": "Data inserted successfully"})
	}
}

func IsErrorEx(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(errors.Error)
	if !ok {
		e = errors.Error{
			Type:    errors.Unknown,
			Message: err.Error(),
		}
	}
	c.JSON(int(e.Type), gin.H{"Error": e.Message})
	return true
}
