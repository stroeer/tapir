module github.com/stroeer/tapir/go

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/stroeer/go-tapir v0.0.0-20210901143243-2da66482ca62
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// replace go-tapir dependency with generated module
replace github.com/stroeer/go-tapir => ../gen/go
