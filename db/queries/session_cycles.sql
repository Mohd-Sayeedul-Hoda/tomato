-- name: CreateSessionCycle :one
INSERT INTO session_cycles(session_id, type, start_time, status)
VALUES (?1, ?2, ?3, ?4)
RETURNING id;

-- name: GetSessionCycleByID :one
SELECT * FROM session_cycles
WHERE id = ?1;

-- name: GetSessionCyclesBySessionID :many
SELECT * FROM session_cycles
WHERE session_id = ?1
ORDER BY id;

-- name: UpdateSessionCycleStatus :exec
UPDATE session_cycles
SET status = ?2
WHERE id = ?1;

-- name: MarkSessionCycleComplete :exec
UPDATE session_cycles
SET status = ?2, end_time = ?3, duration = ?4
WHERE id = ?1;

-- name: DeleteSessionCycle :exec
DELETE FROM session_cycles
WHERE id = ?1;

-- name: GetSessionCyclesByType :many
SELECT * FROM session_cycles
WHERE type = ?1;

-- name: GetSessionCyclesByStatus :many
SELECT * FROM session_cycles
WHERE status = ?1;

-- name: MarkSessionCycleCompleted :exec
UPDATE session_cycles
SET status = 'completed', end_time = CURRENT_TIMESTAMP
WHERE id = ?1;

-- name: GetLatestSessionCycleByStatus :one
SELECT * FROM session_cycles
WHERE status = ?1
ORDER BY created_at DESC
LIMIT 1;

-- name: GetSessionCycleByStatusWithMetadata :many
SELECT
    sc.id,
    sc.session_id,
    sc.type,
    sc.created_at,
    sc.start_time,
    sc.end_time,
    sc.duration,
    sc.status,

    s.work_duration,
    s.break_duration,
    s.long_break_duration,
    s.long_break_cycle
FROM
    session_cycles AS sc
INNER JOIN
    sessions AS s ON sc.session_id = s.id
WHERE
    sc.status = ?1
