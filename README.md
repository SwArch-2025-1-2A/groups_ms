# Main API

## Database migrations

The project uses [migrate](https://github.com/golang-migrate/migrate/). Here
are some useful commands for working with migrations:

### Create a new migration

To create a new migration, run the following command:

```sh
migrate create -ext sql -dir db/migrations -seq <migration_name>
```

Replace `<migration_name>` with a descriptive name for your migration. This
command will create a new migration file in the `db/migrations` directory with
the specified name. The `-ext sql` flag specifies that the migration files
should have the `.sql` extension, and the `-seq` flag specifies that the
migration should be created with a sequential number.

### Apply migrations

To apply all pending migrations, run the following command:

```sh
migrate -path db/migrations -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -verbose up
```

<!-- ### Rollback migrations

To rollback the last applied migration, run the following command:

```sh
migrate -path db/migrations -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" down
``` -->

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
2. Install sqlc by running the following command. Or
visit [this page](https://docs.sqlc.dev/en/latest/overview/install.html#installing-sqlc):

    ```sh
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

3. Install the required Go modules by running the following command in the
project directory:

    ```sh
    go mod tidy
    ```

4. Generate the required SQLC modules by running the following command in the
project directory. Make sure you have
[applied all your migrations](#apply-migrations):

    ```sh
    sqlc generate
    ```

   This command will generate the necessary go code based on the SQL files
   in the `db` directory. The generated code will be placed in the `./repository`
   directory.

5. Generate the GraphQL models and resolver(s) by running the following command
in the project directory:

    ```sh
    go run github.com/99designs/gqlgen generate
    ```

#### Local build instructions

> [!IMPORTANT]
> In order to build and run the project, make sure you meet all the
> [prerequisites](#prerequisites-for-local-build).

After that you can build and run the project. Like this:

```sh
go build -v -o main .
```

This command compiles the current directory and all its subdirectories, and produces
an executable file named `main` in the current directory. After running this
command, you may [run the project](#local-execution) using the generated executable.

### Building with docker

> [!IMPORTANT]
> Make sure your `.dockerignore` file is set up correctly to exclude any
> unnecessary files. Like so:

```sh
cat .gitignore .prodignore > .dockerignore
```

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

> [!IMPORTANT]
> In order to build and run the project, make sure you meet all the
> [prerequisites](#prerequisites-for-local-build).

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
