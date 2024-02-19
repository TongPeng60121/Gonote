package repository

import "context"

type TrustWebRepository interface {
	GetTrustWebSites(ctx context.Context, clientID string) ([]TrustWeb, error)
	GetUrlCounts(ctx context.Context, limit string) ([]UrlCount, error)
	GetTrustWebFromDB(ctx context.Context, sessionID string, clientID string) error
	GetTrustUrlFromDB(ctx context.Context, sessionID string, Url string) (string, error)
	CreateTrustWeb(ctx context.Context, data TrustWebTable) error
	CreateTrustUrl(ctx context.Context, data TrustUrlTable) error
}
