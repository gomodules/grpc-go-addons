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

package cors

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_env "gomodules.xyz/x/env"
	"k8s.io/klog/v2"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
// https://fetch.spec.whatwg.org/#cors-protocol-and-credentials
// For requests without credentials, the server may specify "*" as a wildcard, thereby allowing any origin to access the resource.
type Handler struct {
	options *options
	reg     PatternRegistry
}

func NewHandler(r PatternRegistry, opts ...Option) *Handler {
	return &Handler{reg: r, options: evaluateOptions(opts)}
}

func (r *Handler) RegisterHandler(mux *runtime.ServeMux) {
	for _, p := range r.reg {
		mux.Handle("OPTIONS", p, r.ServeHTTP)
	}
}

func (r Handler) ServeHTTP(w http.ResponseWriter, req *http.Request, _ map[string]string) {
	headers := map[string]string{
		"access-control-allow-methods": "POST,GET,OPTIONS,PUT,DELETE",
		"access-control-allow-headers": req.Header.Get("access-control-request-headers"),
	}
	if r.options.allowHost == "*" {
		headers["access-control-allow-origin"] = "*"
	} else if r.options.allowHost != "" {
		origin := req.Header.Get("Origin")

		u, err := url.Parse(origin)
		if err != nil {
			klog.Errorln("Failed to parse CORS origin header", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ok := u.Host == r.options.allowHost ||
			(r.options.allowSubdomain && strings.HasSuffix(u.Host, "."+r.options.allowHost))
		if !ok {

			klog.Errorf("CORS request from prohibited domain %v", origin)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !_env.FromHost().DevMode() {
			u.Scheme = "https"
		}
		headers["access-control-allow-origin"] = u.String()
		headers["access-control-allow-credentials"] = "true"
		headers["vary"] = "Origin"
	}
	for k, v := range headers {
		w.Header().Set(k, v)
	}
}
