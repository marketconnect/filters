package filter_data_provider_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"filters/app/internal/data_provider/filter_data_provider"
)

const (
	testDBConnectionString = "postgres://test_user:password@localhost:5432/test_db"
	testRedisAddr          = "localhost:6379"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), testDBConnectionString)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}

	// Clean up and prepare the database
	_, err = pool.Exec(context.Background(), "TRUNCATE filters RESTART IDENTITY")
	if err != nil {
		t.Fatalf("Unable to clean up database: %v", err)
	}

	// Insert test data
	_, err = pool.Exec(context.Background(), `INSERT INTO filters (name, filter_name) VALUES ('Test1', 'category'), ('Test2', 'category')`)
	if err != nil {
		t.Fatalf("Unable to insert test data: %v", err)
	}
	// Clean up and prepare the database
	_, err = pool.Exec(context.Background(), "TRUNCATE search_phrases RESTART IDENTITY")
	if err != nil {
		t.Fatalf("Unable to clean up database: %v", err)
	}

	// Insert test data for search phrases and their frequencies
	_, err = pool.Exec(context.Background(), `INSERT INTO search_phrases (phrase, frequency) VALUES ('phrase1', 10), ('phrase2', 20)`)
	if err != nil {
		t.Fatalf("Unable to insert test data: %v", err)
	}

	return pool
}

func setupTestRedis(t *testing.T) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     testRedisAddr,
		Password: "",
		DB:       0,
	})

	// Clean up Redis
	err := client.FlushDB(context.Background()).Err()
	if err != nil {
		t.Fatalf("Unable to clean up Redis: %v", err)
	}

	return client
}

func TestGetDistinctNames(t *testing.T) {
	ctx := context.Background()
	pool := setupTestDB(t)
	defer pool.Close()

	redisClient := setupTestRedis(t)
	defer redisClient.Close()

	// Initialize your filter storage with the pool (Postgres client)
	// This assumes that your NewFilterStorage function or similar
	// takes a database client (pool) and a Redis client as parameters.
	// If it doesn't currently, you'll need to adjust your implementation
	// to allow for Redis client injection, which is a good practice for testing.
	storage := filter_data_provider.NewFilterStorage(pool) // Update this as per your actual constructor

	names, err := storage.GetDistinctNames(ctx, "category")

	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"Test1", "Test2"}, names)

	// Check if result is cached
	cachedResult, err := redisClient.Get(ctx, "category").Result()
	assert.NoError(t, err)

	var cachedNames []string
	err = json.Unmarshal([]byte(cachedResult), &cachedNames)
	assert.NoError(t, err)
	assert.ElementsMatch(t, names, cachedNames)
}

func TestGetFrequencies(t *testing.T) {
	ctx := context.Background()
	dbPool := setupTestDB(t)
	defer dbPool.Close()

	// Assume search_phrases table is already populated with some test data
	storage := filter_data_provider.NewFilterStorage(dbPool)

	// Define test phrases that exist in your test data
	testPhrases := []string{"phrase1", "phrase2"}

	frequencies, err := storage.GetFrequencies(ctx, testPhrases)
	print(frequencies)
	assert.NoError(t, err)
	assert.NotEmpty(t, frequencies)
	res := []uint32{10, 20}
	// You may need to adjust the expected values based on your test data
	for i, frequency := range frequencies {
		// This check depends on your test data. Adjust accordingly.
		freq := res[i]
		assert.True(t, frequency == freq, "Expected frequency to be %v, but got %v", freq, frequency)
	}
}
