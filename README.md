# Api-Audit

Auditācijas API

## Built with Azugo Go Web Framework

This project is built using the [Azugo Go Web Framework](https://azugo.io), a powerful and flexible framework for building modern web applications in Go. Check out the [Azugo GitHub page](https://github.com/azugo) for more information and documentation.


<!-- TOC -->

- [Api-Audit](#repo_name_title)
  - [Development](#development)
    - [Prepare dependecies](#prepare-dependecies)
    - [Local development](#local-development)
    - [Before commit](#before-commit)
  - [Environment variables](#environment-variables)
    - [Local example](#local-example)

<!-- /TOC -->

## Development

### Prepare dependecies

```sh
go mod download
go generate ./...
```

### Local development

To build in VS Code use `Ctrl`+`Shift`+`B`.

To debug project in VS Code use `F5`.

### Before commit

> CI requires linted, formatted code

You should run:

```sh
gofmt -s -w ./..
```

or

```sh
gofumpt -w ./..
```

and fix any errors reported by

```sh
golangci-lint run
```

## Environment variables

In order to run the service you need configure environment variables. List of environment variables:

| Variable | Description | Default value | Required |
| --- | --- | --- | --- |
| `SERVER_URLS` | An server URL or multiple URLS separated by semicolon to listen on. | 0.0.0.0:8080 | Yes |
| `ENVIRONMENT` | Environment name. Possible values: `Development`, `Staging`, `Production` | `Development` | Yes |
| `BASE_PATH` | Base path for all routes | `/` (or take value from `SERVER_URLS` path if exists) | No |
| `ACCESS_LOG_ENABLED` | Enable access log | `true` | Yes |
| `REVERSE_PROXY_LIMIT` | Limit for reverse proxy. | `1` | No |
| `REVERSE_PROXY_TRUSTED_IPS` | List of trusted IP addresses for reverse proxy. Separated by `;` | `"127.0.0.1"` | No |
| `REVERSE_PROXY_TRUSTED_HEADERS` | List of trusted headers for reverse proxy. Separated by `;` | `X-Real-IP; X-Forwarded-For` | No |
| `LOG_LEVEL` | Minimal log level. Allowed values are `debug`, `info`, `warn`, `error`, `fatal`, `panic` | `info` | Yes |
| `POSTGRES_HOST` | Database host name or URL | | Yes |
| `POSTGRES_USER` | Database user name | | Yes |
| `POSTGRES_DB` | Database name | | Yes |
| `POSTGRES_PASSWORD_FILE` | Database connection password | | Yes |
| `IDAUTH_URL` | IDAuth API URL | | Yes |
| `IDAUTH_CLIENT_ID` | IDAuth client ID | | Yes |
| `IDAUTH_CLIENT_SECRET` | IDAuth client secret | | Yes |

### Local example

In local development you must create `.env` file in the root of the project. Example:

```sh
ENVIRONMENT=Development
BASE_PATH=/
LOG_LEVEL=debug

IDAUTH_URL=<idauth_URL>
IDAUTH_CLIENT_ID=edim.api-audit
IDAUTH_CLIENT_SECRET=<idauth_client_secret>

POSTGRES_HOST=<database_host_name>
POSTGRES_USER=<database_user_name>
POSTGRES_DB=edim
POSTGRES_PASSWORD=<database_password>
```
