package sqlite

const sqlUserTable = `
CREATE TABLE IF NOT EXISTS user (
  id          INTEGER PRIMARY KEY,
  name        TEXT    NOT NULL,
  passphrase  TEXT    NOT NULL,
  first_name  TEXT    NOT NULL,
  last_name   TEXT,
  joined_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_at DATETIME,
  deleted_at  DATETIME
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name
  ON user (name);
`
