package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Engine() *echo.Echo {
	e := echo.New()

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		ExposeHeaders:    []string{echo.HeaderContentLength},
		AllowCredentials: true,
	}

	e.Use(middleware.CORSWithConfig(corsConfig))
	e.Use(middleware.Logger())

	return e
}
