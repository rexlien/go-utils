module github.com/rexlien/go-utils/go-utils

go 1.12

require (
	github.com/Shopify/toxiproxy v2.1.4+incompatible
	github.com/aws/aws-sdk-go v1.35.8
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.33.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0
)

replace github.com/rexlien/go-utils/xln-utils => ./xln-utils
