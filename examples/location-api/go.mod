module github.com/krisnasw/go-grst/examples/location-api

go 1.15

// replace github.com/krisnasw/go-grst/grst => ../../grst

// replace github.com/krisnasw/go-grst => ../../

// replace github.com/krisnasw/go-grst/protoc-gen-cdd => ../../protoc-gen-cdd

require (
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/krisnasw/go-grst/grst v0.0.0-20230416062753-8ba2ab080f14
	github.com/krisnasw/go-grst/protoc-gen-cdd v0.0.0-20230416062753-8ba2ab080f14
	github.com/mcuadros/go-defaults v1.2.0
	github.com/sirupsen/logrus v1.8.0
	google.golang.org/genproto v0.0.0-20210310155132-4ce2db91004e
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.1-0.20201208041424-160c7477e0e8
	gopkg.in/validator.v2 v2.0.0-20200605151824-2b28d334fa05
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)
