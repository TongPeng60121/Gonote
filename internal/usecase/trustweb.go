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

// 使用 ClientID 查詢
func (t *trustWebSiteUsecase) SearchTrustWeb(ctx context.Context, clientID int) ([]repository.TrustWeb, error) {
	trustweb, err := t.trustWebSiteRepo.GetTrustWebSites(ctx, clientID)
	if err != nil {
		// 處理錯誤
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
		// 處理錯誤
		return nil, err
	}
	return TrustCount, nil
}

func (t *trustWebSiteUsecase) InsertSessionToDB(ctx context.Context, session []repository.TrustJson) error {

	for _, trustWeb := range session {
		// 檢查是否已存在相同的 TrustWeb 記錄
		result, err := t.trustWebSiteRepo.GetTrustWebFromDB(ctx, trustWeb.SessionID, trustWeb.ClientID)

		if err != nil {
			return err
		} else if result {
			newTrustWeb := repository.TrustWebTable{
				Session_id: trustWeb.SessionID,
				Client_id:  trustWeb.ClientID,
				Cdate:      time.Now(),
			}
			if err := t.trustWebSiteRepo.CreateTrustWeb(ctx, newTrustWeb); err != nil {
				// 返回新增錯誤
				return err
			}
		}

		for _, trustUrl := range trustWeb.TrustWeb {
			result, err := t.trustWebSiteRepo.GetTrustUrlFromDB(ctx, trustWeb.SessionID, trustUrl.Url)
			if err != nil {
				return err
			} else if result {
				newTrustUrl := repository.TrustUrlTable{
					Session_id: trustWeb.SessionID,
					Url:        trustUrl.Url,
					Cdate:      time.Now(),
				}
				if err := t.trustWebSiteRepo.CreateTrustUrl(ctx, newTrustUrl); err != nil {
					// 返回新增錯誤
					return err
				}
			}
		}
	}

	return nil
}
