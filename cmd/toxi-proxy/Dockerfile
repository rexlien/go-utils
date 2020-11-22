FROM golang:1.12

WORKDIR /go/src/app
RUN apt-get update && apt install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get github.com/Shopify/toxiproxy/cmd

COPY go.mod go.mod
RUN export GO111MODULE="on" && go mod download
COPY . .
RUN chmod +x xln-proto/protoGen-go.sh && sh xln-proto/protoGen-go.sh
RUN export GO111MODULE="on" && go build cmd/toxi-proxy/toxi-proxy.go

CMD ["toxi-proxy"]