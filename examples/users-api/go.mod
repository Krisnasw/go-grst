module github.com/krisnasw/go-grst/examples/users-api

go 1.15

replace github.com/krisnasw/go-grst/grst => ../../grst

replace github.com/krisnasw/go-grst => ../../

replace github.com/krisnasw/go-grst/protoc-gen-cdd => ../../protoc-gen-cdd

require (
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/krisnasw/go-grst/grst v0.0.0-00010101000000-000000000000
	github.com/krisnasw/go-grst/protoc-gen-cdd v0.0.0-00010101000000-000000000000
	github.com/mcuadros/go-defaults v1.2.0
	github.com/sirupsen/logrus v1.8.0
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/validator.v2 v2.0.0-20200605151824-2b28d334fa05
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)
