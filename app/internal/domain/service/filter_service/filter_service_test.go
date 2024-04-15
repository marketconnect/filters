package filter_service_test

import (
	"context"
	pb "filters/app/gen/proto"

	"testing"

	filter_service "filters/app/internal/domain/service/filter_service"

	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc/metadata"
)

// FakeFilterDataProvider simulates the filter data provider.
type FakeFilterDataProvider struct{}

func (f *FakeFilterDataProvider) GetDistinctNames(ctx context.Context, filterName string) ([]string, error) {
	// Simulate retrieving a list of filter names.
	return []string{"FilterValue1", "FilterValue2"}, nil
}

func (f *FakeFilterDataProvider) GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error) {
	// Simulate retrieving frequencies for the provided phrases.
	// The frequencies here are arbitrary and should match your test cases.
	return []uint32{5, 10}, nil
}

func (f *FakeFilterDataProvider) GetKeywordsByFilter(ctx context.Context, filterID int64) (*pb.GetKeywordsByFilterResp, error) {
	// Return mock keywords based on a filterID
	keywords := []*pb.KeywordByFilter{
		{Normquery: "keyword1", Frequency: 10, Competition: 100, Count: 500},
		{Normquery: "keyword2", Frequency: 20, Competition: 150, Count: 450},
	}
	return &pb.GetKeywordsByFilterResp{Keywords: keywords}, nil
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

func TestGetFilterValues(t *testing.T) {
	ctx := createContextWithMetadata("validToken")

	service := filter_service.NewFilterService(
		&FakeFilterDataProvider{},
		&FakeTokenManager{},
		nil, // Assuming logger is not critical for the test. Use a real logger if needed.
	)

	req := &pb.GetFilterValuesReq{
		FilterName: "exampleFilter",
	}

	resp, err := service.GetFilterValues(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Values, 2) // Assuming the FakeFilterDataProvider returns 2 filter values
}

func TestGetSearchQuery(t *testing.T) {
	ctx := createContextWithMetadata("validToken")

	service := filter_service.NewFilterService(
		&FakeFilterDataProvider{},
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

func TestGetKeywordsByFilter(t *testing.T) {
	ctx := createContextWithMetadata("validToken")

	service := filter_service.NewFilterService(
		&FakeFilterDataProvider{},
		&FakeTokenManager{},
		nil, // Logger is not needed for this test
	)

	req := &pb.GetKeywordsByFilterReq{
		FilterID: 1, // Example filter ID for testing
	}

	resp, err := service.GetKeywordsByFilter(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Keywords, 2) // Check if two keywords are returned as mocked

	// Verify the details of the first keyword
	assert.Equal(t, "keyword1", resp.Keywords[0].Normquery)
	assert.Equal(t, int32(10), resp.Keywords[0].Frequency)
	assert.Equal(t, int32(100), resp.Keywords[0].Competition)
	assert.Equal(t, int32(500), resp.Keywords[0].Count)

	// Verify the details of the second keyword
	assert.Equal(t, "keyword2", resp.Keywords[1].Normquery)
	assert.Equal(t, int32(20), resp.Keywords[1].Frequency)
	assert.Equal(t, int32(150), resp.Keywords[1].Competition)
	assert.Equal(t, int32(450), resp.Keywords[1].Count)
}
