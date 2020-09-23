# Description
Generate kube-context containing all AKS clusters within the Azure tenant or within a specific subscription. The azure cli config is used for authentication.

## prerequisites
Install azure cli: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest

Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/

## install and run (Linux)
Create kube-context for all AKS clusters within the Azure tenant:
``` shell
$ az login
$ git clone https://github.com/bartvanbenthem/aks-contextgen.git
$ cd aks-contextgen/bin
$ ./run.sh
```

Create kube-context for all AKS clusters within a specific subscription:
``` shell
$ az login
$ git clone https://github.com/bartvanbenthem/aks-contextgen.git
$ export AZURE_SUBSCRIPTION_ID='<<subscriptionID>>'
$ cd aks-contextgen/bin
$ ./aks-contextgen
```

install and run (Windows)
``` shell
# in development
```
