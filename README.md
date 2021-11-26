# compose-dev-env
Example used to try a Compose application with Docker Dev Environments.

This example is based on the `nginx-golang-mysql` sample of [`awesome-compose` repository](https://github.com/docker/awesome-compose/).

## Compose sample application
### Go server with an Nginx proxy and a MariaDB database

Project structure:
```
.
├── backend
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── db
│   └── password.txt
├── docker-compose.yaml
├── proxy
│   ├── conf
│   └── Dockerfile
└── README.md
```

[_.docker/docker-compose.yaml_](.docker/docker-compose.yaml)
```
services:
  backend:
    build: backend
    ...
  db:
    image: mariadb
    ...
  proxy:
    build: proxy
    ports:
    - 8080:80
    ...
```
The compose file defines an application with three services `proxy`, `backend` and `db`.
When deploying the application, docker-compose maps port 80 of the proxy service container to port 8080 of the host as specified in the file.
Make sure port 8080 on the host is not already being in use.

## Deploy with docker-compose

```
$ docker-compose up -d
Creating network "compose-dev-env_default" with the default driver
Creating volume "compose-dev-env_db-data" with default driver
Building backend
Step 1/8 : FROM golang:1.13-alpine AS build
1.13-alpine: Pulling from library/golang
...
Successfully built 5f7c899f9b49
Creating compose-dev-env_db_1 ... done
Creating compose-dev-env_backend_1 ... done
Creating compose-dev-env_proxy_1   ... done
```

## Expected result

Listing containers must show three containers running and the port mapping as below:
```
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                  NAMES
d6752b317e6d   compose-dev-env_proxy                 "nginx -g 'daemon of…"   About a minute ago   Up About a minute   0.0.0.0:8080->80/tcp, :::80->80/tcp   compose-dev-env_proxy_1
70bf9182ea52   compose-dev-env_backend               "/server"                About a minute ago   Up About a minute   8000/tcp                            compose-dev-env_backend_1
c67de604799d   mariadb                               "docker-entrypoint.s…"   About a minute ago   Up About a minute   3306/tcp                            compose-dev-env_db_1
```

After the application starts, navigate to `http://localhost:8080` in your web browser or run:
```
$ curl localhost:8080
["Blog post #0","Blog post #1","Blog post #2","Blog post #3","Blog post #4"]
```

Stop and remove the containers
```
$ docker-compose down
```

