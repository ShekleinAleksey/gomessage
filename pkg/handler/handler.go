package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ShekleinAleksey/gomessage"
	"github.com/ShekleinAleksey/gomessage/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	var msg gomessage.Message

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		stats := api.Group("/stats")
		{
			stats.GET("/", h.getStats)
		}
	}

	return router
}
