FROM golang:1.13-alpine3.10 as build

COPY . /go/src/github.com/kevin-monteiro/http-stub-provider

WORKDIR /go/src/github.com/kevin-monteiro/http-stub-provider

RUN go build -o http-stub-provider cmd/main.go

FROM alpine:3.10 AS release

COPY --from=build /go/src/github.com/kevin-monteiro/http-stub-provider/http-stub-provider /go/bin/http-stub-provider

ENTRYPOINT ["/go/bin/http-stub-provider"]