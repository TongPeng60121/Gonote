package repository

type TrustWebRepository interface {
	GetTrustWebSites(clientID string) ([]Trustweb, error)
	//GetUrlCounts(db *gorm.DB) ([]Trustweb, error)
}
