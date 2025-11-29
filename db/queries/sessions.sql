-- name: CreateSession :one
INSERT INTO sessions (label, note, status, session_estimate, is_tracked, start_time, work_duration, break_duration, long_break_duration, long_break_cycle)
VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9, ?10)
RETURNING id;

-- name: GetSessionById :one
SELECT * FROM sessions WHERE id = ?1;

-- name: GetActiveSessions :many
SELECT * FROM sessions WHERE status = 'running';

-- name: UpdateSessionStatus :exec
UPDATE sessions SET status = ?2 WHERE id = ?1;

-- name: UpdateSessionEndTime :exec
UPDATE sessions SET end_time = ?2 WHERE id = ?1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = ?1;

-- name: GetAllSessions :many
SELECT * FROM sessions ORDER BY created_at DESC;

-- name: GetCompletedSessions :many
SELECT * FROM sessions WHERE status = 'completed' ORDER BY end_time DESC;

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
    is_tracked = ?6,
    start_time = ?7,
    end_time = ?8,
    work_duration = ?9,
    break_duration = ?10,
    long_break_duration = ?11,
    long_break_cycle = ?12
WHERE id = ?1;

-- name: MarkSessionCompleted :exec
UPDATE sessions
SET status = 'completed', end_time = CURRENT_TIMESTAMP
WHERE id = ?1;