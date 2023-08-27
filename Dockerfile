FROM golang:alpine
WORKDIR /go/src/project
COPY tools tools
RUN tools/install.sh
COPY go.mod ./
CMD sleep infinity
