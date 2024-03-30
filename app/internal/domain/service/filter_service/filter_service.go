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
