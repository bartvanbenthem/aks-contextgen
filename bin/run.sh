#!/bin/bash

# get all subscriptions in the Azure tenant
subscriptions=$(az account list -o tsv --query [].id)

for s in $subscriptions 
do
    export AZURE_SUBSCRIPTION_ID=$s
    # check if subscription contains AKS clusters
    clusters=$(az aks list --subscription $s -o tsv --query [].name)
    length=${#clusters}
    if [ $length -gt 0 ]
    then
        ./aks-contextgen
    fi
done

