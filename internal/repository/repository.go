package repository

import "context"

type TrustWebRepository interface {
	GetTrustWebSites(ctx context.Context, clientID string) ([]TrustWeb, error)
	GetUrlCounts(ctx context.Context) ([]UrlCount, error)
}
