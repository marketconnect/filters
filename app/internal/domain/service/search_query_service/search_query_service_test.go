package search_query_service_test

import (
	"context"
	pb "filters/app/gen/proto"
	search_query_service "filters/app/internal/domain/service/search_query_service"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

// FakeSearchPhraseDataProvider simulates the search phrase data provider.
type FakeSearchPhraseDataProvider struct{}

func (f *FakeSearchPhraseDataProvider) GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error) {
	// Simulate retrieving frequencies for the provided phrases.
	// The frequencies here are arbitrary and should match your test cases.
	return []uint32{5, 10}, nil
}

// FakeTokenManager simulates the token manager.
type FakeTokenManager struct{}

func (f *FakeTokenManager) Verify(accessToken string) (*uint64, error) {
	// Simulate successful token verification.
	userID := uint64(123)
	return &userID, nil
}

// createContextWithMetadata creates a context with metadata simulating an incoming request with an authorization token.
func createContextWithMetadata(token string) context.Context {
	md := metadata.New(map[string]string{"authorization": token})
	return metadata.NewIncomingContext(context.Background(), md)
}

func TestGetSearchQuery(t *testing.T) {
	ctx := createContextWithMetadata("validToken")

	service := search_query_service.NewSearchQueryService(
		&FakeSearchPhraseDataProvider{},
		&FakeTokenManager{},
		nil, // Assuming logger is not critical for the test. Use a real logger if needed.
	)

	req := &pb.GetSearchQueryReq{
		Queries: []string{"query1", "query2"},
	}

	resp, err := service.GetSearchQuery(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// The expected frequencies should match what's defined in FakeSearchPhraseDataProvider.GetFrequencies
	expectedFrequencies := []int32{5, 10}
	assert.Equal(t, expectedFrequencies, resp.Frequencies)
}
