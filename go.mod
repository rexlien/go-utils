module github.com/rexlien/go-utils/go-utils

go 1.12

require (
	github.com/Shopify/toxiproxy v2.1.4+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rexlien/go-utils/xln-utils v0.0.0-00010101000000-000000000000
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace github.com/rexlien/go-utils/xln-utils => ./xln-utils
