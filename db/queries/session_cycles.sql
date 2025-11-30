-- name: CreateSessionCycle :one
INSERT INTO session_cycles(session_id, type, start_time, status, timer_profile_id)
VALUES (?1, ?2, ?3, ?4, ?5)
RETURNING id;

-- name: GetSessionCycleByID :one
SELECT * FROM session_cycles
WHERE id = ?1;

-- name: ListSessionCycles :many
SELECT * FROM session_cycles
WHERE
  -- Filter by session_id: If NULL, ignore
  (@session_id IS NULL OR session_id = @session_id)
  AND
  -- Filter by status: If NULL, ignore
  (@status IS NULL OR status = @status)
  AND
  -- Filter by type: If NULL, ignore
  (@type IS NULL OR type = @type)
ORDER BY created_at DESC
LIMIT ?;

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

-- name: MarkSessionCycleCompleted :exec
UPDATE session_cycles
SET status = 'completed', end_time = CURRENT_TIMESTAMP
WHERE id = ?1;

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
    tp.work_duration as work_duration,
    tp.break_duration as break_duration,
    tp.long_break_duration as long_break_duration,
    tp.long_break_cycle as long_break_cycle
FROM
    session_cycles AS sc
INNER JOIN
    sessions AS s ON sc.session_id = s.id
LEFT JOIN
    time_profiles AS tp ON sc.timer_profile_id = tp.id
WHERE
    sc.status = ?1;
