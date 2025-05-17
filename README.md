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

### Rollback migrations

To rollback the last applied migration, run the following command:

```sh
migrate -path db/migrations -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" down
```

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
docker compose build
```

This command will create a Docker the needed images for each service based on the
[docker compose](docker-compose.yml) configuration file. And you can run the Docker
image using the provided [run](#using-docker-to-run) command.

## Running the project

### Local execution

To run the project locally, you can use the following command.
This will not run your database or any other dependency/service that you may need:

```sh
go run main.go
```

Alternatively, if you have built the project using the
[build](#local-build-instructions) command, you can run the generated executable:

```sh
./main
```

### Using Docker to run

We personally recommend [this option](#running-while-rebuilding-images).

#### Running without rebuilding

It doesn't contemplate the current state of your files but the most recent built
images.

```sh
docker compose up
```

#### Running while rebuilding images

This will allow you to start from the current version of your repository.
This means that if you modified any file included in Docker, a new image will be
built.

```sh
docker compose up --build
```