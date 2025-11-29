-- +goose Up
-- +goose StatementBegin
CREATE TABLE session_cycles(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
  timer_profile_id INTEGER NOT NULL REFERENCES time_profiles(id),
  type TEXT CHECK(type IN ('work', 'break', 'long_break')),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  start_time DATETIME,
  end_time DATETIME,
  duration INTEGER,
  status TEXT CHECK(status IN ('completed', 'running', 'skipped', 'cancelled', 'paused' ))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS session_cycles;
-- +goose StatementEnd

