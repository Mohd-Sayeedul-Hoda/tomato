-- +goose Up
-- +goose StatementBegin
INSERT INTO time_profiles 
  (name, work_duration, break_duration, long_break_duration, long_break_cycle, is_default)
VALUES
  ('Standard', 25, 5, 15, 4, TRUE),
  ('Hour Focus Simple', 52, 17, 17, 3, FALSE), 
  ('Hour Focus', 52, 17, 30, 3, FALSE), 
  ('Long Work', 80, 17, 50, 4, FALSE),
  ('Deep Work', 112, 26, 80, 2, FALSE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM time_profiles 
WHERE name IN ('Standard', 'Hour Focus', 'Long Work', 'Deep Work', 'Hour Focus Simple');
-- +goose StatementEnd
