package osugo

import "time"

type Score struct {
	Accuracy              float64               `json:"accuracy"`
	BestID                int64                 `json:"best_id"`
	CreatedAt             time.Time             `json:"created_at"`
	ID                    int64                 `json:"id"`
	MaxCombo              int                   `json:"max_combo"`
	Mode                  string                `json:"mode"`
	ModeInt               int                   `json:"mode_int"`
	Mods                  []string              `json:"mods"`
	Passed                bool                  `json:"passed"`
	Perfect               bool                  `json:"perfect"`
	Pp                    float64               `json:"pp"`
	Rank                  string                `json:"rank"`
	Replay                bool                  `json:"replay"`
	Score                 int                   `json:"score"`
	Statistics            PlayStatistics        `json:"statistics"`
	UserID                int                   `json:"user_id"`
	CurrentUserAttributes CurrentUserAttributes `json:"current_user_attributes"`
	Beatmap               Beatmap               `json:"beatmap"`
	Beatmapset            Beatmapset            `json:"beatmapset"`
	User                  User                  `json:"user"`
	Weight                Weight                `json:"weight"`
}
type PlayStatistics struct {
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	Count50   int `json:"count_50"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
}
type CurrentUserAttributes struct {
	Pin interface{} `json:"pin"`
}
type Weight struct {
	Percentage int     `json:"percentage"`
	Pp         float64 `json:"pp"`
}
