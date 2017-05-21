package probe

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"os"
)

func DoDockerSwarmProbing() {
	fmt.Println("===============================")
	fmt.Println("===============================")
	fmt.Println("===== Docker Swarm Probe ======")
	fmt.Println("===============================")

	host := "unix:///var/run/docker.sock"
	if len(os.Getenv("DF_DOCKER_HOST")) > 0 {
		host = os.Getenv("DF_DOCKER_HOST")
	}
	fmt.Println("Host: " + host)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	fmt.Println("===============================")
	fmt.Println("== List Running Containers ====")
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("== %s \t %s\n", container.ID[:10], container.Image)
	}
	fmt.Println("===============================")

	fmt.Println("===============================")
	fmt.Println("== List Services ====")

	filter := filters.NewArgs()
	filter.Add("label", fmt.Sprintf("%s=true", "com.df.notify"))
	filter.Add("label", fmt.Sprintf("%s", "nl.flusso.drove.service"))
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{Filters: filter})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, service := range services {
			fmt.Println("--------")
			fmt.Println(service)
			fmt.Println("--------")
		}
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
