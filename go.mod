module github.com/rexlien/go-utils/go-utils

go 1.12

require (
	github.com/Shopify/toxiproxy v2.1.4+incompatible
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/rexlien/go-utils/xln-utils v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace github.com/rexlien/go-utils/xln-utils => ./xln-utils
