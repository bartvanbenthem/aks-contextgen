package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerservice/mgmt/containerservice"
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
