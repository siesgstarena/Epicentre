package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/logger"
	"github.com/siesgstarena/epicentre/web"
	"github.com/siesgstarena/epicentre/model"
	"github.com/siesgstarena/epicentre/services/github"
	"github.com/siesgstarena/epicentre/services/heroku"
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
		user.GET(":id",model.UserInfo)
		user.POST("", model.CreateUser)
		user.PUT(":id", model.EditUser)
		user.DELETE(":id", model.DeleteUser)
		user.GET(":id/project/:projectid",model.RuleForAProjectConnectedToUser)
	}

	users := router.Group("/users")
	{
		users.GET("all",model.AllUsers)
	}

	admin := router.Group("/admin")
	{
		admin.GET(":id", model.ProjectsWhereUserAdmin)
	}

	project := router.Group("/project")
	{
		project.GET(":id",model.ProjectInfo)
		project.POST("", model.CreateProject)
		project.PUT(":id", model.EditProject)
		project.DELETE(":id", model.DeleteProject)
		project.GET(":id/users",model.AllUsersInProject)
	}

	projects := router.Group("/projects")
	{
		projects.GET("/all",model.AllProjects)
	}

	rule := router.Group("/rule")
	{
		rule.GET(":id",model.RuleInfo)
		rule.POST("", model.CreateRule)
		rule.PUT(":id", model.EditRule)
		rule.DELETE(":id", model.DeleteRule)
	}

	webhook := router.Group("/webhook")
	{
		webhook.POST("heroku", heroku.ReceiveWebhooks)
		webhook.POST("github", github.ReceiveWebhooks)
	}

	herokuRouter := router.Group("/heroku")
	{
		herokuRouter.POST(":id", heroku.SubscribeWebhook)
		herokuRouter.DELETE(":id", heroku.DeleteWebhook)
	}

	githubRouter := router.Group("/github")
	{
		githubRouter.POST(":id", github.SubscribeWebhook)
		githubRouter.DELETE(":id", github.DeleteWebhook)
	}

	logger.Log.Info("Initialization of routers Finished")
}