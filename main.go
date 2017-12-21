package main

import (
	"github.com/gin-gonic/gin"
)

type SlackDetails struct {
	TeamID      string   `json:"team_id"`
	ChannelID   string   `json:"channel_id"`
	APIKey      string   `json:"api_key"`
	ClsuterList []string `json:"clsuter_list"`
}

func main() {
	router := gin.Default()

	router.POST("/details", func(c *gin.Context) {
		// message := c.PostForm("message")
		// nick := c.DefaultPostForm("nick", "anonymous")
		teamID := c.GetHeader("team_id")
		channelID := c.GetHeader("channel_id")
		c.JSON(200, gin.H{
			"status":     "posted",
			"API_KEY":    "api-key-goes-here",
			"team_id":    teamID,
			"channel_id": channelID,
		})
	})
	router.Run(":8080")
}
