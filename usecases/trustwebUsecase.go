package usecases

import (
	models "gonote/models"
	repository "gonote/repository"
	"time"

	"gorm.io/gorm"
)

// 使用 ClientID 進行獲取
func SearchTrust(db *gorm.DB, clientID string) []models.Searchtrustweb {
	trustweb := repository.GetTrustWeb(db, clientID)
	return trustweb
}

// 統計被使用的信任網站數量
func SearchTrustCount(db *gorm.DB) []models.UrlCount {
	trustweb := repository.GetUrlCounts(db)
	return trustweb
}

func InsertSessionToDB(db *gorm.DB, session models.Session) error {
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
			return err
		}
	}
	return nil
}
