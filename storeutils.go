package main

import (
	"reflect"

	"github.com/golang/protobuf/proto"

	pb "github.com/brotherlogic/store/proto"
)

func (s *Server) applyMutation(base proto.Message, mutation *pb.Mutation) (proto.Message, error) {
	val := reflect.ValueOf(base).Elem()

	switch x := mutation.Value.(type) {
	case *pb.Mutation_StringValue:
		val.Field(int(mutation.GetField())).SetString(x.StringValue)
	case *pb.Mutation_Int64Value:
		val.Field(int(mutation.GetField())).SetInt(x.Int64Value)
	}

	return base, nil
}
