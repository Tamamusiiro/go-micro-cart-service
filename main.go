package main

import (
	"fmt"
	"log"

	"github.com/Tamamusiiro/go-micro-cart-service/domian/respository"
	"github.com/Tamamusiiro/go-micro-cart-service/domian/service"
	"github.com/Tamamusiiro/go-micro-cart-service/handler"
	pb "github.com/Tamamusiiro/go-micro-cart-service/proto/cart"
	common "github.com/Tamamusiiro/go-micro-common"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var QPS = 100

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1:8500", "/micro/config")
	if err != nil {
		log.Fatal(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	// 链路追踪
	tracer, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	// 数据库配置
	mysqlConfig := common.GetMysqlConfig(consulConfig, "mysql")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlConfig.User, mysqlConfig.PWD, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DB,
		),
	}))
	if err != nil {
		log.Fatal(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8084"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	srv.Init()

	repository := respository.NewCartRepository(db)
	err = repository.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
	cartService := service.NewCartService(repository)

	// Register handler
	err = pb.RegisterCartHandler(srv.Server(), &handler.CartHandler{CartService: cartService})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
