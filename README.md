# Web Hashing Server

## Description

The application saves URL under its hash value and
returns this URL while passing hash back.

## Requirements

1. GO 1.19
2. Installed libraries from *go.mod*
3. (Optional) Docker, Docker Compose

## Usage

1. Clone the repository
2. Install dependencies

    ``$ go mod download``

3. Run Postgres

4. Initilize *.env* file with settings specified in *.env.example*

5. Run the application

    DEBUG mode - ``$ go run .``

    PRODUCTION mode - ``$ GIN_MODE=release go run .``

    By default the application stores data in-memory.

    To store it in the database run the command with flag *-d*

    `` $ go run . -d ``

6. Access the application on <http://127.0.0.1:8080>

## Usage with Docker, Docker Compose

1. Initilize *.env* file with settings specified in *.env.example*

2. Start the application

    `` $ docker-compose up -d ``

3. Access the application on <http://127.0.0.1:8080>

## Contacts

Danila Moriakov (d.moriakov@gmail.com)
