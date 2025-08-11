package controllers

import "github.com/labstack/echo/v4"

func CreateNewChatRoom(c echo.Context) error {
	return c.String(200, "post")
}

func FetchExistingChatRoom(c echo.Context) error {
	return c.String(200, "get")
}

func DeleteExistingChatRoom(c echo.Context) error {
	return c.String(200, "delete")
}
