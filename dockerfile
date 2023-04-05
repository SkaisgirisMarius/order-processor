FROM mysql:latest

ENV MYSQL_DATABASE=mydatabase
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password

COPY ./schema.sql /docker-entrypoint-initdb.d/