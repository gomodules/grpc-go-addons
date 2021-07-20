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

type SecureServingOptions struct {
	SecureAddr    string
	PlaintextAddr string
	APIDomain     string
	CACertFile    string
	CertFile      string
	KeyFile       string
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		SecureAddr:    ":8443",
		PlaintextAddr: ":8080",
	}
}

func (o *SecureServingOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.SecureAddr, "secure-addr", o.SecureAddr, "host:port used to serve secure apis")
	fs.StringVar(&o.PlaintextAddr, "plaintext-addr", o.PlaintextAddr, "host:port used to serve http json apis")

	fs.StringVar(&o.APIDomain, "api-domain", o.APIDomain, "Domain used for apiserver (prod: api.appscode.com")
	fs.StringVar(&o.CACertFile, "tls-ca-file", o.CACertFile, "File containing CA certificate")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "File container server TLS certificate")
	fs.StringVar(&o.KeyFile, "tls-private-key-file", o.KeyFile, "File containing server TLS private key")
}

func (o *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	gfs := flag.NewFlagSet("grpc-serving", flag.ExitOnError)
	o.AddGoFlags(gfs)
	fs.AddGoFlagSet(gfs)
}

func (o *SecureServingOptions) ApplyTo(cfg *server.Config) error {
	cfg.SecureAddr = o.SecureAddr
	cfg.PlaintextAddr = o.PlaintextAddr
	cfg.APIDomain = o.APIDomain
	cfg.CACertFile = o.CACertFile
	cfg.CertFile = o.CertFile
	cfg.KeyFile = o.KeyFile

	return nil
}

func (o *SecureServingOptions) Validate() []error {
	return nil
}
