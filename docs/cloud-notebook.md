# Running a notebook remotely

More generally, you can run a notebook on a remote machine through VS Code. This gives the benefit of greater processing power, with the ease of editing code on your local machine. 

## Setting up

1. **Starting a remote container**

    First, you need to spin up a remote Jupyter server container. This is done by running the following command in a terminal:

    ```
    az deployment group create --resource-group {RG_NAME} --template-file ./resources/az-jupyter-server-deploy.bicep --parameters Cores={CORES} Memory={MEMORY}
    ```

    where {RG_GROUP} is the name of the resource group you want to use, {CORES} is the number of cores you want to allocate to the container, and {MEMORY} is the amount of memory you want to allocate to the container.

2. **Connecting to the remote Jupyter server**

    There are two ways to use the remote Jupyter server. This first is to access the notebook through the web interface. This is done by following the URL that is displayed when you start the container.

    An easier way to interact with the container is to connect a notebook in VS code to the remote kernel. To do this:

    1. Open a notebook in VS Code.
    2. In the notebook, use `ctrl+shift p` and type: `Jupyter: Specify Jupyter server connection`.
    3. Paste in the URL that is displayed when you start the container.
