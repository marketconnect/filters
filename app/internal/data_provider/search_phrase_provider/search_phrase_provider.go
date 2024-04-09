package search_phrase_provider

import (
	"context"
	"fmt"

	client "github.com/marketconnect/db_client/postgresql"
)

type searchPhraseStorage struct {
	client client.PostgreSQLClient
}

func NewSearchPhraseStorage(client client.PostgreSQLClient) *searchPhraseStorage {
	return &searchPhraseStorage{client: client}
}
func (s *searchPhraseStorage) GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error) {
	valuesStr := "VALUES "
	for _, phrase := range phrases {
		valuesStr += fmt.Sprintf("('%s'),", phrase)
	}
	valuesStr = valuesStr[:len(valuesStr)-1] // Remove the last comma

	// Construct the full SQL query
	query := fmt.Sprintf(`
    WITH input_phrases(phrase) AS (
        %s
    )
    SELECT ip.phrase, COALESCE(sp.frequency, 0) AS frequency
    FROM input_phrases ip
    LEFT JOIN search_phrases sp ON ip.phrase = sp.phrase;
    `, valuesStr)

	// Execute the query without additional arguments
	rows, err := s.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	frequencyMap := make(map[string]uint32)
	for _, phrase := range phrases {
		frequencyMap[phrase] = 0 // Initialize all frequencies to 0
	}

	for rows.Next() {
		var phrase string
		var frequency uint32
		if err := rows.Scan(&phrase, &frequency); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		frequencyMap[phrase] = frequency
	}

	frequencies := make([]uint32, len(phrases))
	for i, phrase := range phrases {
		frequencies[i] = frequencyMap[phrase]
	}

	return frequencies, nil
}
