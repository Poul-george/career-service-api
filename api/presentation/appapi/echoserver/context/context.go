package context

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
}

func UserIDKey() string {
	return "user_id"
}
