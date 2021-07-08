FROM golang:1.16 AS build
WORKDIR /go/src/github.com/org/repo
COPY . .

RUN go build -o server .

FROM build AS development
RUN apt-get update \
    && apt install git
CMD ["go", "run", "main.go"]

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /go/src/github.com/org/repo/server /server
CMD ["/server"]
