/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"gomodules.xyz/grpc-go-addons/cors"
	"gomodules.xyz/grpc-go-addons/endpoints"

	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Config struct {
	SecureAddr         string
	PlaintextAddr      string
	APIDomain          string
	CACertFile         string
	CertFile           string
	KeyFile            string
	EnableCORS         bool
	CORSOriginHost     string
	CORSAllowSubdomain bool

	grpcRegistry  endpoints.GRPCRegistry
	proxyRegistry endpoints.ProxyRegistry
	corsRegistry  cors.PatternRegistry

	grpcOptions  []grpc.ServerOption
	gwMuxOptions []gwrt.ServeMuxOption
}

func NewConfig() *Config {
	return &Config{}
}

func (s *Config) UseTLS() bool {
	return !(s.CACertFile == "" && s.CertFile == "" && s.KeyFile == "")
}

func (s *Config) SetGRPCRegistry(reg endpoints.GRPCRegistry) {
	s.grpcRegistry = reg
}

func (s *Config) SetProxyRegistry(reg endpoints.ProxyRegistry) {
	s.proxyRegistry = reg
}

func (s *Config) SetCORSRegistry(reg cors.PatternRegistry) {
	s.corsRegistry = reg
}

func (s *Config) GRPCServerOption(opt ...grpc.ServerOption) {
	s.grpcOptions = opt
}

func (s *Config) GatewayMuxOption(opt ...gwrt.ServeMuxOption) {
	s.gwMuxOptions = opt
}

func (c Config) New() (*Server, error) {
	return &Server{Config: c}, nil
}
