// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: postgres.proto

/*
Package postgres is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package postgres

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_PostgresService_GetDatums_0(ctx context.Context, marshaler runtime.Marshaler, client PostgresServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetDatums(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PostgresService_GetDatums_0(ctx context.Context, marshaler runtime.Marshaler, server PostgresServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetDatums(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterPostgresServiceHandlerServer registers the http handlers for service PostgresService to "mux".
// UnaryRPC     :call PostgresServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterPostgresServiceHandlerFromEndpoint instead.
func RegisterPostgresServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PostgresServiceServer) error {

	mux.Handle("POST", pattern_PostgresService_GetDatums_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/postgres.PostgresService/GetDatums", runtime.WithHTTPPathPattern("/datums"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PostgresService_GetDatums_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PostgresService_GetDatums_0(ctx, mux, outboundMarshaler, w, req, response_PostgresService_GetDatums_0{resp}, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterPostgresServiceHandlerFromEndpoint is same as RegisterPostgresServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPostgresServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPostgresServiceHandler(ctx, mux, conn)
}

// RegisterPostgresServiceHandler registers the http handlers for service PostgresService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPostgresServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPostgresServiceHandlerClient(ctx, mux, NewPostgresServiceClient(conn))
}

// RegisterPostgresServiceHandlerClient registers the http handlers for service PostgresService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PostgresServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PostgresServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PostgresServiceClient" to call the correct interceptors.
func RegisterPostgresServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PostgresServiceClient) error {

	mux.Handle("POST", pattern_PostgresService_GetDatums_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/postgres.PostgresService/GetDatums", runtime.WithHTTPPathPattern("/datums"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PostgresService_GetDatums_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PostgresService_GetDatums_0(ctx, mux, outboundMarshaler, w, req, response_PostgresService_GetDatums_0{resp}, mux.GetForwardResponseOptions()...)

	})

	return nil
}

type response_PostgresService_GetDatums_0 struct {
	proto.Message
}

func (m response_PostgresService_GetDatums_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*GetDatumsResponse)
	return response.Value
}

var (
	pattern_PostgresService_GetDatums_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"datums"}, ""))
)

var (
	forward_PostgresService_GetDatums_0 = runtime.ForwardResponseMessage
)