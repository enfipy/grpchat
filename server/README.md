# Server for grpchat

Golang grpchat server

## Usage:

To begin development:

```
docker-compose up --build
```

## Production build:

To build lightweight production image under `7 megabytes` just run:

```
docker build -f docker/prod.Dockerfile -t <image_name> .
```
