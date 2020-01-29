package main

import (
	"fmt"
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
	return s.saveAny(path, an)
}

func (s *Server) saveAny(path string, an *any.Any) error {
	data, _ := proto.Marshal(an)

	if len(path) < 3 {
		return fmt.Errorf("No filename specified")
	}

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

func (s *Server) initialWrite(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	savepath := fmt.Sprintf("%v.%v", req.GetKey().GetPath(), 1)
	err := s.saveAny(savepath, req.GetInit())
	if err != nil {
		return nil, err
	}
	return &pb.WriteResponse{Key: &pb.Key{Path: req.GetKey().GetPath(), Version: 1}}, nil
}

func (s *Server) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	loadpath := fmt.Sprintf("%v.%v", req.GetKey().GetPath(), req.GetKey().GetVersion())

	if req.GetKey().GetVersion() == 0 {
		return s.initialWrite(ctx, req)
	}

	cdata, err := s.load(loadpath)
	if err != nil {
		return nil, err
	}

	for _, mutation := range req.GetMutations() {
		cdata, err = s.applyMutation(cdata, mutation)
		if err != nil {
			return nil, err
		}
	}

	savepath := fmt.Sprintf("%v.%v", req.GetKey().GetPath(), req.GetKey().GetVersion()+1)

	err = s.save(savepath, cdata)
	return &pb.WriteResponse{Key: &pb.Key{Path: req.GetKey().GetPath(), Version: req.GetKey().GetVersion() + 1}}, err
}
