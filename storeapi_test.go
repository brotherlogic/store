package main

import (
	"context"
	"testing"
	"time"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/store/proto"
)

func TestWrite(t *testing.T) {
	s := InitTestServer()

	_, err := s.Write(context.Background(), &pb.WriteRequest{})

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
