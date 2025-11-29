-- name: CreateTimeProfile :one
INSERT INTO time_profiles (
  name, 
  work_duration, 
  break_duration, 
  long_break_duration, 
  long_break_cycle,
  is_default
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetTimeProfile :one
SELECT * FROM time_profiles
WHERE id = ? LIMIT 1;

-- name: GetDefaultTimeProfile :one
SELECT * FROM time_profiles
WHERE is_default = true LIMIT 1;

-- name: ListTimeProfiles :many
SELECT * FROM time_profiles
ORDER BY name;

-- name: UpdateTimeProfile :one
UPDATE time_profiles
SET 
  name = ?,
  work_duration = ?,
  break_duration = ?,
  long_break_duration = ?,
  long_break_cycle = ?,
  is_default = ?
WHERE id = ?
RETURNING *;

-- name: DeleteTimeProfile :exec
DELETE FROM time_profiles
WHERE id = ?;
