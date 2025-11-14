-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  label TEXT NOT NULL,
  work_duration INTEGER NOT NULL,
  break_duration INTEGER NOT NULL,
  long_break_duration INTEGER NOT NULL,
  long_break_cycle INTEGER DEFAULT 4,
  start_time DATETIME NOT NULL,
  end_time DATETIME,
  status TEXT NOT NULL CHECK(status IN ('running', 'completed', 'cancelled')),
  session_estimate INTEGER,
  is_tracked BOOLEAN DEFAULT FALSE,  -- false = just a timer, true = logged work
  note TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
