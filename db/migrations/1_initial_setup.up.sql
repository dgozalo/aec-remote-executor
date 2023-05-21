CREATE SCHEMA IF NOT EXISTS aec_executor;

CREATE TABLE aec_executor.accounts (
                          username VARCHAR ( 50 ) UNIQUE NOT NULL,
                          password VARCHAR ( 50 ) NOT NULL,
                          email VARCHAR ( 255 ) UNIQUE NOT NULL,
                          created_on TIMESTAMP NOT NULL,
                          last_login TIMESTAMP
);