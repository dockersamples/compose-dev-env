FROM golang:1.16 AS build
WORKDIR /go/src/github.com/org/repo
COPY . .

ENV CGO_ENABLED=0
RUN go build -o server .

FROM build AS development
RUN apt-get update \
    && apt-get install -y git
CMD ["go", "run", "main.go"]

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /go/src/github.com/org/repo/server /server
CMD ["/server"]
