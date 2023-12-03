################## builder ##################
ARG BUILDER_DOCKER_REGISTRY="docker.io/library"
FROM $BUILDER_DOCKER_REGISTRY/golang:1.21.4-alpine AS builder
SHELL ["/bin/sh", "-c"]
WORKDIR /go/src/project
USER 0

ARG GOPROXY="https://goproxy.io"
ENV GOPROXY=$GOPROXY

COPY app app
COPY go.mod go.sum ./
RUN go build -o /go/bin/app ./app

################## final ##################
ARG DIST_DOCKER_REGISTRY="https://gcr.io/distroless"
FROM gcr.io/distroless/static-debian11 AS final
COPY --from=builder /go/bin/app /
USER 1001
CMD ["/app"]
