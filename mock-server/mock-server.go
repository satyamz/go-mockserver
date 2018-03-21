package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type Topology struct {
	HideIfEmpty bool   `json:"hide_if_empty"`
	Name        string `json:"name"`
	Options     []struct {
		DefaultValue string `json:"defaultValue"`
		ID           string `json:"id"`
		Options      []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"options"`
		NoneLabel  string `json:"noneLabel,omitempty"`
		SelectType string `json:"selectType,omitempty"`
	} `json:"options"`
	Rank  int `json:"rank"`
	Stats struct {
		EdgeCount          int `json:"edge_count"`
		FilteredNodes      int `json:"filtered_nodes"`
		NodeCount          int `json:"node_count"`
		NonpseudoNodeCount int `json:"nonpseudo_node_count"`
	} `json:"stats"`
	SubTopologies []struct {
		HideIfEmpty bool   `json:"hide_if_empty"`
		Name        string `json:"name"`
		Options     []struct {
			DefaultValue string `json:"defaultValue"`
			ID           string `json:"id"`
			Options      []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"options"`
			NoneLabel  string `json:"noneLabel,omitempty"`
			SelectType string `json:"selectType,omitempty"`
		} `json:"options"`
		Rank  int `json:"rank"`
		Stats struct {
			EdgeCount          int `json:"edge_count"`
			FilteredNodes      int `json:"filtered_nodes"`
			NodeCount          int `json:"node_count"`
			NonpseudoNodeCount int `json:"nonpseudo_node_count"`
		} `json:"stats"`
		URL string `json:"url"`
	} `json:"sub_topologies,omitempty"`
	URL string `json:"url"`
}

func main() {
	router := gin.Default()
	router.GET("/test", topology)
	router.Run(":4040")

}

/*-------------------------------------------------------*/

func topology(c *gin.Context) {
	raw, err := ioutil.ReadFile("json/topology.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var x []Topology
	json.Unmarshal(raw, &x)
	c.JSON(200, x)
}
