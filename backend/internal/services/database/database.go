package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"opflow-backend/internal/config"
	"opflow-backend/internal/models"
)

type DatabaseService struct {
	DB     *gorm.DB
	config *config.Config
}

func NewDatabaseService(config *config.Config) (*DatabaseService, error) {
	db, err := gorm.Open(sqlite.Open(config.Database.Path), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&models.Hotspot{}, &models.Content{}, &models.Schedule{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &DatabaseService{
		DB:     db,
		config: config,
	}, nil
}

func (s *DatabaseService) GetDB() *gorm.DB {
	return s.DB
}

func (s *DatabaseService) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
