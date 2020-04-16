# FROM circleci/golang:1.13 AS build_base

# # Install some dependencies needed to build the project
# RUN apk add bash ca-certificates git gcc g++ libc-dev
# WORKDIR /go/src/github.com/rosspatil/go-kit-example

# # Force the go compiler to use modules
# ENV GO111MODULE=on

# # We want to populate the module cache based on the go.{mod} files.
# COPY go.mod .
# # COPY go.sum .

# RUN go mod download
# # This image builds the weavaite server
# FROM build_base AS binary_builder
# # Here we copy the rest of the source code
# COPY . .
# # And compile the project
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o go-kit-example

# #In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
# FROM alpine AS final
# # We add the certificates to be able to get access to internet
# RUN apk add ca-certificates
# # Finally we copy the statically compiled Go binary.
# COPY --from=binary_builder /go/src/github.com/rosspatil/go-kit-example ./go-kit-example

# # Expose the application on port BUILD_toko_PORT.
# EXPOSE 8080

# ENTRYPOINT ["./go-kit-example"]




FROM golang:1.12-alpine AS build-env
EXPOSE 8090 8091 8092
ENV CGO_ENABLED 0

ADD .  /go/src/github.com/rosspatil/go-kit-example
WORKDIR /go/src/github.com/rosspatil/go-kit-example
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

ENV GO111MODULE=on

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" 
RUN go mod download

ENV GO111MODULE=off
RUN go get github.com/derekparker/delve/cmd/dlv

COPY go-kit-example /go/src/github.com/rosspatil/go-kit-example

ENV GO111MODULE=on
WORKDIR /go/src/github.com/rosspatil/go-kit-example
ADD main.go .
CMD ["dlv", "debug", "github.com/rosspatil/go-kit-example","--headless", "--listen=:8092", "--api-version=2"]
