package filter_service

import (
	"context"
	pb "filters/app/gen/proto"

	"github.com/marketconnect/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/status"
)

type FilterDataProvider interface {
	GetDistinctNames(ctx context.Context, filterName string) ([]string, error)
	GetFrequencies(ctx context.Context, phrases []string) ([]uint32, error)
	GetKeywordsByFilter(ctx context.Context, filterID int64, limit int, offset int) (*pb.GetKeywordsByFilterResp, error)
}
type TokenManager interface {
	Verify(accessToken string) (*uint64, error)
}

type FilterService struct {
	filterDataProvider FilterDataProvider
	logger             logger.Logger
	tokenManager       TokenManager

	pb.UnimplementedFilterServiceServer
}

func NewFilterService(filterDataProvider FilterDataProvider, tokenManager TokenManager, logger logger.Logger) *FilterService {
	return &FilterService{
		filterDataProvider: filterDataProvider,
		tokenManager:       tokenManager,
		logger:             logger,
	}
}

func (service *FilterService) GetFilterValues(ctx context.Context, req *pb.GetFilterValuesReq) (*pb.GetFilterValuesResp, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		service.logger.Error("metadata is not provided")
		return &pb.GetFilterValuesResp{}, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		service.logger.Error("authorization token is not provided")
		return &pb.GetFilterValuesResp{}, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userID, err := service.tokenManager.Verify(accessToken)
	if err != nil {
		service.logger.Error("access token is invalid (userID): %v", err, userID)
		return &pb.GetFilterValuesResp{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	// Validate input parameters
	if req == nil {
		service.logger.Error("request is nil (userID: %d)", userID)
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	// Fetch filter values
	filterValues, err := service.filterDataProvider.GetDistinctNames(ctx, req.FilterName)
	if err != nil {
		service.logger.Error("could not add subscription info (userID: %d): %v", err, userID)
		return nil, status.Errorf(codes.Internal, "could not add subscription info: %v", err)
	}

	// Return the current price
	return &pb.GetFilterValuesResp{
		Values: filterValues,
	}, nil
}

func (service *FilterService) GetSearchQuery(ctx context.Context, req *pb.GetSearchQueryReq) (*pb.GetSearchQueryResp, error) {
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
	if len(req.Queries) == 0 {
		service.logger.Error("request is nil or queries are empty")
		return nil, status.Error(codes.InvalidArgument, "request is nil or queries are empty")
	}

	// Fetch frequencies for the provided queries
	frequencies, err := service.filterDataProvider.GetFrequencies(ctx, req.Queries)
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

func (service *FilterService) GetKeywordsByFilter(ctx context.Context, req *pb.GetKeywordsByFilterReq) (*pb.GetKeywordsByFilterResp, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		service.logger.Error("metadata is not provided")
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		service.logger.Error("authorization token is not provided")
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	userID, err := service.tokenManager.Verify(accessToken)
	if err != nil {
		service.logger.Error("access token is invalid (userID): %v", err, userID)
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if req == nil {
		service.logger.Error("request is nil")
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	if req.FilterID == 0 {
		service.logger.Error("Filter ID is required")
		return nil, status.Error(codes.InvalidArgument, "Filter ID is required")
	}

	// Fetch keywords by filter
	keywordsResp, err := service.filterDataProvider.GetKeywordsByFilter(ctx, req.FilterID, int(req.Limit), int(req.Offset))
	if err != nil {
		service.logger.Error("could not get keywords by filter (userID: %d): %v", userID, err)
		return nil, status.Errorf(codes.Internal, "could not get keywords by filter: %v", err)
	}

	return keywordsResp, nil
}
