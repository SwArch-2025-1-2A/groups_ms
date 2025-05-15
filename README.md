# Main API

## Building the project

### Local build

In order to run the project, you just need to run the following command:

```sh
go build -v -o main .
```

This command compiles the current directory and all its subdirectories, and produces
an executable file named `main` in the current directory. After running this
command, you may [run the project](#local-execution) using the generated executable.

### Building with docker

To build the project using Docker, you can use the provided Dockerfile. The
Dockerfile is set up to build the Go application and create a Docker image.
To build the Docker image, run the following command in the directory where the
Dockerfile is located:

```sh
docker build -t go-api .
```

This command will create a Docker image named `go-api` based on the instructions
in the Dockerfile. And you can run the Docker image using the provided
[run](#using-docker-to-run) command.

## Running the project

### Local execution

To run the project locally, you can use the following command:

```sh
go run main.go
```

Alternatively, if you have built the project using the [build](#local-build) command,
you can run the generated executable:

```sh
./main
```

### Using Docker to run

To run the project using Docker, you can use the following command:

```sh
docker run -p 8080:8080 go-api
```

This command will run the Docker container based on the `go-api` image and map
port 8080 of the container to port 8080 of your host machine. You can then
access the API at `http://localhost:8080`.
