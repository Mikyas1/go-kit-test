// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "notificator/pkg/endpoint"
	pb "notificator/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	sendEmail grpc.Handler
	pb.UnimplementedNotificatorServer
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.NotificatorServer {
	return &grpcServer{sendEmail: makeSendEmailHandler(endpoints, options["SendEmail"])}
}
