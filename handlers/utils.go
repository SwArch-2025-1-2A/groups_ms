package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SwArch-2025-1-2A/groups_ms/app"
	"github.com/SwArch-2025-1-2A/groups_ms/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetApp(c *gin.Context) (*app.App, bool) {
	appInterface, exists := c.Get("app")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("app configuration not available")
		return nil, false
	}

	app, ok := appInterface.(*app.App)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("wrong app type")
		return nil, false
	}

	return app, true
}

func ToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func ToPgBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *b, Valid: true}
}

// Generate the Image URL
func GenerateImageURL(id uuid.UUID) string {
	port := os.Getenv("PORT")
	hostname := os.Getenv("LOCALHOST")
	return "http://" + hostname + ":" + port + "/api/images/" + id.String()
}

// Bind the DB response to the API response when it has group(s) to give back
func BindGroupResponse(grp repository.Group) GroupResponse {
	grpResponse := GroupResponse{
		ID:            grp.ID,
		Name:          grp.Name,
		Description:   grp.Description.String,
		ProfilePicURL: GenerateImageURL(grp.ID),
		IsVerified:    grp.IsVerified,
		IsOpen:        grp.IsOpen,
		CreatedAt:     grp.CreatedAt.Time,
		UpdatedAt:     grp.UpdatedAt.Time,
	}

	return grpResponse
}
