package repository

import (
	"log"

	"gorm.io/gorm"
)

type Trustweb struct {
	SessionID string
	Url       string
	Cdate     string
}

type Session struct {
	SessionID string
	ClientID  string
}

type TrustwebSession struct {
	SessionID uint32
	ClientID  uint32
	Url       string
}

type UrlCount struct {
	TrustwebSession
	Count uint32
}

type trustWebRepository struct {
	db *gorm.DB
}

func NewTrustWebRepository(db *gorm.DB) TrustWebRepository {
	return &trustWebRepository{
		db: db,
	}
}

func (t *trustWebRepository) GetTrustWebSites(clientID string) ([]Trustweb, error) {
	var existingTrustWebs []Trustweb
	result := t.db.Raw("SELECT trustweb.SessionID, Url, trusturl.Cdate FROM trustweb INNER JOIN trusturl ON trustweb.SessionID = trusturl.SessionID WHERE ClientID = ?", clientID).Scan(&existingTrustWebs)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果未找到相同數據，則傳回 nil
		return nil, nil
	} else if result.Error != nil {
		// 處理其他查詢錯誤
		log.Println("Error querying database:", result.Error)
		return nil, result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return existingTrustWebs, nil
}

/*func (t *trustWebRepository) GetUrlCounts(db *gorm.DB) ([]UrlCount, error) {
	var urlCounts []UrlCount
	result := db.Raw("SELECT Url, COUNT(*) as Count FROM trustweb INNER JOIN trusturl ON trustweb.SessionID = trusturl.SessionID GROUP BY url").Scan(&urlCounts)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果未找到相同數據，則傳回 nil
		return nil, nil
	} else if result.Error != nil {
		// 處理其他查詢錯誤
		log.Println("Error querying database:", result.Error)
		return nil, result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return urlCounts, nil
}*/
