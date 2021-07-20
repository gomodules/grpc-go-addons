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

type CorsOptions struct {
	Enable         bool
	OriginHost     string
	AllowSubdomain bool
}

func NewCORSOptions() *CorsOptions {
	return &CorsOptions{
		OriginHost:     "*",
		AllowSubdomain: true,
	}
}

func (o *CorsOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.BoolVar(&o.Enable, "enable-cors", o.Enable, "Enable CORS support")
	fs.StringVar(&o.OriginHost, "cors-origin-host", o.OriginHost, `Allowed CORS origin host e.g, domain[:port]`)
	fs.BoolVar(&o.AllowSubdomain, "cors-origin-allow-subdomain", o.AllowSubdomain, "Allow CORS request from subdomains of origin")
}

func (o *CorsOptions) AddFlags(fs *pflag.FlagSet) {
	gfs := flag.NewFlagSet("grpc-cors", flag.ExitOnError)
	o.AddGoFlags(gfs)
	fs.AddGoFlagSet(gfs)
}

func (o *CorsOptions) ApplyTo(cfg *server.Config) error {
	cfg.EnableCORS = o.Enable
	cfg.CORSOriginHost = o.OriginHost
	cfg.CORSAllowSubdomain = o.AllowSubdomain

	return nil
}

func (o *CorsOptions) Validate() []error {
	return nil
}
