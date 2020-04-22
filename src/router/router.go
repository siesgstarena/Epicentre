package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/services/logger"
	"github.com/siesgstarena/epicentre/src/web"
)

// LoadRouter Configures all routes
func LoadRouter(router *gin.Engine)  {
	logger.Log.Info("Initializing routers")

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "URL Does not exist",
		})
		logger.Log.Warn("Some one trying URL which does not exist")
	})

	handler := router.Group("/")
	{
		handler.GET("health", web.HeathHandler)
		handler.GET("version", web.VersionHandler)
	}

	logger.Log.Info("Initialization of routers Finished")
}