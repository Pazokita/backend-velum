package maps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMapsHandler(c *gin.Context) {
	maps := []MapMetadata{
		{
			ID:       1,
			Title:    "Paris 1750",
			ImageURL: "https://example.com/paris-1750.png",
			BBox:     [4]float64{2.3319, 48.8566, 2.3519, 48.8666},
			Opacity:  0.8,
		},
	}
	c.JSON(http.StatusOK, maps)
}
