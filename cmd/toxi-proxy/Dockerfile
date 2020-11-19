FROM golang:1.12

WORKDIR /go/src/app
RUN go get github.com/Shopify/toxiproxy/cmd
COPY . .


RUN export GO111MODULE="on" && go mod download
RUN chmod +x xln-proto/protoGen-go.sh
RUN export GO111MODULE="on" && go build cmd/toxi-proxy/toxi-proxy.go

CMD ["toxi-proxy"]