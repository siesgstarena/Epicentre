package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/logger"
	"github.com/siesgstarena/epicentre/web"
	"github.com/siesgstarena/epicentre/model"
)

// LoadRouter Configures all routes
func LoadRouter(router *gin.Engine) {
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

	user := router.Group("/user")
	{
		user.GET("/:id",model.UserInfo)
		user.POST("create", model.CreateUser)
		user.PUT("edit/:id", model.EditUser)
		user.DELETE("delete/:id", model.DeleteUser)
		user.GET("/:id/project/:projectid",model.RuleForAProjectConnectedToUser)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/:id", model.ProjectsWhereUserAdmin)
	}

	project := router.Group("/project")
	{
		project.GET("/:id",model.ProjectInfo)
		project.POST("create", model.CreateProject)
		project.PUT("edit/:id", model.EditProject)
		project.DELETE("delete/:id", model.DeleteProject)
		project.GET("/:id/users",model.AllUsersInProject)
	}

	rule := router.Group("/rule")
	{
		rule.GET("/:id",model.RuleInfo)
		rule.POST("create", model.CreateRule)
		rule.PUT("edit/:id", model.EditRule)
		rule.DELETE("delete/:id", model.DeleteRule)
	}

	webhook := router.Group("/webhook")
	{
		webhook.POST("heroku", web.ReceiveWebhooks)
	}

	logger.Log.Info("Initialization of routers Finished")
}
