FROM golang:1.14 AS builder
WORKDIR /usr/src/app
env GOPROXY=https://proxy.golang.org
env GOPRIVATE=git.lifemiles.net
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . /usr/src/app/
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o app .
RUN go test ./... -coverprofile=coverage.out

FROM nexus-release.lifemiles.net/alpine:3.11.3 AS deploy
WORKDIR /home/lifemiles/
add ./application.yaml /home/lifemiles/properties/application.yaml
run ls properties/
RUN apk add --no-cache tzdata
ENV TZ America/El_Salvador
COPY --from=builder /usr/src/app/app .
CMD ./app --base_path /home/lifemiles/properties/ --env_path /home/lifemiles/profiles/ --active_env ${ENVIRONMENT}

