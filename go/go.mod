module github.com/stroeer/tapir/go

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/stroeer/go-tapir v0.0.0-20210901143243-2da66482ca62
	google.golang.org/genproto v0.0.0-20210708141623-e76da96a951f // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
)

// replace go-tapir dependency with generated module
replace github.com/stroeer/go-tapir => ../gen/go
