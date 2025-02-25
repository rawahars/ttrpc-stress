module github.com/rawahars/ttrpc-stress/v1_0_2

go 1.23.3

replace github.com/ttrpc-stress/payload => ../payload

require (
	github.com/containerd/ttrpc v1.0.2
	github.com/ttrpc-stress/payload v0.0.0-00010101000000-000000000000
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	google.golang.org/genproto v0.0.0-20200117163144-32f20d992d24 // indirect
	google.golang.org/grpc v1.26.0 // indirect
)
