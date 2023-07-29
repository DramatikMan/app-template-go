FROM golang:alpine
WORKDIR /go/src/project
COPY tools tools
RUN cd tools && ./install.sh
COPY go.mod go.work ./
CMD sleep infinity
