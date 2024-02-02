package repository

import (
	models "gonote/models"
	"log"

	"gorm.io/gorm"
)

// 查詢
func GetTrustWebFromDB(db *gorm.DB, sessionID, clientID int, url string) *models.Trustweb {
	var existingTrustWeb models.Trustweb
	result := db.Table("trustweb").Where("SessionID = ? AND ClientID = ? AND Url = ?", sessionID, clientID, url).First(&existingTrustWeb)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果未找到相同數據，則傳回 nil
		return nil
	} else if result.Error != nil {
		// 處理其他查詢錯誤
		log.Println("Error querying database:", result.Error)
		return nil
	}
	// 傳回查詢到的 TrustWeb 數據
	return &existingTrustWeb
}

// 取得Url
func GetTrustWeb(db *gorm.DB, ClientID string) []models.Searchtrustweb {
	var allTrustweb []models.Searchtrustweb
	db.Table("trustweb").Select("url, cdate").Where("ClientID = ?", ClientID).Find(&allTrustweb)

	// 转换时间为字符串形式
	for i, trustweb := range allTrustweb {
		allTrustweb[i].CdateString = trustweb.Cdate.Format("2006-01-02 15:04:05")
	}

	return allTrustweb
}

// 取得每個Url的數量
func GetUrlCounts(db *gorm.DB) []models.UrlCount {
	var urlCounts []models.UrlCount
	db.Table("trustweb").Select("url, COUNT(*) as count").Group("url").Find(&urlCounts)

	return urlCounts
}

// 新增
func CreateTrustWeb(db *gorm.DB, trustWeb *models.Trustweb) error {
	result := db.Table("trustweb").Create(trustWeb)
	if result.Error != nil {
		// 處理新增錯誤
		log.Println("Error creating TrustWeb:", result.Error)
		return result.Error
	}
	return nil
}
