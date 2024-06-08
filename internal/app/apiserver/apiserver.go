package apiserver

import (
	"fmt"
)


type ApiServer struct {
	config *Config
}

func NewApiServer(config *Config) *ApiServer {
    return &ApiServer{
		config: config
	}
}

func (s *ApiServer) Run() error {
    fmt.Println("Assalam Alaykum")
	return nil
}