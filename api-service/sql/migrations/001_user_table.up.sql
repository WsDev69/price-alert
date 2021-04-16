BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id         UUID      NOT NULL DEFAULT uuid_generate_v4(),
    email      text      NOT NULL UNIQUE,

    PRIMARY KEY (id)
);

COMMIT;