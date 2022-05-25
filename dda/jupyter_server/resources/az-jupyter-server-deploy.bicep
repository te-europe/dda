param Location string = resourceGroup().location
param Cores int = 1
param Memory int = 1

resource jupyterServer 'Microsoft.ContainerInstance/containerGroups@2021-10-01' = {
  name: 'jupyter-server'
  location: Location
  properties: {
    osType: 'Linux'
    containers: [
      {
        name: 'jupyter-server'
        properties: {
          image: 'jupyter/scipy-notebook:6b49f3337709'
          resources: {
            requests: {
              memoryInGB: Cores
              cpu: Memory
            }
          }
          ports: [
            {
              port: 8888
              protocol: 'TCP'
            }
          ]
        }
      }
    ]
    restartPolicy: 'OnFailure'
    ipAddress: {
      type: 'Public'
      ports: [
        {
          protocol: 'TCP'
          port: 8888
        }
      ]
    }
  }
}
