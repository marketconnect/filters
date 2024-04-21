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
	GetLemmasByFilterID(ctx context.Context, filterID int64) ([]*pb.LemmaByFilter, error)
	GetKeywordsByLemmas(ctx context.Context, req *pb.GetKeywordsByLemmasReq) (*pb.GetKeywordsByLemmasResp, error)
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

func (service *FilterService) GetLemmasByFilterID(ctx context.Context, req *pb.GetLemmasByFilterIDReq) (*pb.GetLemmasByFilterIDResp, error) {
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

	lemmas, err := service.filterDataProvider.GetLemmasByFilterID(ctx, req.FilterID)
	if err != nil {
		service.logger.Error("could not get lemmas: %v", err)
		return nil, status.Errorf(codes.Internal, "could not get lemmas: %v", err)
	}
	return &pb.GetLemmasByFilterIDResp{Lemmas: lemmas}, nil
}

func (service *FilterService) GetKeywordsByLemmas(ctx context.Context, req *pb.GetKeywordsByLemmasReq) (*pb.GetKeywordsByLemmasResp, error) {
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

	if len(req.LemmasIDs) == 0 {
		service.logger.Error("lemmasIDs are required")
		return nil, status.Error(codes.InvalidArgument, "lemmasIDs are required")
	}

	// Fetch keywords for the provided lemmas
	keywordsResp, err := service.filterDataProvider.GetKeywordsByLemmas(ctx, req)
	if err != nil {
		service.logger.Error("could not get keywords by lemmas: %v", err)
		return nil, status.Errorf(codes.Internal, "could not get keywords by lemmas: %v", err)
	}

	return keywordsResp, nil
}
