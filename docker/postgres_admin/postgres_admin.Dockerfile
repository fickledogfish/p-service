FROM dpage/pgadmin4:6.14

USER pgadmin

COPY --chown=pgadmin:pgadmin docker/postgres_admin/servers.json /pgadmin4/servers.json
COPY --chown=pgadmin:pgadmin SQL/ /var/lib/pgadmin/storage/admin_example.com/
