### Go Gin framework



#### Setting up the containers
----
We are dealing with multiple containers as small services. Example each of the below services will have a container of its own 

- Api Application 
- Database
- Nginx reverse proxy

When dealing with multiple containers as such we need some orchestration between them - hence we have chosen the simplest of them to learn : `docker-compose`

##### Whats `docker-compose` ?
---

[Read more here](https://docs.docker.com/get-started/08_using_compose/)
Docker Compose is a tool that was developed to help define and share multi-container applications. With Compose, we can create a YAML file to define the services and with a single command, can spin everything up or tear it all down. The big advantage of using Compose is you can define your application stack in a file, keep it at the root of your project repo (it’s now version controlled), and easily enable someone else to contribute to your project. Someone would only need to clone your repo and start the compose app. In fact, you might see quite a few projects on GitHub/GitLab doing exactly this now.

##### Getting it installed:
----

```sh
# If you have no clue on how to install yay 
# Plus if your system is brand new, here is procedure to install yay using pacman
# Im assuming you are on a Arch based system 
yay docker
yay docker-compose
# if you dont modify the group you would then have to append sudo to each docker command i
sudo usermod -aG docker <username>
# systemctl is specific to arch based systems
# If you are on a Debian derivative the commands will defer
sudo systemctl enable docker.service
sudo systemctl start docker.service
```
If the installation was successful the below commands shall give you the . Output will confirm you are now ready to build your application stack and take your setup for a spin

```sh
$ docker --version
Docker version 20.10.21, build baeda1f82a
```
```sh
$ docker images   
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
```

##### Forming the stack
----

```yml
version: '3.1'

services: 
  apiapp:
    build:
      context: .
      dockerfile: Dockerfile
      args:
        - APPNAME=learn-gogin
        - ORGNAME=ncs
    ports: 
      - 8080:8080
    tty: true
    stdin_open: true
    entrypoint: [/usr/bin/ncs/learn-gogin]
    container_name: ctn_apiapp
```

Lets build a docker container stack by defining services in a docker-compose.yml file. API application would be a service that has its own dockerfile.
Port 8080 from within the container is mapped to 8080 on the host machine. Entrypoint is the GO executable

```dockerfile
FROM golang:1.18-alpine
ARG APPNAME
ARG ORGNAME
RUN apk update 
RUN mkdir -p /usr/src/${ORGNAME} /usr/bin/${ORGNAME} /var/log/${ORGNAME} 
WORKDIR /usr/src/${ORGNAME}
COPY go.sum go.mod ./
RUN go mod download 
COPY . .
RUN go build -o /usr/bin/{ORGNAME}/${APPNAME} .
```
Shift back to the directory (If you are not already in it) where `docker-compose` resides and issue the following commands

```sh 
$ docker-compose build
$ docker images   
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
learn-gogin-apiapp   latest    84f2b3d2cbd1   8 minutes ago   506MB
```
Voila! you now have a docker image ready to be spun up as a go application container.

```sh
$ docker-compose up    
[+] Running 1/1
 ⠿ Container ctn_apiapp  Recreated                                                                                                                                                       0.1s
Attaching to ctn_apiapp
ctn_apiapp  | INFO[0000] Starting the go gin application              
ctn_apiapp  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
ctn_apiapp  | 
ctn_apiapp  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
ctn_apiapp  |  - using env:	export GIN_MODE=release
ctn_apiapp  |  - using code:	gin.SetMode(gin.ReleaseMode)
ctn_apiapp  | 
ctn_apiapp  | [GIN-debug] GET    /ping                     --> main.PingHandler (3 handlers)
ctn_apiapp  | [GIN-debug] GET    /employees/               --> main.main.func1 (4 handlers)
ctn_apiapp  | [GIN-debug] GET    /employees/:id            --> main.EmployeeHandler (4 handlers)
ctn_apiapp  | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
ctn_apiapp  | Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
ctn_apiapp  | [GIN-debug] Listening and serving HTTP on :8080

```
Here you can see above, when I run the container with a single service it actually starts to serve the application from behind the container.
