FROM golang:1.18-alpine AS builder

LABEL maintainer="leoo.j <minkj1992@gmail.com> (https://minkj1992.github.io)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# https://stackoverflow.com/a/56103707
RUN apk add --no-cache ca-certificates openssl

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver ./application/cmd/main.go

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]