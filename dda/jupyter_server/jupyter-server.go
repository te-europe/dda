package jupter_server

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
	ffs "jackminchin.me/dda/file_service"
)

const (
	JupyterServer string = "https://raw.githubusercontent.com/te-europe/dda/master/resources/az-jupyter-server-deploy.bicep"
	DaskScheduler string = "https://raw.githubusercontent.com/te-europe/dda/master/resources/az-dask-scheduler-deploy.bicep"
	DaskWorker    string = "https://raw.githubusercontent.com/te-europe/dda/master/resources/az-dask-worker-deploy.bicep"
)

func DeployJupyterServer(c *cli.Context) error {
	// Read in the cores and memory flags
	cores, err := strconv.Atoi(c.String("cores"))
	if err != nil {
		log.Fatal(err)
	}

	memory, err := strconv.Atoi(c.String("memory"))
	if err != nil {
		log.Fatal(err)
	}

	resourceGroup := c.String("resource-group")

	cores_ph := 0.03887
	memory_ph := 0.00427

	cost_ph := float64(cores)*cores_ph + float64(memory)*memory_ph
	cost_pm := cost_ph * 730

	// Print out the task and await confirmation
	// Create table with the resources for confirmatin

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Resource", "Cores", "Memory", "Cost (p/h)", "Cost (p/m)"})
	t.AppendRow(table.Row{"Jupyter Server", cores, fmt.Sprintf("%s GB", strconv.Itoa(memory)), fmt.Sprintf("£%.2f", cost_ph), fmt.Sprintf("£%.2f", cost_pm)})

	t.Render()

	fmt.Printf("Are you sure you want to deploy this? (y/n) ")

	// Read in the user's response
	var response string
	fmt.Scanln(&response)

	// If the user responded with "y", deploy the task
	if response == "y" {
		// Start the loading bar
		s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
		s.Suffix = " Deploying Jupyter Server Container"
		s.Start()

		// Get the bicep file content
		tempSpecPath, err := ffs.DownloadFileFromRepo(JupyterServer)
		if err != nil {
			color.Red("Error downloading jupyter-server bicep file")
			return &DeploymentError{}
		}

		// print out files
		deployCommand := exec.Command("az", "deployment", "group", "create", "--name", "jupyter-server", "--resource-group", resourceGroup, "--template-file", tempSpecPath, "--parameters", "Cores="+strconv.Itoa(cores), "Memory="+strconv.Itoa(memory))
		output, err := deployCommand.CombinedOutput()
		if err != nil {
			color.Red("Error deploying jupyter-server")
			fmt.Println(string(output))
			return &DeploymentError{}
		}

		// Start the command
		deployCommand.Start()

		// Check the status of the Command
		err = deployCommand.Wait()
		if err != nil {
			s.Stop()
			color.Red(err.Error())
			return &DeploymentError{}
		}

		// Delete the temporary file
		err = ffs.DeleteFromTemporary(tempSpecPath)
		if err != nil {
			s.Stop()
			log.Fatal(err)
		}

		// Print success
		s.Stop()

		// Marshall output into DeploymentResponse
		// var dr DeploymentResponse
		// err = json.Unmarshal(output, &dr)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Println(output)
		// Print the output
		// fmt.Println(dr.Properties.OutputResources[0].Id)

	} else {
		color.Green("Deployment cancelled")
	}

	return nil
}

type DeploymentError struct{}

func (m *DeploymentError) Error() string {
	return "Deployment failed"
}

type DeploymentResponseProperties struct {
	CorrelationId     string           `json:"correlationId"`
	DebugSetting      string           `json:"debugSetting"`
	Dependencies      []string         `json:"dependencies"`
	Duration          string           `json:"duration"`
	Error             string           `json:"error"`
	Mode              string           `json:"mode"`
	OnErrorDeployment string           `json:"onErrorDeployment"`
	OutputResources   []OutputResource `json:"outputResources"`
}

type DeploymentResponse struct {
	Id            string                       `json:"id"`
	Location      string                       `json:"location"`
	Name          string                       `json:"name"`
	Properties    DeploymentResponseProperties `json:"properties"`
	ResourceGroup string                       `json:"resourceGroup"`
}

type OutputResource struct {
	Id            string `json:"id"`
	ResourceGroup string `json:"resourceGroup"`
}
