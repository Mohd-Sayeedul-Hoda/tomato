package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
	repo "github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo"
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/sqlc"
)

type sessionCycleRepo struct {
	queries *sqlc.Queries
}

func NewSessionCycleRepository(db *sql.DB) (*sessionCycleRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection cannot be nil")
	}

	return &sessionCycleRepo{
		queries: sqlc.New(db),
	}, nil
}

func (s *sessionCycleRepo) CreateSessionCycle(ctx context.Context, cycle models.SessionCycle) (int64, error) {
	params := sqlc.CreateSessionCycleParams{
		SessionID:      cycle.SessionID,
		TimerProfileID: cycle.TimerProfileID,
		Type:           sql.NullString{String: "", Valid: false},
		StartTime:      sql.NullTime{Time: time.Time{}, Valid: false},
		Status:         sql.NullString{String: "", Valid: false},
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
		return 0, fmt.Errorf("failed to create session cycle: %w", err)
	}
	return id, nil
}

func (s *sessionCycleRepo) GetSessionCycleByID(ctx context.Context, id int64) (*models.SessionCycle, error) {
	cycle, err := s.queries.GetSessionCycleByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get session cycle by id: %w", err)
	}
	return s.mapSQLCSessionCycleToModel(cycle), nil
}

func (s *sessionCycleRepo) ListSessionCycles(ctx context.Context, filter models.SessionCycleFilter) ([]*models.SessionCycle, error) {
	params := sqlc.ListSessionCyclesParams{
		Limit: 1000,
	}

	if filter.SessionID != nil {
		params.SessionID = *filter.SessionID
	}
	if filter.Status != nil {
		params.Status = *filter.Status
	}
	if filter.Type != nil {
		params.Type = *filter.Type
	}
	if filter.Limit != nil {
		params.Limit = int64(*filter.Limit)
	}

	cycles, err := s.queries.ListSessionCycles(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list session cycles: %w", err)
	}

	result := make([]*models.SessionCycle, len(cycles))
	for i, cycle := range cycles {
		result[i] = s.mapSQLCSessionCycleToModel(cycle)
	}
	return result, nil
}

func (s *sessionCycleRepo) UpdateSessionCycleStatus(ctx context.Context, id int64, status string) error {
	params := sqlc.UpdateSessionCycleStatusParams{
		ID:     id,
		Status: sql.NullString{String: status, Valid: true},
	}
	if err := s.queries.UpdateSessionCycleStatus(ctx, params); err != nil {
		return fmt.Errorf("failed to update session cycle status: %w", err)
	}
	return nil
}

func (s *sessionCycleRepo) MarkSessionCycleComplete(ctx context.Context, id int64, status string, endTime time.Time, duration int64) error {
	params := sqlc.MarkSessionCycleCompleteParams{
		ID:       id,
		Status:   sql.NullString{String: status, Valid: true},
		EndTime:  sql.NullTime{Time: endTime, Valid: true},
		Duration: sql.NullInt64{Int64: duration, Valid: true},
	}
	if err := s.queries.MarkSessionCycleComplete(ctx, params); err != nil {
		return fmt.Errorf("failed to mark session cycle complete: %w", err)
	}
	return nil
}

func (s *sessionCycleRepo) MarkSessionCycleCompleted(ctx context.Context, id int64) error {
	if err := s.queries.MarkSessionCycleCompleted(ctx, id); err != nil {
		return fmt.Errorf("failed to mark session cycle completed: %w", err)
	}
	return nil
}

func (s *sessionCycleRepo) DeleteSessionCycle(ctx context.Context, id int64) error {
	if err := s.queries.DeleteSessionCycle(ctx, id); err != nil {
		return fmt.Errorf("failed to delete session cycle: %w", err)
	}
	return nil
}

func (s *sessionCycleRepo) mapSQLCSessionCycleToModel(cycle sqlc.SessionCycle) *models.SessionCycle {
	result := &models.SessionCycle{
		ID:             cycle.ID,
		SessionID:      cycle.SessionID,
		TimerProfileID: cycle.TimerProfileID,
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

func (s *sessionCycleRepo) mapSQLCSessionCycleWithMetadataToModel(row sqlc.GetSessionCycleByStatusWithMetadataRow) *models.SessionCycleWithMetadata {
	result := &models.SessionCycleWithMetadata{
		ID:        row.ID,
		SessionID: row.SessionID,
	}

	if row.WorkDuration.Valid {
		result.WorkDuration = row.WorkDuration.Int64
	}
	if row.BreakDuration.Valid {
		result.BreakDuration = row.BreakDuration.Int64
	}
	if row.LongBreakDuration.Valid {
		result.LongBreakDuration = row.LongBreakDuration.Int64
	}

	if row.Type.Valid {
		result.Type = &row.Type.String
	}
	if row.StartTime.Valid {
		result.StartTime = &row.StartTime.Time
	}
	if row.EndTime.Valid {
		result.EndTime = &row.EndTime.Time
	}
	if row.Duration.Valid {
		result.Duration = &row.Duration.Int64
	}
	if row.Status.Valid {
		result.Status = &row.Status.String
	}
	if row.LongBreakCycle.Valid {
		result.LongBreakCycle = &row.LongBreakCycle.Int64
	}

	return result
}

func (s *sessionCycleRepo) GetSessionCycleByStatusWithMetadata(ctx context.Context, status string) ([]*models.SessionCycleWithMetadata, error) {
	sqlNullString := sql.NullString{String: status, Valid: true}
	rows, err := s.queries.GetSessionCycleByStatusWithMetadata(ctx, sqlNullString)
	if err != nil {
		return nil, fmt.Errorf("failed to get session cycle by status with metadata: %w", err)
	}

	result := make([]*models.SessionCycleWithMetadata, len(rows))
	for i, row := range rows {
		result[i] = s.mapSQLCSessionCycleWithMetadataToModel(row)
	}
	return result, nil
}
