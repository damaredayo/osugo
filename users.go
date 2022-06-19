package osugo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	AvatarURL                        string                 `json:"avatar_url"`
	CountryCode                      string                 `json:"country_code"`
	DefaultGroup                     string                 `json:"default_group"`
	ID                               int                    `json:"id"`
	IsActive                         bool                   `json:"is_active"`
	IsBot                            bool                   `json:"is_bot"`
	IsDeleted                        bool                   `json:"is_deleted"`
	IsOnline                         bool                   `json:"is_online"`
	IsSupporter                      bool                   `json:"is_supporter"`
	LastVisit                        time.Time              `json:"last_visit"`
	PmFriendsOnly                    bool                   `json:"pm_friends_only"`
	ProfileColour                    interface{}            `json:"profile_colour"`
	Username                         string                 `json:"username"`
	CoverURL                         string                 `json:"cover_url"`
	Discord                          string                 `json:"discord"`
	HasSupported                     bool                   `json:"has_supported"`
	Interests                        string                 `json:"interests"`
	JoinDate                         time.Time              `json:"join_date"`
	Kudosu                           Kudosu                 `json:"kudosu"`
	Location                         string                 `json:"location"`
	MaxBlocks                        int                    `json:"max_blocks"`
	MaxFriends                       int                    `json:"max_friends"`
	Occupation                       string                 `json:"occupation"`
	Playmode                         string                 `json:"playmode"`
	Playstyle                        []string               `json:"playstyle"`
	PostCount                        int                    `json:"post_count"`
	ProfileOrder                     []string               `json:"profile_order"`
	Title                            string                 `json:"title"`
	TitleURL                         string                 `json:"title_url"`
	Twitter                          string                 `json:"twitter"`
	Website                          string                 `json:"website"`
	Country                          Country                `json:"country"`
	Cover                            Cover                  `json:"cover"`
	AccountHistory                   []interface{}          `json:"account_history"`
	ActiveTournamentBanner           interface{}            `json:"active_tournament_banner"`
	Badges                           []Badges               `json:"badges"`
	BeatmapPlaycountsCount           int                    `json:"beatmap_playcounts_count"`
	CommentsCount                    int                    `json:"comments_count"`
	FavouriteBeatmapsetCount         int                    `json:"favourite_beatmapset_count"`
	FollowerCount                    int                    `json:"follower_count"`
	GraveyardBeatmapsetCount         int                    `json:"graveyard_beatmapset_count"`
	Groups                           []interface{}          `json:"groups"`
	GuestBeatmapsetCount             int                    `json:"guest_beatmapset_count"`
	LovedBeatmapsetCount             int                    `json:"loved_beatmapset_count"`
	MappingFollowerCount             int                    `json:"mapping_follower_count"`
	MonthlyPlaycounts                []MonthlyPlaycounts    `json:"monthly_playcounts"`
	Page                             Page                   `json:"page"`
	PendingBeatmapsetCount           int                    `json:"pending_beatmapset_count"`
	PreviousUsernames                []string               `json:"previous_usernames"`
	RankedBeatmapsetCount            int                    `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts             []ReplaysWatchedCounts `json:"replays_watched_counts"`
	ScoresBestCount                  int                    `json:"scores_best_count"`
	ScoresFirstCount                 int                    `json:"scores_first_count"`
	ScoresPinnedCount                int                    `json:"scores_pinned_count"`
	ScoresRecentCount                int                    `json:"scores_recent_count"`
	Statistics                       Statistics             `json:"statistics"`
	SupportLevel                     int                    `json:"support_level"`
	UserAchievements                 []UserAchievements     `json:"user_achievements"`
	RankHistory                      RankHistory            `json:"rank_history"`
	RankedAndApprovedBeatmapsetCount int                    `json:"ranked_and_approved_beatmapset_count"`
	UnrankedBeatmapsetCount          int                    `json:"unranked_beatmapset_count"`
}
type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}
type Badges struct {
	AwardedAt   time.Time `json:"awarded_at"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	URL         string    `json:"url"`
}
type MonthlyPlaycounts struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}
type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}
type ReplaysWatchedCounts struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}
type Rank struct {
	Country int `json:"country"`
}
type Statistics struct {
	Level                  Level       `json:"level"`
	GlobalRank             int         `json:"global_rank"`
	Pp                     float64     `json:"pp"`
	RankedScore            int64       `json:"ranked_score"`
	HitAccuracy            float64     `json:"hit_accuracy"`
	PlayCount              int         `json:"play_count"`
	PlayTime               int         `json:"play_time"`
	TotalScore             int64       `json:"total_score"`
	TotalHits              int         `json:"total_hits"`
	MaximumCombo           int         `json:"maximum_combo"`
	ReplaysWatchedByOthers int         `json:"replays_watched_by_others"`
	IsRanked               bool        `json:"is_ranked"`
	GradeCounts            GradeCounts `json:"grade_counts"`
	CountryRank            int         `json:"country_rank"`
	Rank                   Rank        `json:"rank"`
}
type UserAchievements struct {
	AchievedAt    time.Time `json:"achieved_at"`
	AchievementID int       `json:"achievement_id"`
}
type RankHistory struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

func (c *Client) GetUserKudosu(user, limit int, offset string) (kudosu []*KudosuHistory, err error) {
	params := map[string]string{
		"limit":  strconv.Itoa(limit),
		"offset": offset,
	}
	b, err := c.request(http.MethodGet, "users/"+strconv.Itoa(user)+"/kudosu", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&kudosu)
	return
}

func (c *Client) GetUserScores(user int, scoreType ScoreType, includeFails bool, mode Gamemode, limit int, offset string) (scores []*Score, err error) {
	params := map[string]string{
		"mode":          mode.String(),
		"include_fails": strconv.FormatBool(includeFails),
		"limit":         strconv.Itoa(limit),
		"offset":        offset,
	}
	b, err := c.request(http.MethodGet, "users/"+strconv.Itoa(user)+"/scores/"+scoreType.String(), params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&scores)
	return
}

func (c *Client) GetUserBeatmaps(user int, beatmapType GetUserBeatmapsType, limit int, offset string) (beatmaps []*Beatmapset, err error) {
	if beatmapType == GetUserBeatmapsTypeMostPlayed {
		return nil, fmt.Errorf("GetUserBeatmapsTypeMostPlayed is not supported in GetUserBeatmaps, use GetUserBeatmapsMostPlayed instead")
	}
	params := map[string]string{
		"type":   beatmapType.String(),
		"limit":  strconv.Itoa(limit),
		"offset": offset,
	}
	b, err := c.request(http.MethodGet, "users/"+strconv.Itoa(user)+"/beatmaps", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmaps)
	return
}

func (c *Client) GetUserBeatmapsMostPlayed(user int, limit int, offset string) (beatmaps []*BeatmapPlaycount, err error) {
	params := map[string]string{
		"type":   GetUserBeatmapsTypeMostPlayed.String(),
		"limit":  strconv.Itoa(limit),
		"offset": offset,
	}
	b, err := c.request(http.MethodGet, "users/"+strconv.Itoa(user)+"/beatmaps", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&beatmaps)
	return
}

func (c *Client) GetUserRecentActivity(user int, limit int, offset string) (activity []*Event, err error) {
	params := map[string]string{
		"limit":  strconv.Itoa(limit),
		"offset": offset,
	}
	b, err := c.request(http.MethodGet, "users/"+strconv.Itoa(user)+"/recent-activity", params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&activity)
	return
}

// queryType can be either id or username to limit lookup by their respective type.
// Passing empty or invalid value will result in id lookup followed by username lookup if not found.
func (c *Client) GetUser(user int, mode Gamemode, queryType string) (fullUser *User, err error) {
	if mode == GamemodeDefault || mode == "" {
	}
	params := map[string]string{}
	if queryType != "" {
		if queryType == "id" || queryType == "username" {
			params["key"] = queryType
		}
	}

	b, err := c.request(http.MethodGet, "users/"+strconv.Itoa(user)+"/"+mode.String(), params, nil)
	if err != nil {
		return
	}

	err = json.NewDecoder(b).Decode(&fullUser)
	return
}
