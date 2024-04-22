package filter_data_provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"time"

	pb "filters/app/gen/proto"

	redis "github.com/go-redis/redis/v8"

	client "github.com/marketconnect/db_client/postgresql"
)

const (
	selectQuery = `SELECT DISTINCT name FROM filters WHERE filter_name = $1`
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

type filterStorage struct {
	client client.PostgreSQLClient
}

func NewFilterStorage(client client.PostgreSQLClient) *filterStorage {
	return &filterStorage{client: client}
}

func (s *filterStorage) GetDistinctNames(ctx context.Context, filterName string) ([]string, error) {
	// Попытка получить результат из Redis с учетом контекста
	cachedResult, err := redisClient.Get(ctx, filterName).Result()
	if err == nil {
		// Десериализация кэшированного результата
		var names []string
		json.Unmarshal([]byte(cachedResult), &names)
		return names, nil
	}

	// Получение результата из PostgreSQL
	rows, err := s.client.Query(ctx, selectQuery, filterName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	// Сериализация результата и его сохранение в Redis на 24 часа
	serializedResult, _ := json.Marshal(names)
	redisClient.Set(ctx, filterName, serializedResult, 24*time.Hour) // Установка TTL в 24 часа

	return names, nil
}

func (s *filterStorage) GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error) {
	valuesStr := "VALUES "
	for _, phrase := range phrases {
		valuesStr += fmt.Sprintf("('%s'),", phrase)
	}
	valuesStr = valuesStr[:len(valuesStr)-1] // Remove the last comma

	// Construct the full SQL query
	query := fmt.Sprintf(`
    WITH input_phrases(kw) AS (
        %s
    )
    SELECT ip.kw, COALESCE(sp.freq, 0) AS freq
    FROM input_phrases ip
    LEFT JOIN search_phrases sp ON ip.kw = sp.kw;
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

func (s *filterStorage) GetLemmasByFilterID(ctx context.Context, filterID int64) ([]*pb.LemmaByFilter, error) {
	query := `
	SELECT 
    l.id AS lemma_id,
    l.lemma,
    SUM(sp.freq) AS total_frequency
	FROM categories c
	JOIN kw ON c.kw_id = kw.id
	JOIN kw_lemmas kl ON kw.id = kl.kw_id
	JOIN lemmas l ON kl.lemma_id = l.id
	JOIN search_phrases sp ON kw.normquery = sp.kw
	WHERE c.filter_id = $1
	GROUP BY l.id, l.lemma
	ORDER BY total_frequency DESC;

    `

	rows, err := s.client.Query(ctx, query, filterID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for filter ID %d: %w", filterID, err)
	}
	defer rows.Close()

	var lemmas []*pb.LemmaByFilter
	for rows.Next() {
		var lemma pb.LemmaByFilter
		if err := rows.Scan(&lemma.LemmaID, &lemma.Lemma, &lemma.TotalFrequency); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		lemmas = append(lemmas, &lemma)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return lemmas, nil
}

func (s *filterStorage) GetKeywordsByLemmas(ctx context.Context, req *pb.GetKeywordsByLemmasReq) (*pb.GetKeywordsByLemmasResp, error) {
	// Converting lemmasIDs from []int64 to a string suitable for IN clause
	lemmaIDs := make([]string, len(req.LemmasIDs))
	for i, id := range req.LemmasIDs {
		lemmaIDs[i] = strconv.FormatInt(id, 10)
	}
	lemmaArray := strings.Join(lemmaIDs, ",")

	query := `
		SELECT DISTINCT
		l.id AS lemma_id,
		l.lemma,
		k.normquery AS keyword,
		sp.freq
		FROM lemmas l
		JOIN kw_lemmas kl ON l.id = kl.lemma_id
		JOIN kw k ON kl.kw_id = k.id
		JOIN categories c ON k.id = c.kw_id
		JOIN search_phrases sp ON k.normquery = sp.kw
		WHERE l.id IN (` + lemmaArray + `)
		AND c.filter_id = $3 
		ORDER BY sp.freq DESC
		LIMIT $1 OFFSET $2;
    `

	rows, err := s.client.Query(ctx, query, req.Limit, req.Offset, req.FilterID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query with lemmas: %w", err)
	}
	defer rows.Close()

	resp := &pb.GetKeywordsByLemmasResp{}
	for rows.Next() {
		var keyword pb.KeywordByLemma
		err := rows.Scan(&keyword.LemmaID, &keyword.Lemma, &keyword.Keyword, &keyword.Freq)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		resp.Keywords = append(resp.Keywords, &keyword)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return resp, nil
}
