package usecases

import (
	models "gonote/models"
	repository "gonote/repository"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 使用 ClientID 進行獲取
func SearchTrust(db *gorm.DB, c *gin.Context) {
	clientID := c.Param("clientID")
	trustweb := repository.GetTrustWeb(db, clientID)
	if len(trustweb) == 0 {
		c.JSON(404, gin.H{"error": "TrustWeb not found"})
		return
	}
	c.JSON(200, trustweb)
}

// 統計被使用的信任網站數量
func SearchTrustCount(db *gorm.DB, c *gin.Context) {
	TrustCount := repository.GetUrlCounts(db)
	if len(TrustCount) == 0 {
		c.JSON(404, gin.H{"error": "TrustWeb not found"})
		return
	}
	c.JSON(200, TrustCount)
}

func InsertSessionToDB(db *gorm.DB, c *gin.Context) {
	var session models.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, trustWeb := range session.TrustWeb {
		// 檢查是否已存在相同的 TrustWeb 記錄
		existingTrustWeb := repository.GetTrustWebFromDB(db, session.SessionID, session.ClientID, trustWeb.Url)
		if existingTrustWeb != nil {
			// 如果已存在相同數據，則直接傳回，不執行插入操作
			continue
		}

		// 插入 TrustWeb 數據
		newTrustWeb := models.Trustweb{
			SessionID: session.SessionID,
			ClientID:  session.ClientID,
			Url:       trustWeb.Url,
			Cdate:     time.Now(),
		}

		if err := repository.CreateTrustWeb(db, &newTrustWeb); err != nil {
			// 处理插入数据库时的错误
			c.JSON(500, gin.H{"error": "Failed to insert data into database"})
			return
		}
	}

	c.JSON(201, session)
}
