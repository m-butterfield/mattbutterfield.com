CREATE TABLE images (
    id character varying(255) PRIMARY KEY,
    caption text,
    location text,
    width integer NOT NULL,
    height integer NOT NULL
);
