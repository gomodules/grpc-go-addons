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
	gort "runtime"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

type RegisterProxyHandlerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type proxyHandler struct {
	Register RegisterProxyHandlerFunc
}

type ProxyRegistry []*proxyHandler

func (r *ProxyRegistry) Register(fn RegisterProxyHandlerFunc) {
	if *r == nil {
		*r = make([]*proxyHandler, 0)
	}
	*r = append(*r, &proxyHandler{
		Register: fn,
	})
}

func (r ProxyRegistry) ApplyTo(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	for _, ep := range r {
		klog.Infof("Registering grpc-gateway endpoint: %s", funcName(ep.Register))
		if err := ep.Register(context.Background(), mux, endpoint, opts); err != nil {
			return nil
		}
	}
	return nil
}

func funcName(i interface{}) string {
	name := gort.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
