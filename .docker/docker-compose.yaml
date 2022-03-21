version: "3.7"
services:
  backend:
    build:
      context: backend
      target: development
    secrets:
      - db-password
    depends_on:
      - db
  db:
    image: mariadb
    restart: always
    healthcheck:
      test: [ "CMD", "mysqladmin", "ping", "-h", "127.0.0.1", "--silent" ]
      interval: 3s
      retries: 5
      start_period: 30s
    secrets:
      - db-password
    volumes:
      - db-data:/var/lib/mysql
    environment:
      - MYSQL_DATABASE=example
      - MYSQL_ROOT_PASSWORD_FILE=/run/secrets/db-password
    expose:
      - 3306
  proxy:
    build: proxy
    ports:
      - 8080:80
    depends_on:
      - backend
volumes:
  db-data:
secrets:
  db-password:
    file: db/password.txt
