package main

import (
	"github.com/SwArch-2025-1-2A/groups_ms/app"
	"github.com/SwArch-2025-1-2A/groups_ms/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	app := app.NewApp()
	defer app.DBPool.Close()

	// Setup gin
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	})

	api := r.Group("/api")
	{
		groups := api.Group("/groups")
		{
			groups.POST("", handlers.CreateGroupsHandler)
			groups.GET("", handlers.GetGroupsHandler)
			groups.GET("/:id", handlers.GetGroupByIDHandler)
			groups.DELETE("/:id", handlers.DeleteGroupHandler)
		}
		images := api.Group("/images/:id")
		{
			images.GET("", handlers.GetImageHandler)
		}
	}

	// Beware: this function blocks anything that goes after it
	r.Run()
}
