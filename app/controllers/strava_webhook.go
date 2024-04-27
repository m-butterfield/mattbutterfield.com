package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StravaWebhookRequest struct {
	AspectType     string         `json:"aspect_type"`
	EventTime      int            `json:"event_time"`
	ObjectId       int            `json:"object_id"`
	ObjectType     string         `json:"object_type"`
	OwnerId        int            `json:"owner_id"`
	SubscriptionId int            `json:"subscription_id"`
	Updates        map[string]any `json:"updates"`
}

func stravaWebhook(c *gin.Context) {
	// The below is required if adding a new webhook subscription (update router to accept GET requests here as well)
	// See:https://developers.strava.com/docs/webhooks/
	//if c.Request.Method == http.MethodGet {
	//	c.JSON(http.StatusOK, gin.H{"hub.challenge": c.Query("hub.challenge")})
	//	return
	//}
	c.Status(http.StatusOK)
	body := &StravaWebhookRequest{}
	err := c.Bind(body)
	if err != nil {
		log.Printf("Error binding request: %v", err)
		return
	}
	log.Printf("Received webhook: %v", body)

	if body.AspectType == "create" && body.ObjectType == "activity" {
		log.Printf("New activity, creating task to update heatmap")
		if _, err = tc.CreateTask("update_heatmap", "update-heatmap", nil); err != nil {
			log.Printf("Error creating task: %v", err)
			return
		}
	}
}
