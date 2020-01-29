package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/store/proto"
)

func TestWrite(t *testing.T) {
	s := InitTestServer()

	_, err := s.Write(context.Background(), &pb.WriteRequest{})

	if err != nil {
		t.Errorf("Bad write: %v", err)
	}
}
