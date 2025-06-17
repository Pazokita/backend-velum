package maps

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMaps(c *gin.Context) {
	maps, err := h.service.GetMaps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}
	log.Println("📦 Maps retournées :", maps)
	c.JSON(http.StatusOK, maps)
}
func (h *Handler) GetMapByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "id invalide"})
		return
	}

	m, err := h.service.GetMap(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "carte non trouvée"})
		return
	}

	c.JSON(200, m)
}
