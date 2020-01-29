package main

import (
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/store/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func (s *Server) save(path string, m proto.Message) error {
	an, _ := ptypes.MarshalAny(m)
	data, _ := proto.Marshal(an)

	fullpath := s.savePath + path
	split := strings.Split(fullpath, "/")
	dirs := strings.Join(split[0:len(split)-1], "/")

	os.MkdirAll(dirs, 0777)

	return ioutil.WriteFile(fullpath, data, 0644)
}

func (s *Server) load(path string) (proto.Message, error) {
	fullpath := s.savePath + path

	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return nil, err
	}

	var result ptypes.DynamicAny
	var an any.Any
	proto.Unmarshal(data, &an)
	err = ptypes.UnmarshalAny(&an, &result)

	return result.Message, nil
}

func (s *Server) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	return &pb.WriteResponse{}, nil
}
