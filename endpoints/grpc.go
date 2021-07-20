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

package endpoints

import (
	"reflect"

	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

type grpcHandler struct {
	Register interface{}
	Server   interface{}
}

type GRPCRegistry []*grpcHandler

func (r *GRPCRegistry) Register(fn, server interface{}) {
	if *r == nil {
		*r = make([]*grpcHandler, 0)
	}
	*r = append(*r, &grpcHandler{
		Register: fn,
		Server:   server,
	})
}

func (r GRPCRegistry) ApplyTo(srv *grpc.Server) error {
	for _, ep := range r {
		klog.Infof("Registering gRPC server: %s", reflect.TypeOf(ep.Server))

		fn := reflect.ValueOf(ep.Register)
		params := []reflect.Value{
			reflect.ValueOf(srv),
			reflect.ValueOf(ep.Server),
		}
		fn.Call(params)
	}
	return nil
}
