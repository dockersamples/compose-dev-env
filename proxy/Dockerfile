FROM nginx:1.21-alpine

RUN apk update \
    && apk add git

COPY conf /etc/nginx/conf.d/default.conf