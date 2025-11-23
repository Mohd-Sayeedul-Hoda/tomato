package models

import (
	"time"
)

type Session struct {
	ID                int64      `json:"id"`
	Label             string     `json:"label"`
	WorkDuration      int64      `json:"work_duration"`
	BreakDuration     int64      `json:"break_duration"`
	LongBreakDuration int64      `json:"long_break_duration"`
	LongBreakCycle    *int64     `json:"long_break_cycle,omitempty"`
	StartTime         time.Time  `json:"start_time"`
	EndTime           *time.Time `json:"end_time,omitempty"`
	Status            string     `json:"status"`
	SessionEstimate   *int64     `json:"session_estimate,omitempty"`
	IsTracked         *bool      `json:"is_tracked,omitempty"`
	Note              *string    `json:"note,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
}

type SessionCycle struct {
	ID        int64      `json:"id"`
	SessionID int64      `json:"session_id"`
	Type      *string    `json:"type,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	Duration  *int64     `json:"duration,omitempty"`
	Status    *string    `json:"status,omitempty"`
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