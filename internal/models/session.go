package models

import (
	"time"
)

type Session struct {
	ID              int64      `json:"id"`
	Label           string     `json:"label"`
	Status          string     `json:"status"`
	SessionEstimate *int64     `json:"session_estimate,omitempty"`
	IsTracked       *bool      `json:"is_tracked,omitempty"`
	Note            *string    `json:"note,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

type SessionFilter struct {
	Status    *string
	Date      *time.Time
	IsTracked *bool
}

type SessionCycleFilter struct {
	SessionID *int64
	Status    *string
	Type      *string
	Limit     *int
}

type TimeProfile struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	WorkDuration      int64  `json:"work_duration"`
	BreakDuration     int64  `json:"break_duration"`
	LongBreakDuration int64  `json:"long_break_duration"`
	LongBreakCycle    int64  `json:"long_break_cycle"`
	IsDefault         bool   `json:"is_default"`
}

type SessionCycle struct {
	ID             int64      `json:"id"`
	SessionID      int64      `json:"session_id"`
	TimerProfileID int64      `json:"timer_profile_id"`
	Type           *string    `json:"type,omitempty"`
	StartTime      *time.Time `json:"start_time,omitempty"`
	EndTime        *time.Time `json:"end_time,omitempty"`
	Duration       *int64     `json:"duration,omitempty"`
	Status         *string    `json:"status,omitempty"`
}

type SessionCycleWithMetadata struct {
	ID                int64      `json:"id"`
	SessionID         int64      `json:"session_id"`
	Type              *string    `json:"type,omitempty"`
	StartTime         *time.Time `json:"start_time,omitempty"`
	EndTime           *time.Time `json:"end_time,omitempty"`
	Duration          *int64     `json:"duration,omitempty"`
	Status            *string    `json:"status,omitempty"`
	WorkDuration      int64      `json:"work_duration"`
	BreakDuration     int64      `json:"break_duration"`
	LongBreakDuration int64      `json:"long_break_duration"`
	LongBreakCycle    *int64     `json:"long_break_cycle,omitempty"`
}
