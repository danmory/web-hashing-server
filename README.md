# Web Hashing Server

## Description

The application gets the URLs and

## Requirements

1. GO 1.19
2. Installed libraries from *go.mod*
3. (Optional) Docker compose

## Usage

1. Clone the repository
2. Install dependencies

    ``$ go mod tidy``

3. Run Postgres

4. Initilize *.env* file with settings specified in *.env.example*

5. Run the application

    DEBUG mode - ``$ go run .``

    PRODUCTION mode - ``$ GIN_MODE=release go run .``

    By default the application stores data in-memory.

    To store it in the database run the command with flag *-d*

    `` $ go run . -d ``

## Usage with Docker Compose

In progress...

## Tests

In progress...

## Contacts

Danila Moriakov (d.moriakov@gmail.com)
