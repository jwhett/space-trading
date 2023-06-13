package api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/jwhett/space-trading/models"
	"github.com/spf13/cobra"
)

var agentUrl = fmt.Sprintf("%s/my/agent", apiBaseUrl)

func GetAgent(cmd *cobra.Command, args []string) {
	response, err := makeRequest("GET", agentUrl, nil)
	if err != nil {
		fmt.Printf("request failure: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v\n", err)
		os.Exit(1)
	}
	var ad models.AgentData
	err = json.Unmarshal([]byte(body), &ad)
	if err != nil {
		fmt.Printf("failed to unmarshall body into an Agent: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Name: %s\nFaction: %s\nHQ: %s\nCredits: %d\n", ad.Data.Symbol, ad.Data.StartingFaction, ad.Data.Headquarters, ad.Data.Credits)
}
