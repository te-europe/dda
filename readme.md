# Distributed Data Analytics

This repository contains the tools you need to quickly provision, run and kill distributed data analytics environments. 

## Getting started

### Deploy a distributed data analysis environment

Run `./deploy/dda.ps1 deploy` in powershell. 

1. Enter the resource group that you wish to deploy your cluster to, and press enter.
2. Enter the number of cores required for the cluster, and press enter.
3. Enter the amount of memory to be allocated to the cluster, and press enter.

Once the deployment has completed, you will be able to collect your local Python notebook to the distributed cluster using Dask. 

### Remove a cluster

Once finished, to remove a cluster run `./deploy/dda.ps1 deploy`. 

1. Enter the resource group that the cluster was deployed to.
2. Press enter, and the resources will be deleted. 