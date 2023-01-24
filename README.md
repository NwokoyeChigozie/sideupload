# upload-ms

### Prerequisites

1. **Go 1.17** or **lastest version** already installed on your local machine.

### Run Project from Root

1. Create and populate a `app.env` file on the project root with its keys and corresponding values as listed in `app-sample.env`
2. Run from project root directory

```bash
$ go run main.go
```

### Run Project as Docker container

1. Create and populate a `app.env` file on the project root with its keys and corresponding values as listed in `app-sample.env`
2. Build Docker Image by running this command on the project root

```bash
$ docker build -t <image tag> .
```

4. Run image built with above command with

```bash
$ docker run -d -p <port>:<port> <image tag>
```

e.g docker run -d -p 8011:8011 vesicash-upload-ms

### Testing

1. Automated unit and integration tests done with golang's builtin [`testing`](https://pkg.go.dev/testing) package.

To run one test file:

```bash
$ go test -v  ./tests/<file name> -timeout 99999s
```

To run all tests:

```bash
$ go test -v  ./tests/<folder name>/<file name> -timeout 99999s
```

```bash
$ go test -v  ./tests/... -timeout 99999s
```

NB: Always add timeout tag to prevent early timeout
