ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

# prefetch the binaries, so that they will be cached and not downloaded on each change
RUN go run github.com/steebchen/prisma-client-go prefetch

# generate the Prisma Client Go client
RUN go run github.com/steebchen/prisma-client-go generate

RUN go build -v -o /usr/src/app/run-app .

FROM debian:bookworm

COPY --from=builder /usr/src/app/run-app /usr/local/bin/
CMD ["run-app"]
