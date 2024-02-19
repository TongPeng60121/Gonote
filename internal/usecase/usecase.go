package usecase

import (
	"context"
	"gonote/internal/repository"
)

type TrustWebSiteUsecase interface {
	SearchTrustWeb(ctx context.Context, clientID string) ([]repository.TrustWeb, error)
	SearchTrustCount(ctx context.Context, limit string) ([]repository.UrlCount, error)
	InsertSessionToDB(ctx context.Context, session []repository.TrustJson) error
}
