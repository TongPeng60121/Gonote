package repository

import "context"

type TrustWebRepository interface {
	GetTrustWebSites(ctx context.Context, clientID int) ([]TrustWeb, error)
	GetUrlCounts(ctx context.Context, limit string) ([]UrlCount, error)
	GetTrustWebFromDB(ctx context.Context, sessionID int, clientID int) (bool, error)
	GetTrustUrlFromDB(ctx context.Context, sessionID int, Url string) (bool, error)
	CreateTrustWeb(ctx context.Context, data TrustWebTable) error
	CreateTrustUrl(ctx context.Context, data TrustUrlTable) error
}
