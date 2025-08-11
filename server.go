package main

import (
	// "fmt"

	"chat-app-portfolio/config"
	"chat-app-portfolio/helper"
	"chat-app-portfolio/servicethread"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	godotenv.Load()
	config.ConnectRedisClient()

	e := echo.New()

	helper.MQTTConnect()
	go servicethread.InitMqttReconnectThread()

	e.Use(middleware.CORSWithConfig(config.GenerateCorsConfig()))

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(30))))

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")))
}
