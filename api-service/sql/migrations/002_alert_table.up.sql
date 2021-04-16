BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS alerts
(
    id          UUID    NOT NULL DEFAULT uuid_generate_v4(),
    from_symbol varchar NOT NULL,
    to_symbol   varchar NOT NULL,
    price       numeric NOT NULL,
    user_id     UUID    NOT NULL references users (id),

    PRIMARY KEY (id)
);

COMMIT;