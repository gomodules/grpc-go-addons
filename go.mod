module gomodules.xyz/grpc-go-addons

go 1.12

require (
	github.com/appscode/go v0.0.0-20190523031839-1468ee3a76e8
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.0
	github.com/pkg/errors v0.8.1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.3.0
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	gomodules.xyz/notify v0.0.0-20190424183923-af47cb5a07a4
	google.golang.org/grpc v1.19.0
)

replace github.com/grpc-ecosystem/grpc-gateway v1.3.1 => github.com/appscode/grpc-gateway v1.3.1-ac
