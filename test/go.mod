module github.com/rawahars/ttrpc-stress/test

go 1.23.3

require (
	github.com/Microsoft/go-winio v0.6.2
	github.com/gogo/protobuf v1.3.2 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.36.5 // indirect
)

require (
	github.com/rawahars/ttrpc-stress/latest v0.0.0
	github.com/rawahars/ttrpc-stress/v1_0_2 v0.0.0
	github.com/rawahars/ttrpc-stress/v1_1_0 v0.0.0
	github.com/rawahars/ttrpc-stress/v1_2_0 v0.0.0
	github.com/rawahars/ttrpc-stress/v1_2_4 v0.0.0
)

require (
	github.com/containerd/ttrpc v1.2.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/ttrpc-stress/payload v0.0.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230731190214-cbb8c96f2d6d // indirect
	google.golang.org/grpc v1.57.1 // indirect
)

replace (
	github.com/containerd/ttrpc v0.0.0 => github.com/containerd/ttrpc v1.2.7
	github.com/rawahars/ttrpc-stress/latest => ../latest
	github.com/rawahars/ttrpc-stress/v1_0_2 => ../1.0.2
	github.com/rawahars/ttrpc-stress/v1_1_0 => ../1.1.0
	github.com/rawahars/ttrpc-stress/v1_2_0 => ../1.2.0
	github.com/rawahars/ttrpc-stress/v1_2_4 => ../1.2.4
	github.com/ttrpc-stress/payload => ../payload
)
