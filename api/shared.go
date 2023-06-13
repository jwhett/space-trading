package api

import (
	"os"
)

const apiBaseUrl = "https://api.spacetraders.io/v2"

var apiToken = os.Getenv("ST_TOKEN")
