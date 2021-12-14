package main

import (
	"database/sql"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"sqlcgrpc/internal/server"
	app_postgres "sqlcgrpc/postgres"
	pb_postgres "sqlcgrpc/proto/postgres"
)

func registerServer(logger *zap.Logger, db *sql.DB) server.RegisterServer {
	return func(grpcServer *grpc.Server) {
		pb_postgres.RegisterPostgresServiceServer(grpcServer, app_postgres.NewService(logger, app_postgres.New(db)))

	}
}

func registerHandlers() []server.RegisterHandler {
	var handlers []server.RegisterHandler

	handlers = append(handlers, pb_postgres.RegisterPostgresServiceHandler)

	return handlers
}
