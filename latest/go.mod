module github.com/rawahars/ttrpc-stress/latest

go 1.23.3

replace (
	github.com/containerd/ttrpc v0.0.0 => github.com/containerd/ttrpc v1.2.7
	github.com/ttrpc-stress/payload => ../payload
)

require (
	github.com/containerd/ttrpc v0.0.0
	github.com/ttrpc-stress/payload v0.0.0
)

require (
	github.com/containerd/log v0.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230731190214-cbb8c96f2d6d // indirect
	google.golang.org/grpc v1.57.1 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
