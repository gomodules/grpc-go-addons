package options

import (
	"flag"
	_env "github.com/appscode/go/env"
	"github.com/appscode/grpc-go-addons/server"
	"github.com/spf13/pflag"
)

type Options struct {
	SecureAddr    string
	PlaintextAddr string
	APIDomain     string
	CACertFile    string
	CertFile      string
	KeyFile       string

	EnableCORS         bool
	CORSOriginHost     string
	CORSAllowSubdomain bool
}

func New() *Options {
	return &Options{
		SecureAddr:    ":8443",
		PlaintextAddr: ":8080",
	}
}

func (s *Options) AddGoFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.SecureAddr, "secure-addr", s.SecureAddr, "host:port used to serve secure apis")
	fs.StringVar(&s.PlaintextAddr, "plaintext-addr", s.PlaintextAddr, "host:port used to serve http json apis")

	fs.StringVar(&s.APIDomain, "api-domain", s.APIDomain, "Domain used for apiserver (prod: api.appscode.com")
	fs.StringVar(&s.CACertFile, "tls-ca-file", s.CACertFile, "File containing CA certificate")
	fs.StringVar(&s.CertFile, "tls-cert-file", s.CertFile, "File container server TLS certificate")
	fs.StringVar(&s.KeyFile, "tls-private-key-file", s.KeyFile, "File containing server TLS private key")

	fs.BoolVar(&s.EnableCORS, "enable-cors", s.EnableCORS, "Enable CORS support")
	fs.StringVar(&s.CORSOriginHost, "cors-origin-host", s.CORSOriginHost, `Allowed CORS origin host e.g, domain[:port]`)
	fs.BoolVar(&s.CORSAllowSubdomain, "cors-origin-allow-subdomain", s.CORSAllowSubdomain, "Allow CORS request from subdomains of origin")
}

func (s *Options) AddFlags(fs *pflag.FlagSet) {
	gfs := flag.NewFlagSet("grpc", flag.ExitOnError)
	s.AddGoFlags(gfs)
	fs.AddGoFlagSet(gfs)
}

func (s *Options) ApplyTo(cfg *server.Config) error {
	cfg.SecureAddr = s.SecureAddr
	cfg.PlaintextAddr = s.PlaintextAddr
	cfg.APIDomain = s.APIDomain
	cfg.CACertFile = s.CACertFile
	cfg.CertFile = s.CertFile
	cfg.KeyFile = s.KeyFile
	cfg.EnableCORS = s.EnableCORS
	cfg.CORSOriginHost = s.CORSOriginHost
	cfg.CORSAllowSubdomain = s.CORSAllowSubdomain

	return nil
}
