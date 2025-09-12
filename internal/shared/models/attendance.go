package models

import "time"

type Attendance struct {
	ID            int64      `json:"id,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Site          *int64     `json:"site,omitempty"`
	Email         *string    `json:"email,omitempty"`
	Date          *string    `json:"date,omitempty"`
	ClockInTs     *time.Time `json:"clock_in_ts,omitempty"`
	ClockInLat    *float64   `json:"clock_in_lat,omitempty"`
	ClockInLng    *float64   `json:"clock_in_lng,omitempty"`
	ClockInPhoto  *string    `json:"clock_in_photo,omitempty"`
	ClockOutTs    *time.Time `json:"clock_out_ts,omitempty"`
	ClockOutLat   *float64   `json:"clock_out_lat,omitempty"`
	ClockOutLng   *float64   `json:"clock_out_lng,omitempty"`
	ClockOutPhoto *string    `json:"clock_out_photo,omitempty"`
	Status        *string    `json:"status,omitempty"`
}
