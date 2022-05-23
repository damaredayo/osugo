package osugo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
