FROM postgres

ENV POSTGRES_DB todo
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password

COPY psql_dump.sql /docker-entrypoint-initdb.d/