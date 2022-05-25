module Scheduler 'az-dask-scheduler-deploy.bicep' = {
  name: 'scheduler-module'
}

param Location string = resourceGroup().location

param NumberOfCores int
param MemoryInGB int

resource containerGroup 'Microsoft.ContainerInstance/containerGroups@2021-03-01' = {
  name: 'daskworker'
  location: Location
  properties: {
    containers: [
      {
        name: 'dask-worker'
        properties: {
          image: 'ghcr.io/dask/dask'
          command: [
            'dask-worker'
            '${Scheduler.outputs.SchedulerIP}:8786'
          ]
          ports: [
            {
              port: 80
            }
          ]
          resources: {
            requests: {
              cpu: NumberOfCores
              memoryInGB: MemoryInGB
            }
          }
        }
      }
    ]
    restartPolicy: 'OnFailure'
    osType: 'Linux'
    ipAddress: {
      type: 'Public'
      ports: [
        {
          protocol: 'TCP'
          port: 80
        }
      ]
    }
  }
}
