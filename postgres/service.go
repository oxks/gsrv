// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package postgres

import (
	"context"

	"go.uber.org/zap"

	pb "immut-api/api/postgres/v1"
)

type Service struct {
	pb.UnimplementedPostgresServiceServer
	logger  *zap.Logger
	querier *Queries
}

func (s *Service) AllDatum(ctx context.Context, req *pb.AllDatumRequest) (*pb.AllDatumResponse, error) {

	result, err := s.querier.AllDatum(ctx)
	if err != nil {
		s.logger.Error("AllDatum sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.AllDatumResponse)
	for _, r := range result {
		res.List = append(res.List, toAllDatumRow(r))
	}
	return res, nil
}
