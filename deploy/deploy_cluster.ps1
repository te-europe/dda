param(
    [Parameter(Position = 0, Mandatory = $true, HelpMessage = "The action to perform.")]
    [string] $Command
)

function Deploy {
    # Deploy a Data Analytics cluster to the specified resource group.

    # The parameter for the resource group
    param(
        [Parameter(Position = 0, Mandatory = $true, HelpMessage = "The resource group name.")]
        [string] $ResourceGroupName,
        [Parameter(Position = 1, Mandatory = $true, HelpMessage = "The number of cores to use.")]
        [int] $NumberOfCores,
        [Parameter(Position = 2, Mandatory = $true, HelpMessage = "The amount of memory available to the cluster.")]
        [int] $MemoryInGB
    )

    # Write-Progress -Activity "Deploying a Data Analytics cluster to the resource group '$ResourceGroupName'..."

    # Download the scheulder BICEP template
    $SchedulerTemplate = "https://raw.githubusercontent.com/te-europe/dda/master/resources/az-dask-scheduler-deploy.bicep"
    $WorkerTemplate = "https://raw.githubusercontent.com/te-europe/dda/master/resources/az-dask-worker-deploy.bicep"
    Invoke-WebRequest -Uri $SchedulerTemplate -OutFile "./az-dask-scheduler-deploy.bicep"
    Invoke-WebRequest -Uri $WorkerTemplate -OutFile "./az-dask-worker-deploy.bicep"

    # Create the resources in Azure
    Invoke-Expression "az deployment group create --resource-group $ResourceGroupName --template-file ./az-dask-scheduler-deploy.bicep"
    Invoke-Expression "az deployment group create --resource-group $ResourceGroupName --template-file ./az-dask-worker-deploy.bicep --parameters NumberOfCores=$NumberOfCores MemoryInGB=$MemoryInGB"

    # Delete the BICEP template
    Remove-Item -Path "./az-dask-worker-deploy.bicep" -Force
    Remove-Item -Path "./az-dask-scheduler-deploy.bicep" -Force

}

function Remove {
    # Remove a Data Analytics cluster from the specified resource group.

    # The parameter for the resource group
    param(
        [Parameter(Position = 0, Mandatory = $true, HelpMessage = "The resource group name.")]
        [string] $ResourceGroupName
    )


    # Delete the resources in Azure
    Invoke-Expression "az container delete --resource-group $ResourceGroupName --name daskscheduler -y"
    Invoke-Expression "az container delete --resource-group $ResourceGroupName --name daskworker -y"


}


switch ($Command) {
    "deploy" { Deploy }
    "remove" { Remove }
    Default {}
}