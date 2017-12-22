package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SlackDetails []struct {
	TeamID      string   `json:"team_id"`
	ChannelID   string   `json:"channel_id"`
	APIKey      string   `json:"api_key"`
	ClsuterList []string `json:"clsuter_list"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/details", GetDetails)
	router.Run(":8080")
}

func GetDetails(c *gin.Context) {
	data := []byte(`[
    {
    "team_id" : "T8HP5N491",
    "channel_id": "C8JMAKKPG",

    "api_key": "YmF0Y2F2ZQ==",
    "clsuter_list": [
        "botcave-clsuter-one", "botcave-cluster-two" ]
 },
{
    "team_id" : "T8HP5N491",
    "channel_id": "C8JT96WDD",

    "api_key": "YmF0Y2F2ZQ==",
    "clsuter_list": [
        "botcave-clsuter-three", "botcave-cluster-four" ]
},

    {
    "team_id" : "T88KY3BFV",
    "channel_id": "C88SBRW21",

    "api_key": "YXNnYXJk=",
    "clsuter_list": [
        "asgard-clsuter-one", "asgard-cluster-two" ]
},

    {
    "team_id" : "T88KY3BFV",
    "channel_id": "C8AAYPM4J",

    "api_key": "YXNnYXJk=",
    "clsuter_list": [
        "asgard-clsuter-three", "asgard-cluster-four" ]
},
    {
    "team_id" : "T88KY3BFV",
    "channel_id": "C89QA2L95",

    "api_key": "YXNnYXJk=",
    "clsuter_list": [
        "asgard-clsuter-five", "asgard-cluster-six" ]
}
]`)

	teamID := c.PostForm("team_id")
	channelID := c.PostForm("channel_id")

	fmt.Println("Team ID: ", teamID, " channelID: ", channelID)
	var slackDetails SlackDetails
	err := json.Unmarshal(data, &slackDetails)
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
