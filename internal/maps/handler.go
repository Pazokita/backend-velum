package maps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMapsHandler(c *gin.Context) {
	maps := GetAllMaps()
	c.JSON(http.StatusOK, maps)
}
