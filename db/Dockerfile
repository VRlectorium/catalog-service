FROM postgres

ADD create.sql /docker-entrypoint-initdb.d/10-init.sql
COPY createTables.sql /docker-entrypoint-initdb.d/20-init.sql