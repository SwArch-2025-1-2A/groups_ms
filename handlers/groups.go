package handlers

import (
	"net/http"

	"github.com/SwArch-2025-1-2A/backend/repository"
	"github.com/gin-gonic/gin"
)

type inputGroupCreation struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ProfilePic  string `json:"profilePic"`
	IsOpen      bool   `json:"isOpen"`
}

func CreateGroupsHandler(c *gin.Context) {

	app, ok := GetApp(c)
	if !ok {
		return
	}

	var input inputGroupCreation

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := repository.CreateGroupParams{
		Name:        ToPgText(&input.Name),
		Description: ToPgText(&input.Description),
		ProfilePic:  ToPgText(&input.ProfilePic),
		IsOpen:      ToPgBool(&input.IsOpen),
	}

	grp, err := app.Queries.CreateGroup(app.Context, args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   grp,
	})

}
