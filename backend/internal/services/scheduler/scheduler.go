package scheduler

import (
	"fmt"
	"time"

	"opflow-backend/internal/models"
	"opflow-backend/internal/services/crawler"
)

type SchedulerService struct {
	crawlerService *crawler.CrawlerService
}

func NewSchedulerService(crawlerService *crawler.CrawlerService) *SchedulerService {
	return &SchedulerService{
		crawlerService: crawlerService,
	}
}

func (s *SchedulerService) StartScheduledCrawling(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			fmt.Println("Starting scheduled crawling...")
			hotspots, err := s.crawlerService.ScrapeAllSources()
			if err == nil {
				fmt.Printf("Successfully scraped %d hotspots\n", len(hotspots))
				// 这里可以添加保存到数据库的逻辑
			} else {
				fmt.Printf("Scheduled crawling failed: %v\n", err)
			}
		}
	}()
}

func (s *SchedulerService) CheckScheduledPublish() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			fmt.Println("Checking scheduled content for publishing...")
			// 这里可以添加检查数据库中到期的内容并模拟发布的逻辑
			// 例如：SELECT * FROM contents WHERE scheduled_at <= NOW() AND status = 'scheduled'
			// 然后更新状态为 'published' 并打印日�?
			fmt.Println("模拟发布到平�?..")
		}
	}()
}
