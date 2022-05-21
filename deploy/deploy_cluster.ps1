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

    Write-Progress -Activity "Deploying a Data Analytics cluster to the resource group '$ResourceGroupName'..."

    Invoke-Expression "az deployment group create --resource-group $ResourceGroupName --template-file ./utils/az-dask-scheduler-deploy.bicep"

    Write-Progress -Completed "Deployment completed."
}


switch ($Command) {
    "deploy" { Deploy }
    Default {}
}