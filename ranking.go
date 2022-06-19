package osugo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type RankingsFull struct {
	Cursor struct {
		Page int `json:"page,omitempty"`
	} `json:"cursor,omitempty"`

	Rankings []*RankingEntry `json:"ranking,omitempty"`
}
type RankingEntry struct {
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
	User                   RankingUser `json:"user"`
}
type Level struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}
type GradeCounts struct {
	Ss  int `json:"ss"`
	SSH int `json:"ssh"`
	S   int `json:"s"`
	Sh  int `json:"sh"`
	A   int `json:"a"`
}
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type Cover struct {
	CustomURL string      `json:"custom_url"`
	URL       string      `json:"url"`
	ID        interface{} `json:"id"`
}
type RankingUser struct {
	AvatarURL     string      `json:"avatar_url"`
	CountryCode   string      `json:"country_code"`
	DefaultGroup  string      `json:"default_group"`
	ID            int         `json:"id"`
	IsActive      bool        `json:"is_active"`
	IsBot         bool        `json:"is_bot"`
	IsDeleted     bool        `json:"is_deleted"`
	IsOnline      bool        `json:"is_online"`
	IsSupporter   bool        `json:"is_supporter"`
	LastVisit     time.Time   `json:"last_visit"`
	PmFriendsOnly bool        `json:"pm_friends_only"`
	ProfileColour interface{} `json:"profile_colour"`
	Username      string      `json:"username"`
	Country       Country     `json:"country"`
	Cover         Cover       `json:"cover"`
}

func (c *Client) GetRankings(mode Gamemode, rankingType RankingType, country string, maxPages int) (rankings []*RankingEntry, err error) {
	params := map[string]string{}
	rankings = make([]*RankingEntry, 0)

	if country != "" {
		params["country"] = country
	}

	if maxPages > 200 {
		maxPages = 200
	}

	for i := 1; i <= maxPages; i++ {
		params["page"] = strconv.Itoa(i)
		b, err := c.request(http.MethodGet, "rankings/"+mode.String()+"/"+rankingType.String(), params, nil)
		if err != nil {
			return nil, err
		}

		// unmarshal the response into a temp struct then append to rankings
		var temp *RankingsFull
		err = json.NewDecoder(b).Decode(&temp)
		if err != nil {
			return nil, err
		}
		rankings = append(rankings, temp.Rankings...)
	}

	return rankings, nil

}
