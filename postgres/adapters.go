package postgres

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "sqlcgrpc/proto/postgres"
)

func toGetDatumsRow(in GetDatumsRow) (out *pb.GetDatumsRow, err error) {
	out = new(pb.GetDatumsRow)
	out.Datum = in.Datum
	if in.Hash.Valid {
		out.Hash = wrapperspb.String(in.Hash.String)
	}
	if in.PreviousHash.Valid {
		out.PreviousHash = wrapperspb.String(in.PreviousHash.String)
	}
	out.CreatedAt = timestamppb.New(in.CreatedAt)
	out.Nickname = in.Nickname
	return
}
