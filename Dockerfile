FROM golang
SHELL ["/bin/bash", "-c"]
WORKDIR /go/src/project
RUN go mod init
CMD sleep infinity