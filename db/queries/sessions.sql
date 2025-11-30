-- name: CreateSession :one
INSERT INTO sessions (label, note, status, session_estimate, is_tracked)
VALUES (?1, ?2, ?3, ?4, ?5)
RETURNING id;

-- name: GetSessionById :one
SELECT * FROM sessions WHERE id = ?1;

-- name: ListSessions :many
SELECT
  id,
  label,
  status,
  session_estimate,
  is_tracked,
  note,
  created_at,
  updated_at
FROM
  sessions
WHERE
  (@status IS NULL OR status = @status)
  AND
  (@date IS NULL OR date(created_at) = @date)
  AND
  (@is_tracked IS NULL OR is_tracked = @is_tracked)
ORDER BY
  created_at DESC;

-- name: UpdateSessionStatus :exec
UPDATE sessions SET status = ?2 WHERE id = ?1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = ?1;

-- name: UpdateSessionNote :exec
UPDATE sessions SET note = ?2 WHERE id = ?1;

-- name: UpdateSession :exec
UPDATE sessions
SET label = ?2,
    note = ?3,
    status = ?4,
    session_estimate = ?5,
    is_tracked = ?6
WHERE id = ?1;

-- name: MarkSessionCompleted :exec
UPDATE sessions
SET status = 'completed'
WHERE id = ?1;
