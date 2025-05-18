# Main API

## Modes

The api can run in one these `MODE`s:

- `release`:
- `debug`:
- `test`:

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

To build the whole project using Docker, you can use the provided
[docker-compose.yml](./docker-compose.yml). This will take care of everything related
to the building and running processes, including each needed service to run the
project locally.

By default the API will be built in `release` [mode](#modes) and is built like this:

```sh
docker compose build
```

In case you need to build it in other [`<mode>`](#modes). Run the following command
replacing `<mode>` with the one you need:

```sh
docker compose build --build-arg MODE=<mode>
```

This command will create the docker images the needed images for each service
based on the [docker compose](docker-compose.yml) configuration file. And you can
run the Docker image using the provided [run](#using-docker-to-run) command.

## Running the project

### Local execution

> [!IMPORTANT]
> In order to build and run the project, make sure you meet all the
> [prerequisites](#prerequisites-for-local-build).
<!-- This fixes renderization issues -->
> [!NOTE]
> These commands will not run your database or any other dependency/service
> that you may need

If you have built the project using the
[build](#local-build-instructions) command, you can run the generated executable:

```sh
./main
```

To run the project locally without building, you can use the following command:

```sh
go run main.go
```

### Using Docker to run

We personally recommend [this option](#running-while-rebuilding-images).

#### Running without rebuilding

It doesn't contemplate the current state of your files but the most recent built
images. Make sure it is in the correct [`mode`](#modes)

```sh
docker compose up
```

#### Running while rebuilding images

This will allow you to start from the current version of your repository.
This means that if you modified any file included in Docker, a new image will be
built in the default mode and run after the building process has been completed

```sh
docker compose up --build
```
