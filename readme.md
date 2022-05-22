# Distributed Data Analytics

This repository contains the tools you need to quickly provision, run and kill distributed data analytics environments. 

## Getting started

### Prerequisites

To use the DDA CLI, and to use the az CLI directly you will need to download it first. You can do this by following these instructions: [Install the az CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-windows?tabs=azure-cli).

You will only be able to provision Azure resource if given the correct permissions from the IT Team. 

### Using the DDA CLI
#### Deploy a distributed data analysis environment

Run `./deploy/dda.ps1 deploy` in powershell. 

1. Enter the resource group that you wish to deploy your cluster to, and press enter.
2. Enter the number of cores required for the cluster, and press enter.
3. Enter the amount of memory to be allocated to the cluster, and press enter.

Once the deployment has completed, you will be able to collect your local Python notebook to the distributed cluster using Dask. 

#### Remove a cluster

Once finished, to remove a cluster run `./deploy/dda.ps1 deploy`. 

1. Enter the resource group that the cluster was deployed to.
2. Press enter, and the resources will be deleted. 


### Using the az CLI

#### Deploying a dask cluster

First, you need to provision a Dask scheduler. To do this, run the following command in your PowerShell terminal:

```az deployment group create --resource-group {Name of target resource group} --template-file ./resources/az-dask-scheduler-deploy.bicep```

Once this has completed, run the command to spin up a worker node, pass it the parameters for how many cpus and memory to allocate to the machine:

```az deployment group create --resource group {RG_NAME} --template-file ./resources/az-dask-worker-deploy.bicep --parameters NumberOfCores=1 MemoryInGB=1```

To spin up multiple worker containers, run the above command multiple times with the required specifications. 

#### Stop the containers

Once finished, you can kill the environment by running the commands:

```az container delete --name daskscheduler --resource-group {RG_NAME} ```

```az container delete --name daskworker --resource-group {RG_NAME} ```