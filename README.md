# GitDesk Forwarder

GitDesk Forwarder is a tool that forwards GitLab Service Desk issues to messaging platforms. Currently, it supports Slack. The configuration is managed through a `config.yaml` file, allowing for customization of the bot username, port, and webhooks.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Docker](#docker)
- [Endpoints](#endpoints)
- [License](#license)

## Installation

To get started with GitDesk Forwarder, clone the repository and install the necessary dependencies.

```bash
git clone https://github.com/nilpntr/gitdesk-forwarder.git
go mod tidy
```

## Configuration

The configuration is done through a `config.yaml` file. Below is the structure of the configuration file:

```yaml
botUsername: support-bot # optional, defaults to "support-bot"
port: 8080               # optional, defaults to 8080
webhooks:
  - secretToken: "your_secret_token" # optional
    listenPath: "/your/listen/path"
    slackWebhookUrl: "https://hooks.slack.com/services/your/webhook/url" # optional
```

### Default Values

- **`botUsername`**: Optional. Defaults to `support-bot`.
- **`port`**: Optional. Defaults to `8080`.

### Config File Location

By default, the application looks for the configuration file at `/app/config.yaml`. You can specify a different path by using the `--config {path to file}` argument.

## Usage

To run GitDesk Forwarder, use the following command:

```bash
go run main.go --config {path to config.yaml} --log-level {info|debug|error}
```

### Log Levels

You can set the log level using the `--log-level` flag. The available log levels are:

- `info`
- `debug`
- `error`

## Docker

GitDesk Forwarder is also available as a Docker image. You can pull the image using the following command:

```bash
docker pull sammobach/gitdesk-forwarder
```

If you want to execute the command manually inside the Docker container, the binary is located at `/app/gitdesk-forwarder`.

## Endpoints

### Health Check

You can check the health of the service by sending a GET request to the following endpoint:

```
GET /health
```

This endpoint returns a `200 OK` status if the service is running properly.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.txt) file for details.