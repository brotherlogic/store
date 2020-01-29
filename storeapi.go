package main

import "golang.org/x/net/context"

import pb "github.com/brotherlogic/store/proto"

func (s *Server) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	return &pb.WriteResponse{}, nil
}
