# Builds the CHACONTAINER API server. Build context must be the repository
# root (see docker-compose.yml: context: .., dockerfile: chacontainer/Dockerfile)
# because this binary lives inside the larger github.com/cli/cli/v2 module
# and needs its go.mod/go.sum.
FROM golang:1.25-alpine AS build
WORKDIR /src

# GOTOOLCHAIN=auto lets Go fetch the exact toolchain pinned in go.mod
# (go.mod may require a newer patch release than this base image ships).
ENV GOTOOLCHAIN=auto CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -o /out/chacontainer-server ./chacontainer/cmd/server

FROM alpine:3.20
RUN apk add --no-cache ca-certificates && adduser -D -u 10001 chacontainer
COPY --from=build /out/chacontainer-server /usr/local/bin/chacontainer-server
USER chacontainer
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/chacontainer-server"]
