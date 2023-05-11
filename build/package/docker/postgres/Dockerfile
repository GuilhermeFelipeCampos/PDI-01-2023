FROM postgres:15

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=go-kafka-db

ARG DB_SQL=postgres/database/squemas/squema_db.sql
ARG TABLE_SQL=postgres/database/squemas/squema_table.sql

COPY $DB_SQL /docker-entrypoint-initdb.d/
COPY $TABLE_SQL /docker-entrypoint-initdb.d/

EXPOSE 5432:5432

CMD ["postgres"]