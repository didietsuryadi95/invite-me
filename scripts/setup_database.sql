CREATE SCHEMA IF NOT EXISTS internal;

-- create user for internal server and set up its privileges

DO
$body$
BEGIN
    IF NOT EXISTS (
        SELECT *
        FROM   pg_catalog.pg_user
        WHERE  usename = 'app_server') THEN

        CREATE USER app_server
        WITH NOSUPERUSER NOCREATEDB NOCREATEROLE
        ENCRYPTED PASSWORD 'password';
    END IF;
END
$body$;

GRANT ALL PRIVILEGES ON SCHEMA internal TO app_server;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA internal TO app_server;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA internal TO app_server;
