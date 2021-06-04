package routes

import (
	"go-ginapi-base/controllers"

	"github.com/gin-gonic/gin"
)

// COnfig Routes
func ConfigRoutes(router *gin.Engine) *gin.Engine {

	c := controllers.NewController()

	// Controllers
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/add", c.AddUser)
			users.GET("/:id", c.GetUser)
		}
		contacts := main.Group("contacts")
		{
			contacts.GET("/:id", c.GetContacts)
		}

	}

	return router
}
