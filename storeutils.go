package main

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"

	pb "github.com/brotherlogic/store/proto"
)

func (s *Server) applyMutation(base proto.Message, mutation *pb.Mutation) (proto.Message, error) {
	val := reflect.ValueOf(base).Elem()

	if val.NumField() < int(mutation.GetField()) {
		return base, fmt.Errorf("Bad field number: %v", mutation.GetField())
	}

	switch x := mutation.GetValue().(type) {
	case *pb.Mutation_StringValue:
		val.Field(int(mutation.GetField())).SetString(x.StringValue)
	case *pb.Mutation_Int64Value:
		val.Field(int(mutation.GetField())).SetInt(x.Int64Value)
	}

	return base, nil
}

func buildMutations(p1, p2 proto.Message) []*pb.Mutation {
	p1e := reflect.ValueOf(p1).Elem()
	p2e := reflect.ValueOf(p2).Elem()

	return buildStruct(p1e, p2e)
}

func buildStruct(p1, p2 reflect.Value) []*pb.Mutation {
	mutations := []*pb.Mutation{}
	for i := 0; i < p1.NumField(); i++ {
		nmut := buildFieldMutations(i, p1.Field(i), p2.Field(i))
		mutations = append(mutations, nmut...)
	}
	return mutations
}

func buildFieldMutations(fieldNum int, p1, p2 reflect.Value) []*pb.Mutation {
	mutations := []*pb.Mutation{}

	switch p1.Kind() {
	case reflect.Int64:
		if p1.Int() != p2.Int() {
			mutations = append(mutations, &pb.Mutation{Field: int32(fieldNum), Value: &pb.Mutation_Int64Value{int64(p2.Int())}})
		}
	}

	return mutations
}
