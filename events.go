package osugo

import "time"

type EventBeatmap struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type EventBeatmapset struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type EventUser struct {
	Username         string `json:"username"`
	Url              string `json:"url"`
	PreviousUsername string `json:"previous_username,omitempty"`
}

type Event struct {
	CreatedAt  time.Time        `json:"created_at"`
	Id         int              `json:"id"`
	Type       string           `json:"type"`
	Beatmap    *EventBeatmap    `json:"beatmap,omitempty"`
	Beatmapset *EventBeatmapset `json:"beatmapset,omitempty"`
	User       *EventUser       `json:"user,omitempty"`
	// TODO: finish events
}
