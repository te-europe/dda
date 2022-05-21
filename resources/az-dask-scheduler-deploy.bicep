param Location string = resourceGroup().location


resource containerGroup 'Microsoft.ContainerInstance/containerGroups@2021-03-01' = {
  name: 'daskscheduler'
  location: Location
  properties: {
    containers: [
      {
        name: 'daskscheduler'
        properties: {
          image: 'ghcr.io/dask/dask'
          ports: [
            {
              port: 8786
            }
            {
              port: 8787
            }
          ]
          command: [
            'dask-scheduler'
          ]
          resources: {
            requests: {
              cpu: 1
              memoryInGB: 4
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
          port: 8786
        }
        {
          protocol: 'TCP'
          port: 8787
        }
      ]
    }
  }
}

output SchedulerIP string = containerGroup.properties.ipAddress.ip
