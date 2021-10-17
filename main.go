package main

import (
	"net/http"
	"time"

	"github.com/andrefebrianto/URL-Shortener-Service/database/cassandra"
	"github.com/andrefebrianto/URL-Shortener-Service/domain/ShortLink/httpcontroller"
	"github.com/andrefebrianto/URL-Shortener-Service/domain/ShortLink/repository/command"
	"github.com/andrefebrianto/URL-Shortener-Service/domain/ShortLink/repository/query"
	"github.com/andrefebrianto/URL-Shortener-Service/domain/ShortLink/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

var GlobalConfig = viper.New()

func setCorsHeader(nextHandler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return nextHandler(context)
	}
}

func responsePing(requestContext echo.Context) error {
	return requestContext.String(http.StatusOK, "Service is running properly")
}

func init() {
	GlobalConfig.SetConfigFile(`configs/configs.json`)
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize database
	cassandra.SetupConnection()

	// Echo instance
	httpServer := echo.New()

	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(setCorsHeader)

	// Routes
	httpServer.GET("/", responsePing)

	shortLinkCommandRepo := command.CreateCassandraCommandRepository(cassandra.GetConnection())
	shortLinkQueryRepo := query.CreateCassandraQueryRepository(cassandra.GetConnection())
	shortLinkUseCase := usecase.CreateShortLinkUseCase(shortLinkCommandRepo, shortLinkQueryRepo, time.Duration(GlobalConfig.GetInt("context.timeout"))*time.Second)

	httpcontroller.CreateShortLinkHttpController(httpServer, shortLinkUseCase)

	// Start server
	httpServer.Logger.Fatal(httpServer.Start(GlobalConfig.GetString("server.port")))
}
