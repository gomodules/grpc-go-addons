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
