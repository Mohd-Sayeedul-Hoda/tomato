-- +goose Up
-- +goose StatementBegin
INSERT INTO sessions 
  (label, status, is_tracked, note)
  VALUES
  ('Default', 'running', false, 'for default use');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM sessions WHERE label = 'Default'
-- +goose StatementEnd
