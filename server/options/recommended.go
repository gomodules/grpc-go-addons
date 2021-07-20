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

package options

import (
	"flag"

	"gomodules.xyz/grpc-go-addons/server"

	"github.com/spf13/pflag"
)

type RecommendedOptions struct {
	Cors          *CorsOptions
	SecureServing *SecureServingOptions
}

func NewRecommendedOptions() *RecommendedOptions {
	return &RecommendedOptions{
		Cors:          NewCORSOptions(),
		SecureServing: NewSecureServingOptions(),
	}
}

func (o *RecommendedOptions) AddGoFlags(fs *flag.FlagSet) {
	o.Cors.AddGoFlags(fs)
	o.SecureServing.AddGoFlags(fs)
}

func (o *RecommendedOptions) AddFlags(fs *pflag.FlagSet) {
	o.Cors.AddFlags(fs)
	o.SecureServing.AddFlags(fs)
}

func (o *RecommendedOptions) ApplyTo(config *server.Config) error {
	if err := o.Cors.ApplyTo(config); err != nil {
		return err
	}
	if err := o.SecureServing.ApplyTo(config); err != nil {
		return err
	}
	return nil
}

func (o *RecommendedOptions) Validate() []error {
	var errors []error
	errors = append(errors, o.Cors.Validate()...)
	errors = append(errors, o.SecureServing.Validate()...)

	return errors
}
