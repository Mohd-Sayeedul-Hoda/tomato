-- name: CreateSession :one
INSERT INTO sessions (label, note, status, session_estimate, is_tracked)
VALUES (?1, ?2, ?3, ?4, ?5)
RETURNING id;

-- name: GetSessionById :one
SELECT * FROM sessions WHERE id = ?1;

-- name: GetActiveSessions :many
SELECT * FROM sessions WHERE status = 'running';

-- name: UpdateSessionStatus :exec
UPDATE sessions SET status = ?2 WHERE id = ?1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = ?1;

-- name: GetAllSessions :many
SELECT * FROM sessions ORDER BY created_at DESC;

-- name: GetCompletedSessions :many
SELECT * FROM sessions WHERE status = 'completed' ORDER BY updated_at DESC;

-- name: GetSessionsByTrackedStatus :many
SELECT * FROM sessions WHERE is_tracked = ?1 ORDER BY created_at DESC;

-- name: UpdateSessionNote :exec
UPDATE sessions SET note = ?2 WHERE id = ?1;

-- name: GetSessionsForDate :many
SELECT * FROM sessions WHERE DATE(created_at) = ?1;

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