package main

import (
	"github.com/SwArch-2025-1-2A/backend/app"
	"github.com/SwArch-2025-1-2A/backend/handlers"
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
			// groups.GET("/groups")
		}
	}

	// Beware: this function blocks anything that goes after it
	r.Run()
}
