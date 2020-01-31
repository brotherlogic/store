package main

import (
	"context"
	"testing"
	"time"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/store/proto"
	"github.com/golang/protobuf/ptypes"
)

func TestWrite(t *testing.T) {
	s := InitTestServer()

	_, err := s.Write(context.Background(), &pb.WriteRequest{Key: &pb.Key{Path: "testing"}})

	if err != nil {
		t.Errorf("Bad write: %v", err)
	}
}

func TestLoadSave(t *testing.T) {
	s := InitTestServer()

	t1 := &rcpb.ReleaseMetadata{DateAdded: time.Now().Unix()}
	err := s.save("testing/t1", t1)
	if err != nil {
		t.Fatalf("Bad save: %v", err)
	}

	val, err := s.load("testing/t1")
	if err != nil {
		t.Fatalf("Bad load: %v", err)
	}

	c, ok := val.(*rcpb.ReleaseMetadata)
	if !ok {
		t.Fatalf("Bad load: %v", c)
	}

	if c.GetDateAdded() != t1.GetDateAdded() {
		t.Errorf("Bad load %v -> %v", t1, c)
	}
}

func TestBadLoad(t *testing.T) {
	s := InitTestServer()

	val, err := s.load("madeup")
	if err == nil {
		t.Errorf("Bad load was successful: %v", val)
	}
}

func TestBadInit(t *testing.T) {
	s := InitTestServer()

	_, err := s.initialWrite(context.Background(), &pb.WriteRequest{Key: &pb.Key{}})

	if err == nil {
		t.Errorf("Bad filename did not fail")
	}
}

func TestBadSave(t *testing.T) {
	s := InitTestServer()

	err := s.saveAny("", nil)
	if err == nil {
		t.Errorf("Should have failed")
	}
}

func TestWriteWithBadRead(t *testing.T) {
	s := InitTestServer()

	_, err := s.Write(context.Background(), &pb.WriteRequest{Key: &pb.Key{Path: "ftest_blah", Version: 1}, Mutations: []*pb.Mutation{&pb.Mutation{Field: 0, Value: &pb.Mutation_Int64Value{time.Now().Add(time.Minute).Unix()}}}})
	if err == nil {
		t.Fatalf("Mutation failed: %v", err)
	}
}

func TestWriteWithBadMutation(t *testing.T) {
	s := InitTestServer()

	t1 := &rcpb.ReleaseMetadata{DateAdded: time.Now().Unix()}
	an, _ := ptypes.MarshalAny(t1)
	_, err := s.Write(context.Background(), &pb.WriteRequest{Key: &pb.Key{Path: "ftest"}, Init: an})
	if err != nil {
		t.Fatalf("Error in write")
	}

	_, err = s.Write(context.Background(), &pb.WriteRequest{Key: &pb.Key{Path: "ftest", Version: 1}, Mutations: []*pb.Mutation{&pb.Mutation{Field: 999, Value: &pb.Mutation_Int64Value{time.Now().Add(time.Minute).Unix()}}}})
	if err == nil {
		t.Fatalf("Mutation did not failed: %v", err)
	}
}
