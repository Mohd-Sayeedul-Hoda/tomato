-- +goose Up
-- +goose StatementBegin
CREATE TABLE time_profiles(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE, -- like standard, hour work
  work_duration INTEGER NOT NULL,
  break_duration INTEGER NOT NULL,
  long_break_duration INTEGER NOT NULL,
  long_break_cycle INTEGER DEFAULT 4,
  is_default BOOLEAN DEFAULT false -- for choosing default one.
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS time_profiles;
-- +goose StatementEnd
