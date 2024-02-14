package repository

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type TrustWeb struct {
	Url         string
	Cdate       time.Time
	CdateString string
}

type UrlCount struct {
	Url   string
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

func (t *trustWebRepository) GetTrustWebSites(_ context.Context, clientID string) ([]TrustWeb, error) {
	var existingTrustWebs []TrustWeb
	result := t.db.Raw("SELECT Url, trusturl.Cdate FROM trustweb INNER JOIN trusturl ON trustweb.SessionID = trusturl.SessionID WHERE ClientID = ? GROUP BY Url", clientID).Scan(&existingTrustWebs)
	if result.Error == gorm.ErrRecordNotFound {
		// 如果未找到相同數據，則傳回 nil
		return nil, nil
	} else if result.Error != nil {
		// 處理其他查詢錯誤
		log.Println("Error querying database:", result.Error)
		return nil, result.Error
	}

	for i := range existingTrustWebs {
		existingTrustWebs[i].CdateString = existingTrustWebs[i].Cdate.Format("2006-01-02 15:04:05")
	}

	// 傳回查詢到的 TrustWeb 數據
	return existingTrustWebs, nil
}

func (t *trustWebRepository) GetUrlCounts(_ context.Context) ([]UrlCount, error) {
	var urlCounts []UrlCount
	result := t.db.Raw("SELECT Url, COUNT(*) as Count FROM trustweb INNER JOIN trusturl ON trustweb.SessionID = trusturl.SessionID GROUP BY url").Scan(&urlCounts)
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
}
