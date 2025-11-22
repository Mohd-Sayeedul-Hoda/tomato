package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/sqlc"
)

type SessionCycleRepo struct {
	queries *sqlc.Queries
}

func NewSessionCycleRepository(db *sql.DB) *SessionCycleRepo {
	return &SessionCycleRepo{
		queries: sqlc.New(db),
	}
}

func (s *SessionCycleRepo) CreateSessionCycle(ctx context.Context, cycle models.SessionCycle) (int64, error) {
	params := sqlc.CreateSessionCycleParams{
		SessionID: cycle.SessionID,
		Type:      sql.NullString{String: "", Valid: false},
		StartTime: sql.NullTime{Time: time.Time{}, Valid: false},
		Status:    sql.NullString{String: "", Valid: false},
	}

	if cycle.Type != nil {
		params.Type = sql.NullString{String: *cycle.Type, Valid: true}
	}
	if cycle.StartTime != nil {
		params.StartTime = sql.NullTime{Time: *cycle.StartTime, Valid: true}
	}
	if cycle.Status != nil {
		params.Status = sql.NullString{String: *cycle.Status, Valid: true}
	}

	id, err := s.queries.CreateSessionCycle(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SessionCycleRepo) GetSessionCycleByID(ctx context.Context, id int64) (models.SessionCycle, error) {
	cycle, err := s.queries.GetSessionCycleByID(ctx, id)
	if err != nil {
		return models.SessionCycle{}, err
	}
	return s.mapSQLCSessionCycleToModel(cycle), nil
}

func (s *SessionCycleRepo) GetSessionCyclesBySessionID(ctx context.Context, sessionID int64) ([]models.SessionCycle, error) {
	cycles, err := s.queries.GetSessionCyclesBySessionID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	result := make([]models.SessionCycle, len(cycles))
	for i, cycle := range cycles {
		result[i] = s.mapSQLCSessionCycleToModel(cycle)
	}
	return result, nil
}

func (s *SessionCycleRepo) GetSessionCyclesByStatus(ctx context.Context, status string) ([]models.SessionCycle, error) {
	sqlNullString := sql.NullString{String: status, Valid: true}
	cycles, err := s.queries.GetSessionCyclesByStatus(ctx, sqlNullString)
	if err != nil {
		return nil, err
	}

	result := make([]models.SessionCycle, len(cycles))
	for i, cycle := range cycles {
		result[i] = s.mapSQLCSessionCycleToModel(cycle)
	}
	return result, nil
}

func (s *SessionCycleRepo) GetSessionCyclesByType(ctx context.Context, cycleType string) ([]models.SessionCycle, error) {
	sqlNullString := sql.NullString{String: cycleType, Valid: true}
	cycles, err := s.queries.GetSessionCyclesByType(ctx, sqlNullString)
	if err != nil {
		return nil, err
	}

	result := make([]models.SessionCycle, len(cycles))
	for i, cycle := range cycles {
		result[i] = s.mapSQLCSessionCycleToModel(cycle)
	}
	return result, nil
}

func (s *SessionCycleRepo) UpdateSessionCycleStatus(ctx context.Context, id int64, status string) error {
	params := sqlc.UpdateSessionCycleStatusParams{
		ID:     id,
		Status: sql.NullString{String: status, Valid: true},
	}
	return s.queries.UpdateSessionCycleStatus(ctx, params)
}

func (s *SessionCycleRepo) MarkSessionCycleComplete(ctx context.Context, id int64, status string, endTime time.Time, duration int64) error {
	params := sqlc.MarkSessionCycleCompleteParams{
		ID:       id,
		Status:   sql.NullString{String: status, Valid: true},
		EndTime:  sql.NullTime{Time: endTime, Valid: true},
		Duration: sql.NullInt64{Int64: duration, Valid: true},
	}
	return s.queries.MarkSessionCycleComplete(ctx, params)
}

func (s *SessionCycleRepo) MarkSessionCycleCompleted(ctx context.Context, id int64) error {
	return s.queries.MarkSessionCycleCompleted(ctx, id)
}

func (s *SessionCycleRepo) DeleteSessionCycle(ctx context.Context, id int64) error {
	return s.queries.DeleteSessionCycle(ctx, id)
}

func (s *SessionCycleRepo) mapSQLCSessionCycleToModel(cycle sqlc.SessionCycle) models.SessionCycle {
	result := models.SessionCycle{
		ID:        cycle.ID,
		SessionID: cycle.SessionID,
	}

	if cycle.Type.Valid {
		result.Type = &cycle.Type.String
	}
	if cycle.StartTime.Valid {
		result.StartTime = &cycle.StartTime.Time
	}
	if cycle.EndTime.Valid {
		result.EndTime = &cycle.EndTime.Time
	}
	if cycle.Duration.Valid {
		result.Duration = &cycle.Duration.Int64
	}
	if cycle.Status.Valid {
		result.Status = &cycle.Status.String
	}

	return result
}

