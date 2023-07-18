package main

import (
	"time"

	crud_city_usecase "github.com/krisnasw/go-grst/examples/location-api/app/usecase/crud_city"
	crud_province_usecase "github.com/krisnasw/go-grst/examples/location-api/app/usecase/crud_province"
	search_usecase "github.com/krisnasw/go-grst/examples/location-api/app/usecase/search"
	"github.com/krisnasw/go-grst/examples/location-api/config"
	"github.com/krisnasw/go-grst/examples/location-api/handler"
	pbCity "github.com/krisnasw/go-grst/examples/location-api/handler/grst/city"
	pbProvince "github.com/krisnasw/go-grst/examples/location-api/handler/grst/province"
	"github.com/krisnasw/go-grst/examples/location-api/repository/city_mysql"
	"github.com/krisnasw/go-grst/examples/location-api/repository/province_mysql"
	"github.com/krisnasw/go-grst/grst"
	loggerInterceptor "github.com/krisnasw/go-grst/grst/interceptor/logger"
	recoveryInterceptor "github.com/krisnasw/go-grst/grst/interceptor/recovery"
	sessionInterceptor "github.com/krisnasw/go-grst/grst/interceptor/session"
	"gorm.io/gorm/logger"

	"github.com/krisnasw/go-grst/examples/location-api/pkg/conn/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName,
		mysql.SetPrintLog(cfg.DBLogEnable, logger.LogLevel(cfg.DBLogLevel), time.Duration(cfg.DBLogThreshold)*time.Millisecond))
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	provinceRepo := province_mysql.New(db)
	cityRepo := city_mysql.New(db)

	provinceUsecase := crud_province_usecase.New(provinceRepo)
	cityUsecase := crud_city_usecase.New(cityRepo)
	citySearchUsecase := search_usecase.New(cityRepo, provinceRepo)

	provinceHndl := handler.NewProvinceHandler(provinceUsecase)
	cityHndl := handler.NewCityHandler(cityUsecase, citySearchUsecase)

	grpcRestSrv, err := grst.NewServer(cfg.GrpcPort, cfg.RestPort, true,
		grst.RegisterGRPCUnaryInterceptor("session", sessionInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("recovery", recoveryInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("log", loggerInterceptor.UnaryServerInterceptor()),
		grst.AddForwardHeaderToContext([]string{"country"}),
	)

	if err != nil {
		logrus.Panicln("Failed to Initialize GRPC-REST Server:", err)
	}

	pbProvince.RegisterProvinceGrstServer(grpcRestSrv, provinceHndl)
	pbCity.RegisterCityGrstServer(grpcRestSrv, cityHndl)
	if err := <-grpcRestSrv.ListenAndServeGrst(); err != nil {
		logrus.Panicln("Failed to Run Grpcrest Server:", err)
	}
}
