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

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
// https://fetch.spec.whatwg.org/#cors-protocol-and-credentials
// For requests without credentials, the server may specify "*" as a wildcard, thereby allowing any origin to access the resource.
type PatternRegistry []runtime.Pattern

func (r *PatternRegistry) Register(f []runtime.Pattern) {
	if *r == nil {
		*r = make([]runtime.Pattern, 0)
	}
	*r = append(*r, f...)
}
