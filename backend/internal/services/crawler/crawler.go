package crawler

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
    
    "opflow-backend/internal/models"
    "opflow-backend/internal/pkg/utils"
)

type Service struct{}

func New() *Service {
    return &Service{}
}

func (s *Service) FetchZhihu() ([]models.Hotspot, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    
    req, err := http.NewRequest("GET", "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total", nil)
    if err != nil {
        return getMockZhihu(), nil
    }
    
    req.Header.Set("User-Agent", utils.GetRandomUserAgent())
    req.Header.Set("Referer", "https://www.zhihu.com")
    
    resp, err := client.Do(req)
    if err != nil {
        return getMockZhihu(), nil
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return getMockZhihu(), nil
    }
    
    // 解析知乎API返回的JSON
    var result struct {
        Data []struct {
            Target struct {
                Title   string `json:"title"`
                URL     string `json:"url"`
                ID      int64  `json:"id"`
            } `json:"target"`
            DetailText string `json:"detail_text"` // 热度文本如"123万热度"
        } `json:"data"`
    }
    
    if err := json.Unmarshal(body, &result); err != nil {
        return getMockZhihu(), nil
    }
    
    var hotspots []models.Hotspot
    for i, item := range result.Data {
        if i >= 10 { // 只取前10条
            break
        }
        
        score := 100 - i*5
        if score < 50 {
            score = 50
        }
        
        hotspots = append(hotspots, models.Hotspot{
            ID:        fmt.Sprintf("zhihu_%d", item.Target.ID),
            Title:     item.Target.Title,
            URL:       item.Target.URL,
            Source:    "zhihu",
            Score:     score,
            CreatedAt: time.Now(),
        })
    }
    
    if len(hotspots) == 0 {
        return getMockZhihu(), nil
    }
    
    return hotspots, nil
}

func (s *Service) FetchWeibo() ([]models.Hotspot, error) {
    // 微博先返回模拟数据（反爬太严）
    return []models.Hotspot{
        {ID: "weibo_1", Title: "微博热搜占位1", URL: "https://weibo.com", Source: "weibo", Score: 100, CreatedAt: time.Now()},
        {ID: "weibo_2", Title: "微博热搜占位2", URL: "https://weibo.com", Source: "weibo", Score: 95, CreatedAt: time.Now()},
    }, nil
}

// 保底模拟数据（API失败时用）
func getMockZhihu() []models.Hotspot {
    return []models.Hotspot{
        {ID: "zhihu_1", Title: "AI发展趋势2026", URL: "https://zhihu.com", Source: "zhihu", Score: 100, CreatedAt: time.Now()},
        {ID: "zhihu_2", Title: "程序员职业规划", URL: "https://zhihu.com", Source: "zhihu", Score: 95, CreatedAt: time.Now()},
        {ID: "zhihu_3", Title: "云原生技术解析", URL: "https://zhihu.com", Source: "zhihu", Score: 90, CreatedAt: time.Now()},
    }
}
