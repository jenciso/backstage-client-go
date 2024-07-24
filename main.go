package main

import (
	"os"

	"github.com/tdabasinskas/go-backstage/v2/backstage"
)

func main() {

	baseURL, ok := os.LookupEnv("BACKSTAGE_BASE_URL")
	if !ok {
		baseURL = "http://localhost:7007/api/"
	}

	c, _ := backstage.NewClient(baseURL, "default", nil)

	PrintDomainsAndSystems(c)
	PrintAppListWithSystemsAndDomains(c)
	PrintDomains(c)
}
