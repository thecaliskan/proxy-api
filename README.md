# Proxy API

Proxy API for handling POST form data requests.

## Overview

This project provides a simple proxy API for POST form data requests. It's built with Go and includes Docker support for easy deployment.

## Features

- Proxy POST form data requests.
- Built with Go.
- Docker support.

## Requirements

- Go 1.22+
- Docker (optional)

## Usage

1. Docker:
    ```sh
   docker run -d 9090:80 --name proxy-api ghcr.io/thecaliskan/proxy-api
    ```

2. Binary:
    Download [release](https://github.com/thecaliskan/proxy-api/releases), extract and run command on cli
    ```sh
   ./proxy-api
    ```

## Configuration

Configuration options can be set via environment variables:

- PROXY_API_PORT: The port on which the server will run (default: 9900, docker: 80).

## Example
```sh
    curl --location 'localhost:9900' \
        --header 'proxy-url: https://httpbin.org/post' \
        --header 'Content-Type: application/x-www-form-urlencoded' \
        --data-urlencode 'john=doe' \
        --data-urlencode 'foo=bar'
```
## Development

1. Clone the repository:
    ```sh
    git clone https://github.com/thecaliskan/proxy-api.git
    cd proxy-api
    ```

2. Build the application:
    ```sh
    go build -o proxy-api
    ```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any bugs or feature requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.