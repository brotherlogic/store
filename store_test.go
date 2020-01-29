package main

import "testing"

func InitTestServer() *Server {
	s := Init()
	return s
}

func TestBasicSaveLoad(t *testing.T) {
	s := InitTestServer()
	s.doNothing()
	s.reallyDoNothing()
}
