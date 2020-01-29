package main

import (
	"testing"
	"time"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/store/proto"
)

func TestApplyInt64Mutation(t *testing.T) {
	s := InitTestServer()

	tstart := time.Now().Unix()
	adj := &rcpb.ReleaseMetadata{DateAdded: tstart}

	ret, err := s.applyMutation(adj, &pb.Mutation{Field: int32(0), Value: &pb.Mutation_Int64Value{time.Now().Add(time.Minute).Unix()}})

	if err != nil {
		t.Errorf("Mutation was not applied")
	}

	adj2, ok := ret.(*rcpb.ReleaseMetadata)
	if !ok {
		t.Fatalf("Unable to channel %v", adj2)
	}

	if adj2.GetDateAdded() == tstart {
		t.Errorf("Mutation error %v -> %v (%v)", ret, adj2, tstart)
	}
}

func TestApplyStringMutation(t *testing.T) {
	s := InitTestServer()

	adj := &rcpb.ReleaseMetadata{FilePath: "first"}

	ret, err := s.applyMutation(adj, &pb.Mutation{Field: int32(2), Value: &pb.Mutation_StringValue{"second"}})

	if err != nil {
		t.Errorf("Mutation was not applied")
	}

	adj2, ok := ret.(*rcpb.ReleaseMetadata)
	if !ok {
		t.Fatalf("Unable to channel %v", adj2)
	}

	if adj2.FilePath == "first" {
		t.Errorf("Mutation error %v -> %v", ret, adj2)
	}
}
