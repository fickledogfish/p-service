FROM postgres:14.5

USER postgres

COPY --chown=postgres:postgres ./SQL/*-up-*.sql /docker-entrypoint-initdb.d/
