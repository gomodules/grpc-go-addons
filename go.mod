module gomodules.xyz/grpc-go-addons

go 1.12

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/pkg/errors v0.9.1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/net v0.0.0-20191002035440-2ec189313ef0
	gomodules.xyz/notify v0.1.1
	gomodules.xyz/x v0.0.0-20201105065653-91c568df6331
	google.golang.org/grpc v1.24.0
)

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2
