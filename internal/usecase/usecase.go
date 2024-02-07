package usecase

import (
	"gonote/internal/repository"
)

type TrustWebSiteUsecase interface {
	SearchTrustWeb(clientID string) ([]repository.Trustweb, error)
	//SearchTrustCount(db *gorm.DB) ([]repository.UrlCount, error)
}
