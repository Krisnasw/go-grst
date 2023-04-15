module github.com/herryg91/cdd/examples/users-api

go 1.15

replace github.com/herryg91/cdd/grst => ../../grst

replace github.com/herryg91/cdd => ../../

replace github.com/herryg91/cdd/protoc-gen-cdd => ../../protoc-gen-cdd

require (
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/herryg91/cdd/grst v0.0.0-00010101000000-000000000000
	github.com/herryg91/cdd/protoc-gen-cdd v0.0.0-00010101000000-000000000000
	github.com/herryg91/hgolib/databases v0.0.0-20201227172554-ac2bb27a5077
	github.com/jinzhu/gorm v1.9.16
	github.com/kelseyhightower/envconfig v1.4.0
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
