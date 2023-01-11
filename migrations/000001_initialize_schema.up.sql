CREATE TABLE users (
  id         UUID NOT NULL PRIMARY KEY,
  username   TEXT NOT NULL,
  created_at TIMESTAMP NULL DEFAULT NOW()
);