package ipad

import "github.com/labstack/echo/v4"

func Routes(r *echo.Group) {
	r.GET("/configs", getConfigs)
	r.POST("/set-config", setActiveConfig)
	r.GET("/active-config", getActiveConfig)

	r.GET("/recommendations", getRecommendations)
	r.POST("/toggle-recommendation", toggleRecommendation)

	r.POST("/chat", ChatCompletion)
}
