-- Filename: 001_create_tables.sql

CREATE TABLE request (
    id          UUID PRIMARY KEY,
    header      JSONB
);

CREATE TABLE response (
    id          UUID PRIMARY KEY,
    status      INTEGER,
    header      JSONB,
    length      INTEGER
);
