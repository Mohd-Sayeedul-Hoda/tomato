package repository

import (
	"context"
	"time"

	"errors"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
)

var (
	ErrUniqueViolation = errors.New("unique key violation in db")
	ErrNotFound        = errors.New("record not found")
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session models.Session) (int64, error)
	GetSessionByID(ctx context.Context, id int64) (*models.Session, error)
	GetAllSessions(ctx context.Context) ([]*models.Session, error)
	GetActiveSessions(ctx context.Context) ([]*models.Session, error)
	GetCompletedSessions(ctx context.Context) ([]*models.Session, error)
	GetSessionsByTrackedStatus(ctx context.Context, isTracked bool) ([]*models.Session, error)
	GetSessionsForDate(ctx context.Context, date time.Time) ([]*models.Session, error)
	UpdateSession(ctx context.Context, session models.Session) error
	UpdateSessionStatus(ctx context.Context, id int64, status string) error

	UpdateSessionNote(ctx context.Context, id int64, note string) error
	DeleteSession(ctx context.Context, id int64) error
	MarkSessionCompleted(ctx context.Context, id int64) error
}

type SessionCycleRepository interface {
	CreateSessionCycle(ctx context.Context, cycle models.SessionCycle) (int64, error)
	GetSessionCycleByID(ctx context.Context, id int64) (*models.SessionCycle, error)
	GetSessionCycleByStatusWithMetadata(ctx context.Context, status string) ([]*models.SessionCycleWithMetadata, error)
	GetSessionCyclesBySessionID(ctx context.Context, sessionID int64) ([]*models.SessionCycle, error)
	GetSessionCyclesByStatus(ctx context.Context, status string) ([]*models.SessionCycle, error)
	GetSessionCyclesByType(ctx context.Context, cycleType string) ([]*models.SessionCycle, error)
	GetLatestSessionCycleByStatus(ctx context.Context, status string) (*models.SessionCycle, error)
	UpdateSessionCycleStatus(ctx context.Context, id int64, status string) error
	MarkSessionCycleComplete(ctx context.Context, id int64, status string, endTime time.Time, duration int64) error
	MarkSessionCycleCompleted(ctx context.Context, id int64) error
	DeleteSessionCycle(ctx context.Context, id int64) error
}

type TimeProfileRepository interface {
	CreateTimeProfile(ctx context.Context, profile models.TimeProfile) (*models.TimeProfile, error)
	GetTimeProfile(ctx context.Context, id int64) (*models.TimeProfile, error)
	GetDefaultTimeProfile(ctx context.Context) (*models.TimeProfile, error)
	ListTimeProfiles(ctx context.Context) ([]*models.TimeProfile, error)
	UpdateTimeProfile(ctx context.Context, profile models.TimeProfile) (*models.TimeProfile, error)
	DeleteTimeProfile(ctx context.Context, id int64) error
}
