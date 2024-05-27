-- This migration is intended to be for Postgres. Make another one for your db if the SQL is not accepted by your db.
CREATE EXTENSION IF NOT EXISTS citext;

DO $$
BEGIN
    -- Check if the domain 'email' does not exist
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type 
        WHERE typname = 'email'
    ) THEN
        -- Create the domain if it doesn't exist
        CREATE DOMAIN email AS citext CHECK (
            value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$'
        );
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS vpn (
    id serial primary key,
    name varchar(255),
    type varchar(255),
    local_as_number int,
    remote_as_number int,
    vni int,
    created_at timestamp,
    updated_at timestamp,    
    CONSTRAINT vpn_id_unique UNIQUE (id)
);
