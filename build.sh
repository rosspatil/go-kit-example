GOOS=linux go build go-kit-example -ldflags "-s -w"
upx go-kit-example
docker build . --build-arg BUILDKIT_INLINE_CACHE=1   -t go-kit-example
docker run --rm -p 8090:8090 -p 8091:8091 -p 8092:8092 --security-opt=seccomp:unconfined --name=go-kit-example go-kit-example:latest