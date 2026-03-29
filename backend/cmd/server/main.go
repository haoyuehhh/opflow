package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "opflow-backend/internal/handlers/api"
    "opflow-backend/internal/services/content"
    "opflow-backend/internal/services/crawler"
)

func main() {
    r := gin.Default()
    
    // CORS 配置
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    }))
    
    crawlerSvc := crawler.New()
    contentGen := content.NewGenerator("sk-test-key")
    hotspotHandler := api.NewHotspotHandler(crawlerSvc, contentGen)
    
    v1 := r.Group("/api/v1")
    {
        v1.GET("/hotspots", hotspotHandler.List)
        v1.POST("/hotspots/:id/generate", hotspotHandler.Generate)
        v1.GET("/health", func(c *gin.Context) {
            c.JSON(200, gin.H{"status": "ok"})
        })
    }
    
    r.Run(":8080")
}
