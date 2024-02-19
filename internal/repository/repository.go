package repository

import "context"

type TrustWebRepository interface {
	GetTrustWebSites(ctx context.Context, clientID string) ([]TrustWeb, error)
	GetUrlCounts(ctx context.Context, limit string) ([]UrlCount, error)
	GetTrustWebFromDB(ctx context.Context, sessionID string, clientID string) (string, error)
	CreateTrustWeb(ctx context.Context, data TrustWebTable) error
}
