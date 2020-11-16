package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerservice/mgmt/containerservice"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	// create an authorizer from the azure cli config
	a, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		panic(err)
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	// list all AKS clusters within the given subscription
	aksClient := containerservice.NewManagedClustersClient(subscriptionID)
	aksClient.Authorizer = a
	clusters, err := aksClient.List(context.Background())
	if err != nil {
		panic(err)
	}

	for _, cluster := range clusters.Values() {
		c := strings.Split(*cluster.ID, "/")
		resourceGroup := c[4]
		name := c[8]
		fmt.Printf("List cluster admin credentials for %v \n", name)
		cmd := exec.Command("az",
			"aks", "get-credentials", "--name", name,
			"--resource-group", resourceGroup, "--subscription", subscriptionID,
			"--admin", "--overwrite-existing")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}
}

func resourceManagerAuthorizer() (autorest.Authorizer, error) {
	var rmAuth autorest.Authorizer
	var err error
	if len(os.Getenv("AZURE_CLIENT_ID")) == 0 || len(os.Getenv("AZURE_CLIENT_SECRET")) == 0 {
		// create an resource manager authorizer from the az cli configuration
		rmAuth, err = auth.NewAuthorizerFromCLI()
		if err != nil {
			return rmAuth, err
		}
	} else {
		// create an resource manager authorizer from the following environment variables
		// AZURE_CLIENT_ID  | AZURE_CLIENT_SECRET | AZURE_TENANT_ID
		rmAuth, err = auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return rmAuth, err
		}
	}
	return rmAuth, err
}
