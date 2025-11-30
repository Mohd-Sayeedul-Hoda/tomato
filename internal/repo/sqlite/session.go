package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
	repo "github.com/Mohd-Sayeedul-Hoda/tomato/internal/repo"
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/sqlc"
)

type sessionRepo struct {
	queries *sqlc.Queries
}

func NewSessionRepository(db *sql.DB) (*sessionRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection cannot be nil")
	}

	// TODO: have to use interface define in sqlc queries
	return &sessionRepo{
		queries: sqlc.New(db),
	}, nil
}

func (s *sessionRepo) CreateSession(ctx context.Context, session models.Session) (int64, error) {
	params := sqlc.CreateSessionParams{
		Label:           session.Label,
		Note:            sql.NullString{String: "", Valid: false},
		Status:          session.Status,
		SessionEstimate: sql.NullInt64{Int64: 0, Valid: false},
		IsTracked:       sql.NullBool{Bool: false, Valid: false},
	}

	if session.Note != nil {
		params.Note = sql.NullString{String: *session.Note, Valid: true}
	}
	if session.SessionEstimate != nil {
		params.SessionEstimate = sql.NullInt64{Int64: *session.SessionEstimate, Valid: true}
	}
	if session.IsTracked != nil {
		params.IsTracked = sql.NullBool{Bool: *session.IsTracked, Valid: true}
	}

	id, err := s.queries.CreateSession(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("failed to create session: %w", err)
	}
	return id, nil
}

func (s *sessionRepo) GetSessionByID(ctx context.Context, id int64) (*models.Session, error) {
	session, err := s.queries.GetSessionById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get session by id: %w", err)
	}
	return s.mapSQLCSessionToModel(session), nil
}

func (s *sessionRepo) ListSessions(ctx context.Context, filter models.SessionFilter) ([]*models.Session, error) {
	params := sqlc.ListSessionsParams{}

	if filter.Status != nil {
		params.Status = *filter.Status
	}
	if filter.Date != nil {
		params.Date = filter.Date.Format("2006-01-02")
	}
	if filter.IsTracked != nil {
		params.IsTracked = *filter.IsTracked
	}

	sessions, err := s.queries.ListSessions(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list sessions: %w", err)
	}

	result := make([]*models.Session, len(sessions))
	for i, session := range sessions {
		result[i] = s.mapSQLCSessionToModel(session)
	}
	return result, nil
}

func (s *sessionRepo) UpdateSession(ctx context.Context, session models.Session) error {
	params := sqlc.UpdateSessionParams{
		ID:              session.ID,
		Label:           session.Label,
		Note:            sql.NullString{String: "", Valid: false},
		Status:          session.Status,
		SessionEstimate: sql.NullInt64{Int64: 0, Valid: false},
		IsTracked:       sql.NullBool{Bool: false, Valid: false},
	}

	if session.Note != nil {
		params.Note = sql.NullString{String: *session.Note, Valid: true}
	}
	if session.SessionEstimate != nil {
		params.SessionEstimate = sql.NullInt64{Int64: *session.SessionEstimate, Valid: true}
	}
	if session.IsTracked != nil {
		params.IsTracked = sql.NullBool{Bool: *session.IsTracked, Valid: true}
	}

	if err := s.queries.UpdateSession(ctx, params); err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}
	return nil
}

func (s *sessionRepo) UpdateSessionStatus(ctx context.Context, id int64, status string) error {
	params := sqlc.UpdateSessionStatusParams{
		ID:     id,
		Status: status,
	}
	if err := s.queries.UpdateSessionStatus(ctx, params); err != nil {
		return fmt.Errorf("failed to update session status: %w", err)
	}
	return nil
}

func (s *sessionRepo) UpdateSessionNote(ctx context.Context, id int64, note string) error {
	params := sqlc.UpdateSessionNoteParams{
		ID:   id,
		Note: sql.NullString{String: note, Valid: true},
	}
	if err := s.queries.UpdateSessionNote(ctx, params); err != nil {
		return fmt.Errorf("failed to update session note: %w", err)
	}
	return nil
}

func (s *sessionRepo) DeleteSession(ctx context.Context, id int64) error {
	if err := s.queries.DeleteSession(ctx, id); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

func (s *sessionRepo) MarkSessionCompleted(ctx context.Context, id int64) error {
	if err := s.queries.MarkSessionCompleted(ctx, id); err != nil {
		return fmt.Errorf("failed to mark session completed: %w", err)
	}
	return nil
}

func (s *sessionRepo) mapSQLCSessionToModel(session sqlc.Session) *models.Session {
	result := &models.Session{
		ID:     session.ID,
		Label:  session.Label,
		Status: session.Status,
	}

	if session.SessionEstimate.Valid {
		result.SessionEstimate = &session.SessionEstimate.Int64
	}
	if session.IsTracked.Valid {
		result.IsTracked = &session.IsTracked.Bool
	}
	if session.Note.Valid {
		result.Note = &session.Note.String
	}
	if session.CreatedAt.Valid {
		result.CreatedAt = &session.CreatedAt.Time
	}
	if session.UpdatedAt.Valid {
		result.UpdatedAt = &session.UpdatedAt.Time
	}

	return result
}
