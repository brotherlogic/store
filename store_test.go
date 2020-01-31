package main

import (
	"context"
	"testing"
	"time"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/store/proto"
	"github.com/golang/protobuf/ptypes"
)

func InitTestServer() *Server {
	s := Init()
	s.savePath = "./.testing/"
	return s
}

func TestBasicWriteRead(t *testing.T) {
	s := InitTestServer()

	t1 := &rcpb.ReleaseMetadata{DateAdded: time.Now().Unix()}
	an, _ := ptypes.MarshalAny(t1)
	_, err := s.Write(context.Background(), &pb.WriteRequest{Key: &pb.Key{Path: "ftest"}, Init: an})
	if err != nil {
		t.Fatalf("Error in write")
	}

	_, err = s.Write(context.Background(), &pb.WriteRequest{Key: &pb.Key{Path: "ftest", Version: 1}, Mutations: []*pb.Mutation{&pb.Mutation{Field: 0, Value: &pb.Mutation_Int64Value{time.Now().Add(time.Minute).Unix()}}}})
	if err != nil {
		t.Fatalf("Mutation failed: %v", err)
	}
}
