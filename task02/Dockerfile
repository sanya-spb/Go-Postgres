FROM postgres:13.4
RUN apt update && apt install -y python3-pip postgresql-plpython3-13 pgxnclient
RUN pip3 install faker
RUN pgxn install postgresql_faker

# Variables needed at runtime to configure postgres and run the initdb scripts
ENV POSTGRES_DB 'sauna'
ENV POSTGRES_USER 'postgres'
ENV POSTGRES_PASSWORD 'passwd'
ENV PGDATA '/var/lib/postgresql/data'

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/docker-entrypoint-initdb.d", "/var/log/postgresql", "/var/lib/postgresql"]

EXPOSE 5432