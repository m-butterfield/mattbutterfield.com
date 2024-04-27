package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/heatmap"
	_ "image/jpeg"
	"log"
)

func updateHeatmap(c *gin.Context) {
	if err := heatmap.UpdateHeatMap(); err != nil {
		log.Printf("Error updating heatmap: %v", err)
		c.JSON(500, gin.H{"error": "Error updating heatmap"})
		return
	}
}
