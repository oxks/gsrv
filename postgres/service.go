// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	pb "immut-apt/api/postgres/v1"
	"immut-apt/internal/validation"
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

func (s *Service) DatumAdd(ctx context.Context, req *pb.DatumAddRequest) (*pb.DatumAddResponse, error) {
	var arg DatumAddParams
	arg.AuthorID = req.GetAuthorId()
	arg.Datum = req.GetDatum()
	if v := req.GetHash(); v != nil {
		arg.Hash = sql.NullString{Valid: true, String: v.Value}
	}
	if v := req.GetCreatedAt(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid CreatedAt: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.CreatedAt = v.AsTime()
	} else {
		err := fmt.Errorf("field CreatedAt is required%w", validation.ErrUserInput)
		return nil, err
	}
	if v := req.GetPreviousHash(); v != nil {
		arg.PreviousHash = sql.NullString{Valid: true, String: v.Value}
	}

	err := s.querier.DatumAdd(ctx, arg)
	if err != nil {
		s.logger.Error("DatumAdd sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DatumAddResponse{}, nil
}

func (s *Service) GetDatums(ctx context.Context, req *pb.GetDatumsRequest) (*pb.GetDatumsResponse, error) {

	result, err := s.querier.GetDatums(ctx)
	if err != nil {
		s.logger.Error("GetDatums sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.GetDatumsResponse)
	for _, r := range result {
		res.List = append(res.List, toGetDatumsRow(r))
	}
	return res, nil
}
