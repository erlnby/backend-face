# backend-face

Service for Financial university project

## Developing

### Setting up dev

To start developing the project use commands below:

```shell
git clone github.com/SeymourCray/backend-face
go mod download
```

### Starting the project

The project is launched using Docker, so make sure it is installed on your machine:

```shell
docker-compose build
docker-compose up
```

### Setting up environment variables

To use example environment variables do:

```shell
cp .env.example .env
```

### Migrations

Migrations are implemented using [golang-migrate](https://github.com/golang-migrate/migrate) utility.
For more detailed information check
[this](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md) tutorial.

#### Creating migrations

Project has its own default migration files. To create your own use the following commands

```shell
migrate create -ext "files' extension" -dir "dir with migration files" -seq "migration name"
```

#### Running migrations

To run migrations:

```shell
migrate -path "dir with migration files" -database "database URL" up
```

To reverse run migrations:

```shell
migrate -path "dir with migration files" -database "database URL" down
```

### Makefile commands

To use commands stated in the Makefile run the command below (without brackets):

```shell
make <make command name>
```

Here is a list of Makefile commands you are probably going to use during development.:

* `upload`

  > Build the docker image of the app corresponding to the current state of the files and push it into
  the DockerHub.
  Make sure that you are listed as a contributor in the repository on the DockerHub before trying to push.

* `up`

  > Launch app in docker containers using .env.docker file to set the environment.

* `down`

  > Stop docker containers
  