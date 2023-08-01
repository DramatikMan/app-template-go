FROM golang:alpine
WORKDIR /go/src/project
COPY tools tools
RUN tools/install.sh
COPY go.mod go.work ./
CMD sleep infinity
