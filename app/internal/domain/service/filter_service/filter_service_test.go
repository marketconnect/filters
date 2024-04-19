package filter_service_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"testing"

	pb "filters/app/gen/proto"
	filter_service "filters/app/internal/domain/service/filter_service"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
)

// Setup a basic logger for testing

func getLogger() *zap.SugaredLogger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger.Sugar()
}

type FakeFilterDataProvider struct{}

func (f *FakeFilterDataProvider) GetDistinctNames(ctx context.Context, filterName string) ([]string, error) {
	return []string{"FilterValue1", "FilterValue2"}, nil
}

func (f *FakeFilterDataProvider) GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error) {
	return []uint32{5, 10}, nil
}

func (f *FakeFilterDataProvider) GetLemmasByFilterID(ctx context.Context, filterID int64) ([]*pb.LemmaByFilter, error) {
	if filterID == 1 {
		return []*pb.LemmaByFilter{
			{Lemma: "lemma1", TotalFrequency: 50},
			{Lemma: "lemma2", TotalFrequency: 30},
		}, nil
	}
	return nil, fmt.Errorf("no lemmas found for filter ID %d", filterID)
}

type FakeTokenManager struct{}

func (f *FakeTokenManager) Verify(accessToken string) (*uint64, error) {
	userID := uint64(123)
	return &userID, nil
}

func createContextWithMetadata(token string) context.Context {
	md := metadata.New(map[string]string{"authorization": token})
	return metadata.NewIncomingContext(context.Background(), md)
}

func TestGetFilterValues(t *testing.T) {
	ctx := createContextWithMetadata("validToken")

	service := filter_service.NewFilterService(&FakeFilterDataProvider{}, &FakeTokenManager{}, getLogger())
	req := &pb.GetFilterValuesReq{FilterName: "exampleFilter"}
	resp, err := service.GetFilterValues(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Values, 2)
}

func TestGetSearchQuery(t *testing.T) {
	ctx := createContextWithMetadata("validToken")
	service := filter_service.NewFilterService(&FakeFilterDataProvider{}, &FakeTokenManager{}, getLogger())
	req := &pb.GetSearchQueryReq{Queries: []string{"query1", "query2"}}
	resp, err := service.GetSearchQuery(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, []int32{5, 10}, resp.Frequencies)
}

func TestGetLemmasByFilterID(t *testing.T) {
	ctx := createContextWithMetadata("validToken")
	service := filter_service.NewFilterService(&FakeFilterDataProvider{}, &FakeTokenManager{}, getLogger())

	validReq := &pb.GetLemmasByFilterIDReq{FilterID: 1}
	invalidReq := &pb.GetLemmasByFilterIDReq{FilterID: 999}

	validResp, validErr := service.GetLemmasByFilterID(ctx, validReq)
	assert.NoError(t, validErr)
	assert.NotNil(t, validResp)
	assert.Len(t, validResp.Lemmas, 2)
	assert.Equal(t, "lemma1", validResp.Lemmas[0].Lemma)
	assert.Equal(t, int32(50), validResp.Lemmas[0].TotalFrequency)
	assert.Equal(t, "lemma2", validResp.Lemmas[1].Lemma)
	assert.Equal(t, int32(30), validResp.Lemmas[1].TotalFrequency)

	invalidResp, invalidErr := service.GetLemmasByFilterID(ctx, invalidReq)
	assert.Error(t, invalidErr)
	assert.Nil(t, invalidResp, "Response should be nil when an error occurs")
}
