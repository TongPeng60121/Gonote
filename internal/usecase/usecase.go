package usecase

import (
	"context"
	"gonote/internal/repository"
)

type TrustWebSiteUsecase interface {
	SearchTrustWeb(ctx context.Context, clientID string) ([]repository.TrustWeb, error)
	SearchTrustCount(ctx context.Context) ([]repository.UrlCount, error)
}
