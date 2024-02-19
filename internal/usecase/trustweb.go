package usecase

import (
	"context"
	"gonote/internal/repository"
	"time"
)

type trustWebSiteUsecase struct {
	trustWebSiteRepo repository.TrustWebRepository
}

// NewTrustWebSiteUsecase 创建 TrustWebSiteUsecase 实例
func NewTrustWebSiteUsecase(repo repository.TrustWebRepository) TrustWebSiteUsecase {
	return &trustWebSiteUsecase{
		trustWebSiteRepo: repo,
	}
}

// 使用 ClientID 进行获取
func (t *trustWebSiteUsecase) SearchTrustWeb(ctx context.Context, clientID string) ([]repository.TrustWeb, error) {
	trustweb, err := t.trustWebSiteRepo.GetTrustWebSites(ctx, clientID)
	if err != nil {
		// 处理错误
		return nil, err
	}

	for i := range trustweb {
		trustweb[i].CdateString = trustweb[i].Cdate.Format("2006-01-02 15:04:05")
	}

	return trustweb, nil
}

// 統計被使用的信任網站數量
func (t *trustWebSiteUsecase) SearchTrustCount(ctx context.Context, limit string) ([]repository.UrlCount, error) {
	TrustCount, err := t.trustWebSiteRepo.GetUrlCounts(ctx, limit)
	if err != nil {
		// 处理错误
		return nil, err
	}
	return TrustCount, nil
}

func (t *trustWebSiteUsecase) InsertSessionToDB(ctx context.Context, session []repository.TrustJson) error {

	for _, trustWeb := range session {
		// 檢查是否已存在相同的 TrustWeb 記錄
		trw_id, err := t.trustWebSiteRepo.GetTrustWebFromDB(ctx, trustWeb.SessionID, trustWeb.ClientID)
		if trw_id != "" {
			// 如果已存在相同數據，則直接傳回，不執行插入操作
			continue
		} else if err != nil {
			return err
		} else {
			newTrustWeb := repository.TrustWebTable{
				Session_id: trustWeb.SessionID,
				Client_id:  trustWeb.ClientID,
				Cdate:      time.Now(),
			}
			if err := t.trustWebSiteRepo.CreateTrustWeb(ctx, newTrustWeb); err != nil {
				// 处理插入数据库时的错误
				return err
			}
		}

		// 插入 TrustWeb 數據
		/*newTrustWeb := models.Trustweb{
			SessionID: session.SessionID,
			ClientID:  session.ClientID,
			Url:       trustWeb.Url,
			Cdate:     time.Now(),
		}

		if err := repository.CreateTrustWeb(db, &newTrustWeb); err != nil {
			// 处理插入数据库时的错误
			c.JSON(500, gin.H{"error": "Failed to insert data into database"})
			return
		}*/
	}

	return nil
}
