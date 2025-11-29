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

type timeProfileRepo struct {
	queries *sqlc.Queries
}

func NewTimeProfileRepository(db *sql.DB) (*timeProfileRepo, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection cannot be nil")
	}

	return &timeProfileRepo{
		queries: sqlc.New(db),
	}, nil
}

func (r *timeProfileRepo) CreateTimeProfile(ctx context.Context, profile models.TimeProfile) (*models.TimeProfile, error) {
	params := sqlc.CreateTimeProfileParams{
		Name:              profile.Name,
		WorkDuration:      profile.WorkDuration,
		BreakDuration:     profile.BreakDuration,
		LongBreakDuration: profile.LongBreakDuration,
		LongBreakCycle:    sql.NullInt64{Int64: profile.LongBreakCycle, Valid: true},
		IsDefault:         sql.NullBool{Bool: profile.IsDefault, Valid: true},
	}

	created, err := r.queries.CreateTimeProfile(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create time profile: %w", err)
	}

	return r.mapSQLCTimeProfileToModel(created), nil
}

func (r *timeProfileRepo) GetTimeProfile(ctx context.Context, id int64) (*models.TimeProfile, error) {
	profile, err := r.queries.GetTimeProfile(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get time profile: %w", err)
	}

	return r.mapSQLCTimeProfileToModel(profile), nil
}

func (r *timeProfileRepo) GetDefaultTimeProfile(ctx context.Context) (*models.TimeProfile, error) {
	profile, err := r.queries.GetDefaultTimeProfile(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get default time profile: %w", err)
	}

	return r.mapSQLCTimeProfileToModel(profile), nil
}

func (r *timeProfileRepo) ListTimeProfiles(ctx context.Context) ([]*models.TimeProfile, error) {
	profiles, err := r.queries.ListTimeProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list time profiles: %w", err)
	}

	result := make([]*models.TimeProfile, len(profiles))
	for i, p := range profiles {
		result[i] = r.mapSQLCTimeProfileToModel(p)
	}
	return result, nil
}

func (r *timeProfileRepo) UpdateTimeProfile(ctx context.Context, profile models.TimeProfile) (*models.TimeProfile, error) {
	params := sqlc.UpdateTimeProfileParams{
		Name:              profile.Name,
		WorkDuration:      profile.WorkDuration,
		BreakDuration:     profile.BreakDuration,
		LongBreakDuration: profile.LongBreakDuration,
		LongBreakCycle:    sql.NullInt64{Int64: profile.LongBreakCycle, Valid: true},
		IsDefault:         sql.NullBool{Bool: profile.IsDefault, Valid: true},
		ID:                profile.ID,
	}

	updated, err := r.queries.UpdateTimeProfile(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update time profile: %w", err)
	}

	return r.mapSQLCTimeProfileToModel(updated), nil
}

func (r *timeProfileRepo) DeleteTimeProfile(ctx context.Context, id int64) error {
	if err := r.queries.DeleteTimeProfile(ctx, id); err != nil {
		return fmt.Errorf("failed to delete time profile: %w", err)
	}
	return nil
}

func (r *timeProfileRepo) mapSQLCTimeProfileToModel(p sqlc.TimeProfile) *models.TimeProfile {
	return &models.TimeProfile{
		ID:                p.ID,
		Name:              p.Name,
		WorkDuration:      p.WorkDuration,
		BreakDuration:     p.BreakDuration,
		LongBreakDuration: p.LongBreakDuration,
		LongBreakCycle:    p.LongBreakCycle.Int64,
		IsDefault:         p.IsDefault.Bool,
	}
}
