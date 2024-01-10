package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"

	"Sarva_Assignment/internal/core/logger"
)

func setupTestRedisDB() (*RedisDB, *miniredis.Miniredis, context.Context) {
	// Setup a miniredis server for testing
	miniRedis, err := miniredis.Run()
	if err != nil {
		panic(fmt.Sprintf("Error starting miniredis: %v", err))
	}

	// Create a Redis client using miniredis server
	client := redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
		DB:   0,
	})

	// Create a RedisDB instance for testing
	logger := logger.NewLogger(hclog.New(nil))
	redisDB := NewRedisDB(client, logger)

	// Context for testing
	ctx := context.TODO()

	return redisDB, miniRedis, ctx
}

func TestRedisDB_SaveMapping(t *testing.T) {
	redisDB, miniRedis, ctx := setupTestRedisDB()

	defer miniRedis.Close()

	defer redisDB.client.Close()

	// Test SaveMapping
	fileName := "testfile"
	fileSize := int64(100)

	err := redisDB.SaveMapping(ctx, fileName, fileSize)
	assert.NoError(t, err, "SaveMapping should not return an error")

	// Verify the value in Redis
	val, err := redisDB.client.Get(ctx, fmt.Sprintf("file:%s:size", fileName)).Int64()
	assert.NoError(t, err, "Error getting value from Redis")
	assert.Equal(t, fileSize, val, "Incorrect value retrieved from Redis")
}

func TestRedisDB_GetFileSize(t *testing.T) {
	// Setup
	redisDB, miniRedis, ctx := setupTestRedisDB()

	defer miniRedis.Close()

	defer redisDB.client.Close()

	// Test GetFileSize for a non-existent key
	fileName := "nonexistentfile"
	val, err := redisDB.GetFileSize(ctx, fileName)

	assert.NoError(t, err, "GetFileSize should not return an error")
	assert.Zero(t, val, "GetFileSize should return zero for a non-existent key")

	// Test GetFileSize for an existing key
	fileName = "testfile"
	fileSize := int64(200)

	err = redisDB.SaveMapping(ctx, fileName, fileSize)

	assert.NoError(t, err, "SaveMapping should not return an error")

	// Retrieve file size using GetFileSize
	val, err = redisDB.GetFileSize(ctx, fileName)

	assert.NoError(t, err, "GetFileSize should not return an error")
	assert.Equal(t, fileSize, val, "Incorrect file size retrieved from Redis")
}
