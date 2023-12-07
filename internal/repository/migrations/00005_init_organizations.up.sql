CREATE TYPE user_role_new AS ENUM ('system_manager', 'admin', 'no_access');

CREATE TABLE users_new (
    id              TEXT          NOT NULL PRIMARY KEY,
    organization_id uuid,
    name            TEXT          NOT NULL,
    email           TEXT          NOT NULL,
    role            user_role_new NOT NULL,
    deleted         BOOLEAN       NOT NULL,
    version         INT           NOT NULL
);

INSERT INTO users_new (id, name, email, role, deleted, version)
SELECT identity_uid, '', '', 'no_access'::user_role_new, FALSE, -1
FROM users;

ALTER TABLE washes
ADD COLUMN new_owner_id TEXT;

UPDATE washes
SET new_owner_id = (SELECT identity_uid FROM users WHERE id = washes.owner_id);

ALTER TABLE washes
DROP CONSTRAINT washes_owner_fk,
DROP COLUMN owner_id;

ALTER TABLE washes
RENAME COLUMN new_owner_id TO owner_id;

ALTER TABLE washes
ADD CONSTRAINT washes_owner_fk FOREIGN KEY (owner_id) REFERENCES users_new (id);

DROP TABLE users;
ALTER TABLE users_new RENAME TO users;

DROP TYPE user_role;

ALTER TYPE user_role_new RENAME TO user_role;

CREATE TABLE organizations (
    id           UUID    PRIMARY KEY,
    name         TEXT    NOT NULL,
    display_name TEXT    NOT NULL,
    description  TEXT    NOT NULL,
    is_default   BOOLEAN NOT NULL,
    deleted      BOOLEAN NOT NULL, 
    version      INTEGER NOT NULL
);

CREATE TABLE wash_groups (
    id              UUID    PRIMARY KEY,
    organization_id UUID    NOT NULL REFERENCES organizations(id),
    name            TEXT    NOT NULL,
    description     TEXT    NOT NULL,
    is_default      BOOLEAN NOT NULL,
    deleted         BOOLEAN NOT NULL, 
    version         INTEGER NOT NULL
);

ALTER TABLE washes
    ADD COLUMN group_id UUID REFERENCES wash_groups(id);