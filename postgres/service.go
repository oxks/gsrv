package postgres

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "sqlcgrpc/proto/postgres"
)

type Service struct {
	pb.UnimplementedPostgresServiceServer
	logger *zap.Logger
	db     *Queries
}

func NewService(logger *zap.Logger, db *Queries) *Service {
	return &Service{logger: logger, db: db}
}

func (s *Service) GetDatums(ctx context.Context, in *emptypb.Empty) (out *pb.GetDatumsResponse, err error) {

	result, err := s.db.GetDatums(ctx)
	if err != nil {
		s.logger.Error("GetDatums sql call failed", zap.Error(err))
		return
	}
	out = new(pb.GetDatumsResponse)
	for _, r := range result {
		var item *pb.GetDatumsRow
		item, err = toGetDatumsRow(r)
		if err != nil {
			return
		}
		out.Value = append(out.Value, item)
	}
	return

}
