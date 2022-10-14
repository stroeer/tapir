module github.com/stroeer/tapir/go

go 1.18

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/stroeer/go-tapir v0.0.0-20210901143243-2da66482ca62
	google.golang.org/genproto v0.0.0-20221010155953-15ba04fc1c0e // indirect
	google.golang.org/grpc v1.50.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace go-tapir dependency with generated module
replace github.com/stroeer/go-tapir => ../gen/go
