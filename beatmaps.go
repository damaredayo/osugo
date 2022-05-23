package osugo

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Beatmap struct {
	Accuracy         float64    `json:"accuracy"`
	ApproachRate     float64    `json:"ar"`
	BeatmapsetId     int        `json:"beatmapset_id"`
	Bpm              float64    `json:"bpm"`
	Convert          bool       `json:"convert"`
	CountCircles     int        `json:"count_circles"`
	CountSliders     int        `json:"count_sliders"`
	CountSpinners    int        `json:"count_spinners"`
	CircleSize       float64    `json:"cs"`
	DeletedAt        *time.Time `json:"deleted_at"`
	DifficultyRating float64    `json:"difficulty_rating"`
	Drain            float64    `json:"drain"`
	HitLength        float64    `json:"hit_length"`
	Id               int        `json:"id"`
	IsScorable       bool       `json:"is_scorable"`
	LastUpdated      *time.Time `json:"last_updated"`
	ModeInt          int        `json:"mode_int"`
	Passcount        int        `json:"passcount"`
	Playcount        int        `json:"playcount"`
	Ranked           int        `json:"ranked"`
	Status           string     `json:"status"`
	TotalLength      int        `json:"total_length"`
	Url              string     `json:"url"`
	UserId           int        `json:"user_id"`
	Version          string     `json:"version"`
}

type BeatmapCompact struct {
	BeatmapsetId     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	Id               int     `json:"id"`
	Mode             string  `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserId           int     `json:"user_id"`
	Version          string  `json:"version"`
}

type BeatmapDifficultyAttributes struct {
	// Shared
	MaxCombo       int     `json:"max_combo"`
	StarRating     float64 `json:"star_rating"`
	ApproachRate   float64 `json:"approach_rate"`    // Excluding osu!mania
	GreatHitWindow float64 `json:"great_hit_window"` // osu!taiko and osu!mania

	// osu! standard
	AimDifficulty        float64 `json:"aim_difficulty"`
	FlashlightDifficulty float64 `json:"flashlight_difficulty"`
	OverallDifficulty    float64 `json:"overall_difficulty"`
	SliderFactor         float64 `json:"slider_factor"`
	SpeedDifficulty      float64 `json:"speed_difficulty"`

	// osu!taiko
	StaminaDifficulty float64 `json:"stamina_difficulty"`
	RhythmDifficulty  float64 `json:"rhythm_difficulty"`
	ColourDifficulty  float64 `json:"colour_difficulty"`

	// osu!catch only contains ApproachRate

	// osu!mania
	ScoreMultiplier float64 `json:"score_multiplier"`
}

type BeatmapPlaycount struct {
	BeatmapId  int                `json:"beatmap_id"`
	Beatmap    *BeatmapCompact    `json:"beatmap"`
	Beatmapset *BeatmapsetCompact `json:"beatmapset"`
	Count      int                `json:"count"`
}

type BeatmapScores struct {
	Scores    []*Score          `json:"scores"`
	UserScore *BeatmapUserScore `json:"userScore"` // because only this field uses camel case? pepega
}

type BeatmapUserScore struct {
	Position int    `json:"position"`
	Score    *Score `json:"score"`
}

type Beatmapset struct {
	Artist                       string       `json:"artist"`
	ArtistUnicode                string       `json:"artist_unicode"`
	AvailabilityDownloadDisabled bool         `json:"availability.download_disabled"`
	AvailabilityMoreInformation  bool         `json:"availability.more_information"`
	Beatmaps                     []*Beatmap   `json:"beatmaps,omitempty"`
	Bpm                          float64      `json:"bpm"`
	CanBeHyped                   bool         `json:"can_be_hyped"`
	Covers                       *Covers      `json:"covers"`
	Creator                      string       `json:"creator"`
	Description                  string       `json:"description,omitempty"`
	DiscussionEnabled            bool         `json:"discussion_enabled"`
	DiscussionLocked             bool         `json:"discussion_locked"`
	FavouriteCount               int          `json:"favourite_count"`
	HasFavourited                bool         `json:"has_favourited"`
	HypeCurrent                  int          `json:"hype.current"`
	HypeRequired                 int          `json:"hype.required"`
	Id                           int          `json:"id"`
	IsScorable                   bool         `json:"is_scorable"`
	LastUpdated                  *time.Time   `json:"last_updated"`
	LegacyThreadURL              string       `json:"legacy_thread_url"`
	NominationsCurrent           int          `json:"nominations.current"`
	NominationsRequired          int          `json:"nominations.required"`
	Nsfw                         bool         `json:"nsfw"`
	Playcount                    int          `json:"playcount"`
	PreviewURL                   string       `json:"preview_url"`
	Ranked                       RankedStatus `json:"ranked"`
	RankedDate                   *time.Time   `json:"ranked_date"`
	Source                       string       `json:"source"`
	Status                       string       `json:"status"`
	Storyboard                   string       `json:"storyboard"`
	SubmittedDate                *time.Time   `json:"submitted_date"`
	Tags                         string       `json:"tags"`
	Title                        string       `json:"title"`
	TitleUnicode                 string       `json:"title_unicode"`
	UserId                       int          `json:"user_id"`
	Video                        bool         `json:"video"`
}

type BeatmapsetCompact struct {
	Artist         string     `json:"artist"`
	ArtistUnicode  string     `json:"artist_unicode"`
	Beatmaps       []*Beatmap `json:"beatmaps,omitempty"`
	Covers         *Covers    `json:"covers"`
	Creator        string     `json:"creator"`
	Description    string     `json:"description,omitempty"`
	FavouriteCount int        `json:"favourite_count"`
	Id             int        `json:"id"`
	Nsfw           bool       `json:"nsfw"`
	Playcount      int        `json:"playcount"`
	PreviewURL     string     `json:"preview_url"`
	Source         string     `json:"source"`
	Status         string     `json:"status"`
	Title          string     `json:"title"`
	TitleUnicode   string     `json:"title_unicode"`
	UserId         int        `json:"user_id"`
	Video          bool       `json:"video"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2x     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2x      string `json:"card@2x"`
	List        string `json:"list"`
	List2x      string `json:"list@2x"`
	SlimCover   string `json:"slimcover"`
	SlimCover2x string `json:"slimcover@2x"`
}

func (c *Client) LookupBeatmap(id, filename, checksum string) (beatmap *Beatmap, err error) {
	params := map[string]string{
		"id":       id,
		"filename": filename,
		"checksum": checksum,
	}
	b, err := c.request(http.MethodGet, "beatmaps/lookup", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmap)
	return
}

func (c *Client) GetUserBeatmapScore(beatmapId, user, mods string, mode Gamemode) (beatmap *Beatmap, err error) {
	params := map[string]string{
		"mode": mode.String(),
		"mods": mods,
	}
	b, err := c.request(http.MethodGet, "beatmaps/"+beatmapId+"/scores/users/"+user, params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmap)
	return
}

func (c *Client) GetUserBeatmapScores(beatmapId, user string, mode Gamemode) (beatmaps []*Beatmap, err error) {
	params := map[string]string{
		"mode": mode.String(),
	}
	b, err := c.request(http.MethodGet, "beatmaps/"+beatmapId+"/scores/users/"+user+"/all", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmaps)
	return
}

func (c *Client) GetBeatmapScores(beatmapId, mods, scoreType string, mode Gamemode) (beatmaps *BeatmapScores, err error) {
	params := map[string]string{
		"mode": mode.String(),
		"mods": mods,
		"type": scoreType,
	}
	b, err := c.request(http.MethodGet, "beatmaps/"+beatmapId+"/scores", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmaps)
	return
}

func (c *Client) GetBeatmaps(beatmapIds []string) (beatmaps []*BeatmapCompact, err error) {
	params := map[string]string{
		"ids[]": strings.Join(beatmapIds, ","),
	}
	b, err := c.request(http.MethodGet, "beatmaps", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmaps)
	return
}

func (c *Client) GetBeatmap(beatmapId string) (beatmap *Beatmap, err error) {
	b, err := c.request(http.MethodGet, "beatmaps/"+beatmapId, nil, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmap)
	return
}

func (c *Client) GetBeatmapAttributes(beatmapId string, mods []string, ruleset Gamemode, ruleset_id int) (beatmap *Beatmap, err error) {

	s := struct {
		Mods      []string `json:"mods,omitempty"`
		Ruleset   string   `json:"ruleset,omitempty"`
		RulesetID int      `json:"ruleset_id,omitempty"`
	}{
		Mods:      mods,
		Ruleset:   ruleset.String(),
		RulesetID: ruleset_id,
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		return
	}

	b, err := c.request(http.MethodGet, "beatmaps/"+beatmapId+"/attributes", nil, bytes)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmap)
	return
}

func (c *Client) BeatmapSearch(filters string) (beatmaps []*BeatmapCompact, err error) {
	b, err := c.request(http.MethodGet, "beatmaps/search/"+filters, nil, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmaps)
	return
}
