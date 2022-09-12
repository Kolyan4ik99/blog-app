package app

import (
	"context"
	"time"

	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
)

type TimeExpiryScheduler struct {
	repoPost repository.PostInterface
	ticker   time.Duration
}

func NewTimeExpiryScheduler(repoPost repository.PostInterface, ticker time.Duration) *TimeExpiryScheduler {
	return &TimeExpiryScheduler{repoPost: repoPost, ticker: ticker}
}

func (t *TimeExpiryScheduler) Run(ctx context.Context) {
	tick := time.NewTicker(t.ticker)

	logger.Logger.Infoln("Start scheduler")
	for {
		select {
		case <-tick.C:
			logger.Logger.Infoln("Start tick")
			ttl := time.Now()
			posts, err := t.repoPost.GetAllPostTTLBefore(ctx, ttl)
			if err != nil {
				logger.Logger.Errorln(err)
			} else {
				for _, post := range posts {
					err = t.repoPost.DeleteById(ctx, post.Id)
					if err != nil {
						logger.Logger.Errorln(err)
					}
					logger.Logger.Infof("Post id=[%d] was deleted by ttl", post.Id)
				}
			}
			logger.Logger.Infoln("End tick")
		case <-ctx.Done():
			logger.Logger.Errorln("Scheduler stop")
			return
		}
	}
}
