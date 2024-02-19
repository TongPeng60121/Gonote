package repository

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
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

type TrustJson struct {
	SessionID string
	ClientID  string
	TrustWeb  []TrustUrl
}

type TrustUrl struct {
	Url string
}

type TrustWebTable struct {
	Session_id string
	Client_id  string
	Cdate      time.Time
}

type TrustUrlTable struct {
	Session_id string
	Url        string
	Cdate      time.Time
}

func NewTrustWebRepository(db *gorm.DB) TrustWebRepository {
	return &trustWebRepository{
		db: db,
	}
}

func (t *trustWebRepository) GetTrustWebSites(ctx context.Context, clientID string) ([]TrustWeb, error) {
	var existingTrustWebs []TrustWeb
	// 預備 SQL 查詢
	stmt, err := t.db.DB().PrepareContext(ctx, "SELECT url, trusturl.cdate FROM trustweb INNER JOIN trusturl ON trustweb.session_id = trusturl.session_id WHERE client_id = ? GROUP BY url")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// 執行查詢
	rows, err := stmt.QueryContext(ctx, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 掃描結果到 existingTrustWebs 切片
	for rows.Next() {
		var trustWeb TrustWeb
		if err := rows.Scan(&trustWeb.Url, &trustWeb.Cdate); err != nil {
			return nil, err
		}
		existingTrustWebs = append(existingTrustWebs, trustWeb)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 傳回查詢到的 TrustWeb 數據
	return existingTrustWebs, nil
}

func (t *trustWebRepository) GetUrlCounts(_ context.Context, limit string) ([]UrlCount, error) {
	var urlCounts []UrlCount
	result := t.db.Raw("SELECT url, COUNT(*) as count FROM trustweb INNER JOIN trusturl ON trustweb.session_id = trusturl.session_id GROUP BY url LIMIT ?", limit).Scan(&urlCounts)
	if result.Error != nil {
		return nil, result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return urlCounts, nil
}

func (t *trustWebRepository) GetTrustWebFromDB(_ context.Context, sessionID string, clientID string) error {
	result := t.db.Raw("SELECT trw_id FROM trustweb WHERE session_id = ? AND client_id = ? LIMIT 1", sessionID, clientID)
	if result.Error != nil {
		return result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return nil
}

func (t *trustWebRepository) GetTrustUrlFromDB(_ context.Context, sessionID string, url string) (string, error) {
	tru_id := ""
	result := t.db.Raw("SELECT trw_id FROM trusturl WHERE session_id = ? AND url = ? LIMIT 1", sessionID, url).Scan(&tru_id)
	if result.Error != nil {
		return "", result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return tru_id, nil
}

func (t *trustWebRepository) CreateTrustWeb(_ context.Context, data TrustWebTable) error {
	result := t.db.Table("trustweb").Create(&data)
	if result.Error != nil {
		return result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return nil
}

func (t *trustWebRepository) CreateTrustUrl(_ context.Context, data TrustUrlTable) error {
	result := t.db.Table("trusturl").Create(&data)
	if result.Error != nil {
		return result.Error
	}
	// 傳回查詢到的 TrustWeb 數據
	return nil
}
