# Description
Generate Kube context file for all AKS clusters within the Azure tenant or within a specific subscription. The azure cli config is used for authentication.

# Description
CLI tool to display all the contacts per namespace on an RBAC enabled Azure Kubernetes cluster. The Azure graph API is consumed directly to query Azure contacts. This tool works best when all the namespace owners on the AKS cluster use the same rolebinding name.

## prerequisites
Install azure cli: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest

Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/

## install and run (Linux)
``` shell
$ git clone https://github.com/bartvanbenthem/aks-contextgen.git
$ ./aks-contextgen/bin/run.sh
```
