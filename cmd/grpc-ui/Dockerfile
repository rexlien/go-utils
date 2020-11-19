FROM golang:1.12

WORKDIR /go/src/app
RUN go get github.com/fullstorydev/grpcui/...
RUN go install github.com/fullstorydev/grpcui/cmd/grpcui

ENTRYPOINT grpcui -plaintext $XLN_GRPC_HOST