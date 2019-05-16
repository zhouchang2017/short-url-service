# short-url-service


This is the short url service

online
> http://t.wewee.com/


## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- DB_USERNAME: database username
- DB_PASSWORD: database password
- DB_HOST: database host
- DB_PORT: database port
- DB_NAME: database name

## Dependencies

client build with vue

```

cd client

# install dependencies
yarn

# run watch
yarn watch

# run production
yarn prod
```

## Usage

A Makefile is included for convenience

Build the production

```
make build
```

Run the dev
```
DB_USERNAME=root -e DB_PASSWORD=12345678 -e DB_HOST=127.0.0.1 -e DB_PORT=3306 -e DB_NAME=micro_book_mall go run main.go
```
