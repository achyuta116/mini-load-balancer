## Mini Load Balancer Using Docker
### Objectives
- Write more Golang
- Learn Docker - images, containers, docker compose, docker networks
- Learn Load Balancing & Implement Service Discovery + Round Robin load balancing

### Build It!
#### Docker Images
```sh
$ docker image build -t mlb-server:latest server
$ docker image build -t mlb-balancer:latest .
```
The current directory contains the source files for the balancer itself <br><br>
The `server` directory holds the source files for the servers we're load balancing on

#### Docker Compose
> [!IMPORTANT]
> In case you're trying to run this on Windows or MacOS, you're going to have to execute the following command
> ```sh
> $ docker compose -f docker-compose.dev.yaml -d
> ```
> This will start up a `Docker in Docker` container where you'll be provided with a `docker0` network interface enabling the load balancer to actually reach the containers deployed <br><br>
> Start up a shell prompt within the `parent` container on your host
> ```sh
> $ docker exec -it parent /bin/sh
> ```
> Following which you'll deploy the configuration within `docker-compose.yaml`

- Deploy `docker-compose.yaml`
```sh
$ docker compose up
```
- Inspect the network created (`src_shared`) using `docker inspect`
- Make a note of the IP address generated for the container running the image `mlb-balancer`
- Send requests to the generated IP and observe the responses sent back from different containers!
> [!NOTE]
> Try stopping and restarting containers and observe the effects request routing has on the resulting deployment
> <br>
> You will notice the load balancer picking up on changes within the docker network

