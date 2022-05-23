package osugo

import "time"

type Score struct {
	Id        string    `json:"id"`
	BestId    string    `json:"best_id"`
	UserId    string    `json:"user_id"`
	Accuracy  float64   `json:"accuracy"`
	Mods      string    `json:"mods"`
	Score     int       `json:"score"`
	MaxCombo  int       `json:"max_combo"`
	Perfect   bool      `json:"perfect"`
	Count50   int       `json:"statistics.count_50"`
	Count100  int       `json:"statistics.count_100"`
	Count300  int       `json:"statistics.count_300"`
	CountGeki int       `json:"statistics.count_geki"`
	CountKatu int       `json:"statistics.count_katu"`
	CountMiss int       `json:"statistics.count_miss"`
	Passed    bool      `json:"passed"`
	PP        float64   `json:"pp"`
	Rank      int       `json:"rank"`
	CreatedAt time.Time `json:"created_at"`
	Mode      string    `json:"mode"`
	ModeInt   int       `json:"mode_int"`
	Replay    string    `json:"replay"`

	Beatmap     *Beatmap    `json:"beatmap,omitempty"`
	Beatmapset  *Beatmapset `json:"beatmapset,omitempty"`
	RankCountry int         `json:"rank_country,omitempty"`
	RankGlobal  int         `json:"rank_global,omitempty"`
	Weight      int         `json:"weight,omitempty"`
	User        string      `json:"user,omitempty"`
	Match       string      `json:"match,omitempty"` // ?
}
