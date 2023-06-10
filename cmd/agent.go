/*
Copyright © 2023 Josh Whetton <whetton.josh@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jwhett/space-trading/models"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: getAgent,
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var agentUrl = fmt.Sprintf("%s/my/agent", ApiBaseUrl)

func getAgent(cmd *cobra.Command, args []string) {
	request, err := http.NewRequest("GET", agentUrl, nil)
	if err != nil {
		fmt.Printf("Failed to build request: %v\n", err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiToken))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Failure when performing the request: %v\n", err)
	}
	body, err := ioutil.ReadAll(response.Body)
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
