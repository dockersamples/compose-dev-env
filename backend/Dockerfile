FROM golang:1.18 AS build
WORKDIR /go/src/aquafarm-management
COPY . .

ENV CGO_ENABLED=0

# Mendefinisikan environment variable DB_HOST, DB_PASSWORD, DB_USER, dan DB_NAME
ENV DB_HOST=$DB_HOST
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_USER=$DB_USER
ENV DB_NAME=$DB_NAME

RUN go build -o server .

FROM build AS development
RUN apt-get update \
    && apt-get install -y git
CMD ["go", "run", "main.go"]

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /go/src/aquafarm-management/server /server
CMD ["/server"]
