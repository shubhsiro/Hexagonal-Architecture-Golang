// redis_adapter.go
package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"Sarva_Assignment/internal/adapter/logger"
	coreDB "Sarva_Assignment/internal/core/database"
)

// RedisAdapter represents the adapter for Redis database.
type RedisAdapter struct {
	RedisDB *coreDB.RedisDB
	Logger  *logger.HCLoggerAdapter // Assuming your logger adapter is named HCLoggerAdapter
}

// NewRedisAdapter creates a new instance of RedisAdapter.
func NewRedisAdapter(redisDB *coreDB.RedisDB, logger *logger.HCLoggerAdapter) *RedisAdapter {
	return &RedisAdapter{
		RedisDB: redisDB,
		Logger:  logger,
	}
}

// SaveMapping saves fileName-fileSize mapping to Redis.
func (r *RedisAdapter) SaveMapping(fileName string, fileSize int64) error {
	err := r.RedisDB.SaveMapping(context.TODO(), fileName, fileSize)
	if err != nil {
		r.Logger.Log(fmt.Sprintf("Error saving mapping to Redis: %v", err))

		return errors.Errorf("failed to save mapping to Redis: %v", err)
	}

	return nil
}

// GetFileSize retrieves the file size from Redis based on the fileName.
func (r *RedisAdapter) GetFileSize(fileName string) (int64, error) {
	fileSize, err := r.RedisDB.GetFileSize(context.TODO(), fileName)
	if err != nil {
		if err == redis.Nil {
			r.Logger.Log(fmt.Sprintf("File not found in Redis for key: %s", fileName))

			return 0, fmt.Errorf("file not found in Redis")
		}

		r.Logger.Log(fmt.Sprintf("Error getting file size from Redis: %v", err))

		return 0, fmt.Errorf("failed to get file size from Redis: %v", err)
	}

	return fileSize, nil
}
