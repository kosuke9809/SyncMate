FROM postgres:16.2-bullseye

WORKDIR /postgres

COPY ./db/sql/init.sql /docker-entrypoint-initdb.d/init.sql
