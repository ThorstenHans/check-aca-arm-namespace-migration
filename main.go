package main

import (
	"context"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
	log "github.com/sirupsen/logrus"
)

const newNamespaceName = "Microsoft.App"

func main() {

	subscription_id := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscription_id) == 0 {
		log.Fatal("Please set AZURE_SUBSCRIPTION_ID")
		os.Exit(2)
	}
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Error while loading Azure Credential '%s'\n", err)
		os.Exit(2)
	}

	log.Infof("Scanning Azure Container Apps in subscription %s", subscription_id)
	client := armappservice.NewKubeEnvironmentsClient(subscription_id, cred, nil)
	pager := client.ListBySubscription(nil)

	hasNewNamespace := false
	for {
		nextPage := pager.NextPage(context.Background())
		if err := pager.Err(); err != nil {
			// failed to advance page
			log.Fatalf("Error while paging through list of Azure Container Apps '%s'\n", err)
			os.Exit(3)
		}
		if !nextPage {
			break
		}

		for _, value := range pager.PageResponse().Value {
			log.Infof("Found Azure Container App Environment '%s'", *value.Name)
			if strings.HasPrefix(*value.Type, newNamespaceName) {
				hasNewNamespace = true
				log.Infof("Your Azure Container App Environment '%s' has been mirgrated to the new namespace (%s)", *value.Name, newNamespaceName)
				break
			}
		}
	}

	if hasNewNamespace {
		log.Info("Your Subscription has been (at least partially) been migrated to the new namespaces for Azure Container Apps")
		os.Exit(1)
	}
	log.Info("Your Azure Container Apps are still using early preview namespace (Microsoft.Web)")
	os.Exit(0)
}
