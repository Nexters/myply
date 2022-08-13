FROM golang:1.18-alpine AS builder

LABEL maintainer="leoo.j <minkj1992@gmail.com> (https://minkj1992.github.io)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# To handle: x509: certificate signed by unknown authority in youtube and mongo atlas
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates


RUN go build -ldflags="-s -w" -o apiserver ./application/cmd/main.go

FROM scratch
# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/"]
# refs: https://gist.github.com/michaelboke/564bf96f7331f35f1716b59984befc50
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]