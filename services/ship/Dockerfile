# BUILD STAGE
FROM golang:1.12.5 as build

WORKDIR /go/src/app

COPY ${PWD} /go/src/app

RUN /go/src/app/compile.sh

# RUNTIME STAGE
FROM alpine:3 as runtime

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/app/bin/ship /

COPY resources /resources

USER nobody

CMD ["/ship"]
