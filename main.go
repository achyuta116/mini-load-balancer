package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

var (
	index        int = 0
	servers      []string
	serversMutex sync.Mutex
)

func updateServices(cli *client.Client) {
	networkName := "src_shared"
	imageName := "mlb-server"

	filter := filters.NewArgs()
	filter.Add("network", networkName)
	filter.Add("ancestor", imageName)

	for {
		list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
			Filters: filter,
		})
		if err != nil {
			log.Fatalf(err.Error())
		}

		serversMutex.Lock()
		servers = make([]string, 0)

		for _, container := range list {
			ipAddress := container.NetworkSettings.Networks[networkName].IPAddress
			servers = append(servers, ipAddress)
		}
		serversMutex.Unlock()

        if len(servers) != 0 {
            index = index % len(servers)
        } else {
            index = 0
        }
		time.Sleep(time.Second * 2)
	}
}

func getNextServer() string {
	serversMutex.Lock()
	server := servers[index]
	index = (index + 1) % len(servers)
	serversMutex.Unlock()
	return server
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	server := getNextServer()
	uri, _ := url.Parse("http://" + server + ":8000")
    fmt.Println(uri)
	reverseProxy := httputil.NewSingleHostReverseProxy(uri)
	reverseProxy.ServeHTTP(w, r)
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	cli.NegotiateAPIVersion(context.Background())

	if err != nil {
		log.Fatal(err.Error())
	}

	go updateServices(cli)

	r := http.NewServeMux()
	r.Handle("/", http.HandlerFunc(handleRequest))
	http.ListenAndServe(":8000", r)
}
