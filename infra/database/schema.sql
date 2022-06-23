CREATE TABLE accounts
(
    id       BIGSERIAL PRIMARY KEY,
    document VARCHAR (11) NOT NULL,
    UNIQUE (document)
);

CREATE TABLE operation_types
(
    id          BIGSERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE transactions
(
    id                BIGSERIAL PRIMARY KEY,
    account_id        BIGSERIAL REFERENCES accounts (id)        NOT NULL,
    operation_type_id BIGSERIAL REFERENCES operation_types (id) NOT NULL,
    amount            FLOAT(4)                                  NOT NULL,
    event_date        DATE                                      NOT NULL DEFAULT CURRENT_DATE
);