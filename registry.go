// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package main

import (
	"database/sql"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb_postgres "immut-api/api/postgres/v1"
	"immut-api/internal/server"
	app_postgres "immut-api/postgres"
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
