package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/sqlc"
)

type sessionRepo struct {
	queries *sqlc.Queries
}

func NewSessionRepository(db *sql.DB) (*sessionRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection cannot be nil")
	}

	return &sessionRepo{
		queries: sqlc.New(db),
	}, nil
}

func (s *sessionRepo) CreateSession(ctx context.Context, session models.Session) (int64, error) {
	params := sqlc.CreateSessionParams{
		Label:             session.Label,
		Note:              sql.NullString{String: "", Valid: false},
		Status:            session.Status,
		SessionEstimate:   sql.NullInt64{Int64: 0, Valid: false},
		IsTracked:         sql.NullBool{Bool: false, Valid: false},
		StartTime:         session.StartTime,
		WorkDuration:      session.WorkDuration,
		BreakDuration:     session.BreakDuration,
		LongBreakDuration: session.LongBreakDuration,
		LongBreakCycle:    sql.NullInt64{Int64: 0, Valid: false},
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
	if session.LongBreakCycle != nil {
		params.LongBreakCycle = sql.NullInt64{Int64: *session.LongBreakCycle, Valid: true}
	}

	id, err := s.queries.CreateSession(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *sessionRepo) GetSessionByID(ctx context.Context, id int64) (*models.Session, error) {
	session, err := s.queries.GetSessionById(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.mapSQLCSessionToModel(session), nil
}

func (s *sessionRepo) GetAllSessions(ctx context.Context) ([]*models.Session, error) {
	sessions, err := s.queries.GetAllSessions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Session, len(sessions))
	for i, session := range sessions {
		result[i] = s.mapSQLCSessionToModel(session)
	}
	return result, nil
}

func (s *sessionRepo) GetActiveSessions(ctx context.Context) ([]*models.Session, error) {
	sessions, err := s.queries.GetActiveSessions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Session, len(sessions))
	for i, session := range sessions {
		result[i] = s.mapSQLCSessionToModel(session)
	}
	return result, nil
}

func (s *sessionRepo) GetCompletedSessions(ctx context.Context) ([]*models.Session, error) {
	sessions, err := s.queries.GetCompletedSessions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Session, len(sessions))
	for i, session := range sessions {
		result[i] = s.mapSQLCSessionToModel(session)
	}
	return result, nil
}

func (s *sessionRepo) GetSessionsByTrackedStatus(ctx context.Context, isTracked bool) ([]*models.Session, error) {
	sqlNullBool := sql.NullBool{Bool: isTracked, Valid: true}
	sessions, err := s.queries.GetSessionsByTrackedStatus(ctx, sqlNullBool)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Session, len(sessions))
	for i, session := range sessions {
		result[i] = s.mapSQLCSessionToModel(session)
	}
	return result, nil
}

func (s *sessionRepo) GetSessionsForDate(ctx context.Context, date time.Time) ([]*models.Session, error) {
	sqlNullTime := sql.NullTime{Time: date, Valid: true}
	sessions, err := s.queries.GetSessionsForDate(ctx, sqlNullTime)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Session, len(sessions))
	for i, session := range sessions {
		result[i] = s.mapSQLCSessionToModel(session)
	}
	return result, nil
}

func (s *sessionRepo) UpdateSession(ctx context.Context, session models.Session) error {
	params := sqlc.UpdateSessionParams{
		ID:                session.ID,
		Label:             session.Label,
		Note:              sql.NullString{String: "", Valid: false},
		Status:            session.Status,
		SessionEstimate:   sql.NullInt64{Int64: 0, Valid: false},
		IsTracked:         sql.NullBool{Bool: false, Valid: false},
		StartTime:         session.StartTime,
		EndTime:           sql.NullTime{Time: time.Time{}, Valid: false},
		WorkDuration:      session.WorkDuration,
		BreakDuration:     session.BreakDuration,
		LongBreakDuration: session.LongBreakDuration,
		LongBreakCycle:    sql.NullInt64{Int64: 0, Valid: false},
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
	if session.EndTime != nil {
		params.EndTime = sql.NullTime{Time: *session.EndTime, Valid: true}
	}
	if session.LongBreakCycle != nil {
		params.LongBreakCycle = sql.NullInt64{Int64: *session.LongBreakCycle, Valid: true}
	}

	return s.queries.UpdateSession(ctx, params)
}

func (s *sessionRepo) UpdateSessionStatus(ctx context.Context, id int64, status string) error {
	params := sqlc.UpdateSessionStatusParams{
		ID:     id,
		Status: status,
	}
	return s.queries.UpdateSessionStatus(ctx, params)
}

func (s *sessionRepo) UpdateSessionEndTime(ctx context.Context, id int64, endTime time.Time) error {
	params := sqlc.UpdateSessionEndTimeParams{
		ID:      id,
		EndTime: sql.NullTime{Time: endTime, Valid: true},
	}
	return s.queries.UpdateSessionEndTime(ctx, params)
}

func (s *sessionRepo) UpdateSessionNote(ctx context.Context, id int64, note string) error {
	params := sqlc.UpdateSessionNoteParams{
		ID:   id,
		Note: sql.NullString{String: note, Valid: true},
	}
	return s.queries.UpdateSessionNote(ctx, params)
}

func (s *sessionRepo) DeleteSession(ctx context.Context, id int64) error {
	return s.queries.DeleteSession(ctx, id)
}

func (s *sessionRepo) MarkSessionCompleted(ctx context.Context, id int64) error {
	return s.queries.MarkSessionCompleted(ctx, id)
}

func (s *sessionRepo) mapSQLCSessionToModel(session sqlc.Session) *models.Session {
	result := &models.Session{
		ID:                session.ID,
		Label:             session.Label,
		WorkDuration:      session.WorkDuration,
		BreakDuration:     session.BreakDuration,
		LongBreakDuration: session.LongBreakDuration,
		StartTime:         session.StartTime,
		Status:            session.Status,
	}

	if session.LongBreakCycle.Valid {
		result.LongBreakCycle = &session.LongBreakCycle.Int64
	}
	if session.EndTime.Valid {
		result.EndTime = &session.EndTime.Time
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

	return result
}
