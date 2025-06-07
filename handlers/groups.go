package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SwArch-2025-1-2A/backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateGroupsHandler(c *gin.Context) {

	app, ok := GetApp(c)
	if !ok {
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	isOpen := c.PostForm("isOpen") == "true"

	file, header, err := c.Request.FormFile("profilePic")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar la imagen: " + err.Error()})
		return
	}

	var profilePic []byte
	if file != nil {
		defer file.Close()

		profilePic, err = io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la imagen"})
			return
		}

		if !strings.HasPrefix(header.Header.Get("Content-Type"), "image/") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El archivo debe ser una imagen"})
			return
		}
	}

	args := repository.CreateGroupParams{
		Name:        name,
		Description: ToPgText(&description),
		ProfilePic:  profilePic,
		IsOpen:      isOpen,
	}

	grp, err := app.Queries.CreateGroup(app.Context, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		log.Println(err.Error())
		return
	}

	port := os.Getenv("PORT")

	grpResponse := GroupResponse{
		ID:            grp.ID,
		Name:          grp.Name,
		Description:   grp.Description.String,
		ProfilePicURL: "http://localhost:" + port + "/api/images/" + grp.ID.String(),
		IsVerified:    grp.IsVerified,
		IsOpen:        grp.IsOpen,
		CreatedAt:     grp.CreatedAt.Time,
		UpdatedAt:     grp.UpdatedAt.Time,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   grpResponse,
	})

}

type GroupResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ProfilePicURL string    `json:"profilePicUrl"`
	IsVerified    bool      `json:"isVerified"`
	IsOpen        bool      `json:"isOpen"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func GetGroupsHandler(c *gin.Context) {

	app, ok := GetApp(c)
	if !ok {
		return
	}

	grps, err := app.Queries.GetGroups(app.Context)

	groups := make([]GroupResponse, 0, len(grps))
	port := os.Getenv("PORT")

	for _, g := range grps {
		group := GroupResponse{
			ID:            g.ID,
			Name:          g.Name,
			Description:   g.Description.String,
			ProfilePicURL: "http://localhost:" + port + "/api/images/" + g.ID.String(),
			IsVerified:    g.IsVerified,
			IsOpen:        g.IsOpen,
			CreatedAt:     g.CreatedAt.Time,
			UpdatedAt:     g.UpdatedAt.Time,
		}
		groups = append(groups, group)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   groups,
	})

}

func GetImageHandler(c *gin.Context) {
	app, ok := GetApp(c)
	if !ok {
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	image, err := app.Queries.GetImage(app.Context, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch the Image",
		})
		log.Println(err.Error())
		return
	}

	if len(image) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Writer.Write(image)
}
