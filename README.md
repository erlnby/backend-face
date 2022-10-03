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

