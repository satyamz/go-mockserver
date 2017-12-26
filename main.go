package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SlackDetails []struct {
	TeamID          string   `json:"team_id"`
	IncomingWebhook string   `json:"inc_webhook"`
	ChannelID       string   `json:"channel_id"`
	APIKey          string   `json:"api_key"`
	ClsuterList     []string `json:"clsuter_list"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/details", GetDetails)
	router.POST("/alerts", GetDetailsUsingCluster)
	router.Run(":8080")
}

// GetDetails returns slack information based on the team id and channel id
func GetDetails(c *gin.Context) {

	teamID := c.PostForm("team_id")
	channelID := c.PostForm("channel_id")

	fmt.Println("Team ID: ", teamID, " channelID: ", channelID)
	var slackDetails SlackDetails
	err := json.Unmarshal(getData(), &slackDetails)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(slackDetails)
	for _, v := range slackDetails {
		if v.TeamID == teamID && v.ChannelID == channelID {
			c.JSON(200, gin.H{
				"status":       "posted",
				"API_KEY":      v.APIKey,
				"team_id":      v.TeamID,
				"channel_id":   v.ChannelID,
				"cluster_list": v.ClsuterList,
			})
		}
	}
}

// GetDetailsUsingCluster returns slack details based on cluster Id
func GetDetailsUsingCluster(c *gin.Context) {
	clusterIdentity := c.PostForm("cluster_id")

	fmt.Println("Cluster ID: ", clusterIdentity)

	var slackDetails SlackDetails

	err := json.Unmarshal(getData(), &slackDetails)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("%+v\n", slackDetails)

	for _, v := range slackDetails {
		for _, clusterID := range v.ClsuterList {
			if clusterIdentity == clusterID {
				c.JSON(200, gin.H{
					"status":      "posted",
					"inc_webhook": v.IncomingWebhook,
					"team_id":     v.TeamID,
					"channel_id":  v.ChannelID,
				})

			}
		}
	}
}

//getData returns mock data
func getData() []byte {

	data := []byte(`[
    {
    "team_id"    : "T8HP5N491",
    "channel_id" : "C8JMAKKPG",
	"inc_webhook": "https://hooks.slack.com/services/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",
    "api_key": "YmF0Y2F2ZQ==",
    "clsuter_list": [
        "botcave-clsuter-one", "botcave-cluster-two" ]
 },
{
    "team_id" : "T8HP5N491",
    "channel_id": "C8JT96WDD",
	"inc_webhook" : "https://hooks.slack.com/services/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",
    "api_key": "YmF0Y2F2ZQ==",
    "clsuter_list": [
        "botcave-clsuter-three", "botcave-cluster-four" ]
},

    {
    "team_id" : "T88KY3BFV",
    "channel_id": "C88SBRW21",
	"inc_webhook": "https://hooks.slack.com/services/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",
    "api_key": "YXNnYXJk=",
    "clsuter_list": [
        "asgard-clsuter-one", "asgard-cluster-two" ]
},

    {
    "team_id" : "T88KY3BFV",
    "channel_id": "C8AAYPM4J",
	"inc_webhook": "https://hooks.slack.com/services/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",
    "api_key": "YXNnYXJk=",
    "clsuter_list": [
        "asgard-clsuter-three", "asgard-cluster-four" ]
},
    {
    "team_id" : "T88KY3BFV",
    "channel_id": "C89QA2L95",
	"inc_webhook": "https://hooks.slack.com/services/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",
    "api_key": "YXNnYXJk=",
    "clsuter_list": [
        "asgard-clsuter-five", "asgard-cluster-six" ]
}
]`)
	return data
}
