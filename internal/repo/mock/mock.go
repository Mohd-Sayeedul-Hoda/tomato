package mock

import (
	"context"
	"time"

	"github.com/Mohd-Sayeedul-Hoda/tomato/internal/models"
)

type MockSessionRepo struct {
	CreateSessionFunc func(ctx context.Context, session models.Session) (int64, error)
}

func (m *MockSessionRepo) CreateSession(ctx context.Context, session models.Session) (int64, error) {
	if m.CreateSessionFunc != nil {
		return m.CreateSessionFunc(ctx, session)
	}
	return 1, nil
}

func (m *MockSessionRepo) GetSessionByID(ctx context.Context, id int64) (*models.Session, error) {
	return nil, nil
}
func (m *MockSessionRepo) ListSessions(ctx context.Context, filter models.SessionFilter) ([]*models.Session, error) {
	return nil, nil
}
func (m *MockSessionRepo) UpdateSession(ctx context.Context, session models.Session) error {
	return nil
}
func (m *MockSessionRepo) UpdateSessionStatus(ctx context.Context, id int64, status string) error {
	return nil
}
func (m *MockSessionRepo) UpdateSessionNote(ctx context.Context, id int64, note string) error {
	return nil
}
func (m *MockSessionRepo) DeleteSession(ctx context.Context, id int64) error {
	return nil
}
func (m *MockSessionRepo) MarkSessionCompleted(ctx context.Context, id int64) error {
	return nil
}

type MockSessionCycleRepo struct{}

func (m *MockSessionCycleRepo) CreateSessionCycle(ctx context.Context, cycle models.SessionCycle) (int64, error) {
	return 0, nil
}
func (m *MockSessionCycleRepo) GetSessionCycleByID(ctx context.Context, id int64) (*models.SessionCycle, error) {
	return nil, nil
}
func (m *MockSessionCycleRepo) GetSessionCycleByStatusWithMetadata(ctx context.Context, status string) ([]*models.SessionCycleWithMetadata, error) {
	return nil, nil
}
func (m *MockSessionCycleRepo) ListSessionCycles(ctx context.Context, filter models.SessionCycleFilter) ([]*models.SessionCycle, error) {
	return nil, nil
}
func (m *MockSessionCycleRepo) UpdateSessionCycleStatus(ctx context.Context, id int64, status string) error {
	return nil
}
func (m *MockSessionCycleRepo) MarkSessionCycleComplete(ctx context.Context, id int64, status string, endTime time.Time, duration int64) error {
	return nil
}
func (m *MockSessionCycleRepo) MarkSessionCycleCompleted(ctx context.Context, id int64) error {
	return nil
}
func (m *MockSessionCycleRepo) DeleteSessionCycle(ctx context.Context, id int64) error {
	return nil
}
