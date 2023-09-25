FROM golang:1.21-alpine
WORKDIR /go/src/project
COPY tools tools
RUN tools/install.sh
COPY app app
COPY go.mod ./
RUN go work init && go work use app
CMD sleep infinity
