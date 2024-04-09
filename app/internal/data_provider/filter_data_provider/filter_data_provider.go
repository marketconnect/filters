package filter_data_provider

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
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
