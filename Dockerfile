FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/sergiorra/hexagonal-arch-api-go
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/mooc-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/mooc-api /go/bin/mooc-api
ENTRYPOINT ["/go/bin/mooc-api"]