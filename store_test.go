package main

func InitTestServer() *Server {
	s := Init()
	s.savePath = "./.testing/"
	return s
}
