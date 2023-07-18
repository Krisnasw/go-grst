package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"service_name" default:"location-api"`
	Environment string `envconfig:"environment" default:"dev"`
	Maintenance bool   `envconfig:"maintenance" default:"false"`
	RestPort    int    `envconfig:"rest_port" default:"18080" required:"true"`
	GrpcPort    int    `envconfig:"grpc_port" default:"19090" required:"true"`

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         int    `envconfig:"DB_PORT" default:"3306"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:""`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"cdd"`
	DBLogEnable    bool   `envconfig:"DB_LOG_ENABLE" default:"true"`
	DBLogLevel     int    `envconfig:"DB_LOG_LEVEL" default:"3"`
	DBLogThreshold int    `envconfig:"DB_LOG_THRESHOLD" default:"200"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
