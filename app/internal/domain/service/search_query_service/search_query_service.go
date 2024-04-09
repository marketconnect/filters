package search_query_service

import (
	"context"
	pb "filters/app/gen/proto"

	"github.com/marketconnect/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type SearchPhraseDataProvider interface {
	GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error)
}
type TokenManager interface {
	Verify(accessToken string) (*uint64, error)
}
type SearchQueryService struct {
	searchPhraseDataProvider SearchPhraseDataProvider
	logger                   logger.Logger
	tokenManager             TokenManager
	pb.UnimplementedSearchQueryServiceServer
}

func NewSearchQueryService(searchPhraseDataProvider SearchPhraseDataProvider, tokenManager TokenManager, logger logger.Logger) *SearchQueryService {
	return &SearchQueryService{
		searchPhraseDataProvider: searchPhraseDataProvider,
		tokenManager:             tokenManager,
		logger:                   logger,
	}
}

func (service *SearchQueryService) GetSearchQuery(ctx context.Context, req *pb.GetSearchQueryReq) (*pb.GetSearchQueryResp, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		service.logger.Error("metadata is not provided")
		return &pb.GetSearchQueryResp{}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		service.logger.Error("authorization token is not provided")
		return &pb.GetSearchQueryResp{}, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userID, err := service.tokenManager.Verify(accessToken)
	if err != nil {
		service.logger.Error("access token is invalid (userID): %v", err, userID)
		return &pb.GetSearchQueryResp{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	// Validate input parameters
	if req == nil {
		service.logger.Error("request is nil (userID: %d)", userID)
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	// Validate input parameters
	if req == nil || len(req.Queries) == 0 {
		service.logger.Error("request is nil or queries are empty")
		return nil, status.Error(codes.InvalidArgument, "request is nil or queries are empty")
	}

	// Fetch frequencies for the provided queries
	frequencies, err := service.searchPhraseDataProvider.GetFrequencies(ctx, req.Queries)
	if err != nil {
		service.logger.Error("could not get frequencies: %v", err)
		return nil, status.Errorf(codes.Internal, "could not get frequencies: %v", err)
	}

	// Convert frequencies from uint32 to int32 for the response
	freqInt32 := make([]int32, len(frequencies))
	for i, freq := range frequencies {
		freqInt32[i] = int32(freq)
	}

	return &pb.GetSearchQueryResp{
		Frequencies: freqInt32,
	}, nil
}
