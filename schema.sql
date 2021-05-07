CREATE TABLE images (
  id character varying(255) PRIMARY KEY,
  caption text,
  location text,
  width integer NOT NULL,
  height integer NOT NULL
);

CREATE TABLE songs (
  id character varying(255) PRIMARY KEY,
  description text,
  created_at timestamp NOT NULL
);
