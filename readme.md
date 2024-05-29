# Gaia-X Interconnect API
This is a RESTful API that allows the creation and management of VPNs for cloud interconnection. 

## Table of Contents

- [Features](#Features)
- [Directory Structure](#Directory-Structure)
- [Description](#Description)
- [Setup](#Setup)
- [Directory Structure Explanation](#Directory-Structure-Explanation)
- [License](#license)

## Features

- Standard responses for success and fail requests
- Swagger API documentation
- Sqlx DB with Postgres - but can be changed as needed.
- Standard for custom errors
- Logger for console and external file.
- Migrations setup
- Hot Reload
- Docker setup
- Intuitive, clean and scalabe structure
---

## Directory Structure
```
- /cmd --> Contains the app's entry-points 
  |- /server
     |- /docs
     |- main.go
     |- Makefile
  |- /another_binary
- /config --> Contains the config structures that the server uses.
- internal --> Contains the app's code
   |- /errors
   |- /handlers
   |- /middleware
   |- /model
   |- /storage
   |- server.go
- /logs --> The folder's files are not in version control. Here you'll have the logs of the apps (the ones you specify to be external)
- /migrations --> Migrations to set up the database schema in your db.
- /pkg --> Packages used in /internal
   |- /httputils
   |- /logger
- .air.toml
- .env --> Not in version control. Need to create your own - see below.
- .gitignore
- docker-compose.yml
- Dockerfile
- go.mod
- go.sum
- LICENSE
- README.md
```

## Setup

Make sure to first install the binaries that will generate the api docs and hot-reload the app.

```
go install github.com/swaggo/swag/cmd/swag@latest
```
and
```
go install github.com/cosmtrek/air@latest
```

Download the libs
```
go mod download
```
```
go mod tidy
```

Create an `.env` file in the root folder and use this template:
```
# DEV, STAGE, PROD
ENV=DEV
PORT=8080
VERSION=0.0.1

DB_HOST=localhost  #when running the app without docker
# DB_HOST=postgres # when running the app in docker
DB_USER=postgres
DB_NAME=postgres
DB_PASSWORD=postgres
DB_PORT=5432
```

If you start the app locally without docker make sure your Postgres DB is up.
Write `air` in terminal and the app will be up and running listening on whatever port you set in .env.

Don't forget to rename the module and all the imports in the code from my github link to yours.

## Setup with Docker

To run the app in docker, you need to have docker installed on your machine.

Create an `.env` file in the root folder and use this template:
```
# DEV, STAGE, PROD
ENV=DEV
PORT=8080
VERSION=0.0.1

# DB_HOST=localhost  #when running the app without docker
DB_HOST=postgres # when running the app in docker
DB_USER=postgres
DB_NAME=postgres
DB_PASSWORD=postgres
DB_PORT=5432
```

Then run the following command to build the image and start the container:
```
docker-compose up --build
```

## Cleaning up

To stop the app running in docker, run the following command:
```
docker-compose down
```

## Build API Docs

Navigate to the `cmd/server` directory and run the following command:
```
make swag
```

## Build the App

Navigate to the `cmd/server` directory and run the following command:
```
make build
```

## Testing

To run the tests, navigate to the `cmd/server` directory and run the following command:
```
make test
```

## Accessing the API Docs

After the app is up and running, you can access the API docs by going to `http://localhost:8080/api/docs/index.html`.

## Directory Structure Explanation

* /cmd
   - The entry point of the server is in `main.go`, where the app loads environment variables and config from `/config`.
   - The config is passed to the server and can be expanded as needed.
   - A goroutine runs an `OnShutdown` function for server crashes or panics.
   - A `Makefile` is included for building the app.

* /internal
   - Contains the main code.
   - `server.go` defines the `AppServer` struct with db, handlers, and server setup.
   - `Run` function configures the server, database, router, middlewares, handlers, and migrations.
   - `Sender` and `Storage` are injected from `handlers.Handlers`.
   - Functions: `OnShutdown`, `NotFoundHandler` (404), and `NotAllowedHandler` (405).

* /errors
   - Defines custom error structs for response handling.
   - The `Sender` formats responses based on these errors.
   - Accepts errors as strings, structs, or custom `Err` structs.

* /handlers
   - `handlers.go` sets up the handlers object with its dependencies.
   - Specific handler groups (e.g., `books.go`, `users.go`) are added here.
   - Functions in these files are received by the `Handlers` struct.

* /middlewares
   - Contains custom middlewares added to `negroni` in `server.go`.

* /model
   - `base_model.go` serves as the base for other models.
   - Each file (e.g., `books.go`) defines database models and request/response structs.

* /storage
   - `storage.go` defines the storage interface for database switching.
   - `postgres_db.go` creates a PostgreSQL storage instance.
   - New databases can be added by implementing the `StorageInterface`.

* /pkg
   - `httputils` contains `http_response.go` for response standards.
   - The `Sender` struct handles automated responses (e.g., JSON).
   - `logger` provides logging to console or `/logs` directory, configurable to a different path (e.g., `/var/log/myapp` for Unix/Linux).


# Sample API

Create a new VPN:
```sh
curl -X 'POST'   'http://0.0.0.0:8080/api/vpn/add'   -H 'accept: application/json'   -H 'Content-Type: application/json'   -d '{"type": "vpn", "local_as_number": 65000, "remote_as_number": 65001, "VNI": 1000, "name": "vpn1"}'
```

## License

This project is licensed under the AGPLv3 License - see the [LICENSE](./LICENSE) file for details.

