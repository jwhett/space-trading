package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jwhett/space-trading/models"
	"github.com/spf13/cobra"
)

var agentUrl = fmt.Sprintf("%s/my/agent", apiBaseUrl)

func GetAgent(cmd *cobra.Command, args []string) {
	request, err := http.NewRequest("GET", agentUrl, nil)
	if err != nil {
		fmt.Printf("Failed to build request: %v\n", err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Failure when performing the request: %v\n", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
	}
	var ad models.AgentData
	err = json.Unmarshal([]byte(body), &ad)
	if err != nil {
		fmt.Printf("Failed to unmarshall body into an Agent: %v\n", err)
	}
	fmt.Printf("Name: %s\nFaction: %s\nHQ: %s\nCredits: %d\n", ad.Data.Symbol, ad.Data.StartingFaction, ad.Data.Headquarters, ad.Data.Credits)
}
