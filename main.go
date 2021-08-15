package main

import (
	"fmt"
	"net/http"
	"time"

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

	if GlobalConfig.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	contextTimeout := time.Duration(GlobalConfig.GetInt("context.timeout")) * time.Second

	// Echo instance
	httpServer := echo.New()

	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(setCorsHeader)

	// Routes
	httpServer.GET("/", responsePing)

	fmt.Println(contextTimeout)
	// Start server
	httpServer.Logger.Fatal(httpServer.Start(GlobalConfig.GetString("server.port")))
}
