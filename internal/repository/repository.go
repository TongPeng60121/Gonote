package repository

import "context"

type TrustWebRepository interface {
	GetTrustWebSites(ctx context.Context, clientID string) ([]Trustweb, error)
	//GetUrlCounts(db *gorm.DB) ([]Trustweb, error)
}
