package search_phrase_provider_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"filters/app/internal/data_provider/search_phrase_provider"
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

func TestGetFrequencies(t *testing.T) {
	ctx := context.Background()
	dbPool := setupTestDB(t)
	defer dbPool.Close()

	// Assume search_phrases table is already populated with some test data
	storage := search_phrase_provider.NewSearchPhraseStorage(dbPool)

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
