package controllers

import (
	"chat-app-portfolio/config"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateNewChatRoom(c echo.Context) error {
	ctx := context.Background()
	redis := config.GetRedisClient()
	res, _ := redis.Set(ctx, "input_counter", uuid.New(), 0).Result()

	return c.JSON(200, res)
}

func FetchExistingChatRoom(c echo.Context) error {
	return c.String(200, "get")
}

func DeleteExistingChatRoom(c echo.Context) error {
	return c.String(200, "delete")
}
