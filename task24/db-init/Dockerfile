FROM postgres:15.1

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=ec2_instances_db

COPY init.sql /docker-entrypoint-initdb.d/