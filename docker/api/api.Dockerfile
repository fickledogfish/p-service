ARG ENDPOINT

ARG EXPOSED_PORT

ARG SRC_DIR=/app/src
ARG SRC_DIR=/app/run

# First step: build the selected endpoint
FROM golang:latest AS api_builder

ARG ENDPOINT
ARG SRC_DIR

COPY . $SRC_DIR
WORKDIR $SRC_DIR

RUN go build -o api endpoints/$ENDPOINT.go

# Second step: run the api
FROM api_builder AS api_publisher

ARG SRC_DIR
ARG RUN_DIR
ARG EXPOSED_PORT

ENV EXPOSED_PORT $EXPOSED_PORT

RUN groupadd api && useradd -g api api

WORKDIR $RUN_DIR
COPY --from=api_builder --chown=api:api $SRC_DIR/api $RUN_DIR

USER api
ENTRYPOINT [ "./api" ]
