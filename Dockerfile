# build step
FROM golang:1.12 as builder

LABEL maintainer="sunnydog0826@gmail.com"

COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    && go build -v -a -tags netgo -o drone-docker ./cmd/drone-docker

# run step
FROM docker:18.09-dind

# copy bin from build step
COPY --from=builder /build/drone-docker /bin/

ENTRYPOINT ["/usr/local/bin/dockerd-entrypoint.sh", "/bin/drone-docker"]
