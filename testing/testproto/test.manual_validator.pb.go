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

// Manual code for validation tests.

package mwitkow_testproto

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

func (p *PingRequest) Validate() (*gojsonschema.Result, error) {
	if p.SleepTimeMs > 10000 {
		return nil, errors.New("cannot sleep for more than 10s")
	}
	return &gojsonschema.Result{}, nil
}
