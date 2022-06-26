-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts
(
    id       BIGSERIAL PRIMARY KEY,
    document VARCHAR(11) NOT NULL,
    UNIQUE (document)
);

CREATE TABLE operation_types
(
    id          BIGSERIAL PRIMARY KEY,
    description VARCHAR NOT NULL,
    type        VARCHAR DEFAULT 'NEGATIVE'
);

CREATE TABLE transactions
(
    id                BIGSERIAL PRIMARY KEY,
    account_id        BIGSERIAL REFERENCES accounts (id)        NOT NULL,
    operation_type_id BIGSERIAL REFERENCES operation_types (id) NOT NULL,
    amount            FLOAT(4)                                  NOT NULL,
    event_date        TIMESTAMP                                 NOT NULL DEFAULT CURRENT_DATE
);

INSERT INTO operation_types (description)
VALUES ('COMPRA A VISTA');
INSERT INTO operation_types (description)
VALUES ('COMPRA PARCELADA');
INSERT INTO operation_types (description)
VALUES ('SAQUE');
INSERT INTO operation_types (description, type)
VALUES ('PAGAMENTO', 'POSITIVE');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
DROP TABLE operation_types;
DROP TABLE accounts;
-- +goose StatementEnd
