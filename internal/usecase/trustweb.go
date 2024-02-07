package usecase

import (
	"gonote/internal/repository"
)

type trustWebSiteUsecase struct {
	trustWebSiteRepo repository.TrustWebRepository
}

// NewTrustWebSiteUsecase 创建 TrustWebSiteUsecase 实例
func NewTrustWebSiteUsecase(repo repository.TrustWebRepository) TrustWebSiteUsecase {
	return &trustWebSiteUsecase{
		trustWebSiteRepo: repo,
	}
}

// 使用 ClientID 进行获取
func (t *trustWebSiteUsecase) SearchTrustWeb(clientID string) ([]repository.Trustweb, error) {
	trustweb, err := t.trustWebSiteRepo.GetTrustWebSites(clientID)
	if err != nil {
		// 处理错误
		return nil, err
	}
	return trustweb, nil
}

// 統計被使用的信任網站數量
/*func (t *trustWebSiteUsecase) SearchTrustCount(db *gorm.DB) {
	TrustCount, err := repository.GetUrlCounts()
	if err != nil {
		// 处理错误
		return nil, err
	}
	return TrustCount, nil
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
*/
