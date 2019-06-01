package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nassimelhormi/adserver/handlers"
	"github.com/nassimelhormi/adserver/models"
)

/***
**		function	parseFlags
**		@param		empty
**		@return		filename string, port string
**		Parse the command line to get the filename and the port for running the adserver
***/
func parseFlags() (filename string, port string) {

	flag.StringVar(&filename, "file", "datas/campaigns.json", "The data file. By default we use our datas.")
	flag.StringVar(&port, "port", "4242", "The server port. By default it's 8080.")
	flag.Parse()

	return filename, port
}

/***
**		function	campaignsToCampaign
**		@param		Campaigns map[string]Campaign
**		@return		Campaign Struct
**		Convert a Campaigns map into a Campaign
***/
func campaignsToCampaign(campaigns models.Campaigns) []models.Campaign {
	temp := []models.Campaign{}

	for id, campaign := range campaigns.Campaigns {
		campaign.ID = id
		fmt.Println(campaign.ID)
		temp = append(temp, campaign)
	}
	return temp
}

/***
**		function	singletonJSON
**		@param		filename string
**		@return		Campaign Struct
** 		Parse a JSON file into a Campaign Struct
***/
func singletonJSON(filename string) []models.Campaign {
	campaigns := models.Campaigns{}
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &campaigns)

	if err != nil {
		log.Fatal(err)
	}
	campaign := campaignsToCampaign(campaigns)

	return campaign
}

/***
**		function main
**
**
***/
func main() {
	filename, port := parseFlags()
	log.Println(filename, port)
	log.Println("<--- Adserver Started ! --->")
	campaign := singletonJSON(filename)

	// Handle the server
	http.Handle("/ad", handlers.AdServerHandler(campaign))
	http.ListenAndServe(":"+port, nil)
}
