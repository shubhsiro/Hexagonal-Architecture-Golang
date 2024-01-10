// redis_db.go
package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"Sarva_Assignment/internal/core/logger"
)

// RedisDB represents the core logic for Redis database interactions.
type RedisDB struct {
	client *redis.Client
	logger *logger.Log
}

// NewRedisDB creates a new instance of RedisDB.
func NewRedisDB(client *redis.Client, logger *logger.Log) *RedisDB {
	return &RedisDB{
		client: client,
		logger: logger,
	}
}

// SaveMapping saves fileName-fileSize mapping to Redis.
func (r *RedisDB) SaveMapping(ctx context.Context, fileName string, fileSize int64) error {
	key := fmt.Sprintf("file:%s:size", fileName)

	if err := r.client.Set(ctx, key, fileSize, 0).Err(); err != nil {
		r.logger.Log(fmt.Sprintf("Error saving mapping to Redis: %v", err))

		return fmt.Errorf("failed to save mapping to Redis: %v", err)
	}

	r.logger.Debug(fmt.Sprintf("Mapping saved to Redis. Key: %s, Value: %d", key, fileSize))

	return nil
}

// GetFileSize retrieves the file size from Redis based on the fileName.
func (r *RedisDB) GetFileSize(ctx context.Context, fileName string) (int64, error) {
	key := fmt.Sprintf("file:%s:size", fileName)

	val, err := r.client.Get(ctx, key).Int64()

	if err == redis.Nil {
		r.logger.Debug(fmt.Sprintf("File not found in Redis for key: %s", fileName))

		return 0, nil // Return nil for a non-existent key
	} else if err != nil {
		r.logger.Log(fmt.Sprintf("Error getting file size from Redis: %v", err))

		return 0, fmt.Errorf("failed to get file size from Redis: %v", err)
	}

	r.logger.Debug(fmt.Sprintf("File size retrieved from Redis. Key: %s, Value: %d", key, val))

	return val, nil
}
