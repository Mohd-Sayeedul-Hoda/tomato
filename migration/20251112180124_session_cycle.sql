-- +goose Up
-- +goose StatementBegin
CREATE TABLE session_cycles(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
  type TEXT CHECK(type IN ('work', 'break', 'long_break')),
  start_time DATETIME,
  end_time DATETIME,
  duration INTEGER,
  status TEXT CHECK(status IN ('completed', 'skipped', 'cancelled'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS session_cycles;
-- +goose StatementEnd

