# Main API

## Building the project

The recommended way to build the project is to use [Docker](#building-with-docker),
as it ensures that the build environment is consistent and eliminates any potential
issues with dependencies on your local machine. However, if you prefer to build the
project locally, you can do so by following the instructions below.

### Local build

#### Prerequisites for local build

> [!TIP]
> You may not need to install all of the dependencies listed below if you are
> using Docker to build the project.

- Go 1.24.3 or later
- sqlc (for generating SQL code)

1. Install Go from the [official website](https://golang.org/dl/).
2. Install sqlc by running the following command:

    ```sh
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

3. Install the required Go modules by running the following command in the
   project directory:

    ```sh
    go mod tidy
    ```

4. Install the required SQLC modules by running the following command in the
   project directory:

    ```sh
    sqlc generate
    ```

   This command will generate the necessary SQL code based on the SQL files
   in the `db` directory. The generated code will be placed in the `db/sqlc`
   directory.

#### Local build instructions

To build the project locally, you need to have Go 1.24.3 or later installed on your
machine.

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

Alternatively, if you have built the project using the
[build](#local-build-instructions) command, you can run the generated executable:

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
