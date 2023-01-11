CREATE TABLE users (
  id         UUID NOT NULL PRIMARY KEY,
  name       TEXT NOT NULL,
  created_at TIMESTAMP NULL DEFAULT NOW()
);