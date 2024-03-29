FROM golang:1.20 AS builder

WORKDIR /app

ENV http_proxy=$http_proxy
ENV https_proxy=$https_proxy

ENV GOPROXY=https://goproxy.io,direct

COPY go.mod go.sum ./
RUN go mod download 

ENV GOCACHE=/go-build

COPY . .

RUN go env -w GO111MODULE=on && CGO_ENABLED=0 GOPROXY="https://goproxy.io" GOOS=linux GOARCH=amd64 go build -o main -tags=jsoniter -ldflags="-w -s" main.go

FROM ubuntu:latest AS prod
ARG PROJECT_NAME=pp_backend_2023_3

WORKDIR /opt/${PROJECT_NAME}

COPY --from=builder /app/main ./${PROJECT_NAME}

RUN echo "./${PROJECT_NAME}" >run.sh && chmod -R 755 /opt/${PROJECT_NAME}

EXPOSE 5000

CMD ./run.sh
