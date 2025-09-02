// Package routes provides - handles API routes
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamveekthorr/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		// Default route
		apiV1.POST("/shorten", services.CreateURLShorner)
		apiV1.PUT("/shorten", services.UpdateURL)
		apiV1.GET("/shorten/:shortcode", services.GetURLShortCode)
		apiV1.DELETE("/shorten/:shortcode", services.DeleteShortCode)
	}

	router.GET("/:shortcode", services.HandleRedirect)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "Unable to match route to a registered route")
	})

	return router
}
