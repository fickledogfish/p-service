FROM dpage/pgadmin4:5.7

USER pgadmin

COPY --chown=pgadmin:pgadmin docker/postgres_admin/servers.json /pgadmin4/servers.json
COPY --chown=pgadmin:pgadmin SQL/ /var/lib/pgadmin/storage/admin_example.com/
