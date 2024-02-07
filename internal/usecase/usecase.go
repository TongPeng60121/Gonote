package usecase

import (
	"context"
	"gonote/internal/repository"
)

type TrustWebSiteUsecase interface {
	SearchTrustWeb(ctx context.Context, clientID string) ([]repository.Trustweb, error)
	//SearchTrustCount(db *gorm.DB) ([]repository.UrlCount, error)
}
