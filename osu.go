package osugo

const BASE_URL = "https://osu.ppy.sh/api/v2/"

type RankedStatus int

const (
	RankedStatusGraveyard RankedStatus = iota - 2
	RankedStatusWIP
	RankedStatusPending
	RankedStatusRanked
	RankedStatusApproved
	RankedStatusQualified
	RankedStatusLoved
)

func (r RankedStatus) String() string {
	switch r {
	case RankedStatusGraveyard:
		return "graveyard"
	case RankedStatusWIP:
		return "wip"
	case RankedStatusPending:
		return "pending"
	case RankedStatusRanked:
		return "ranked"
	case RankedStatusApproved:
		return "approved"
	case RankedStatusQualified:
		return "qualified"
	case RankedStatusLoved:
		return "loved"
	default:
		return "unknown"
	}
}

type Gamemode string

const (
	GamemodeStandard Gamemode = "osu"
	GamemodeTaiko    Gamemode = "taiko"
	GamemodeCatch    Gamemode = "fruits"
	GamemodeMania    Gamemode = "mania"
)

func (g Gamemode) Description() string {
	switch g {
	case GamemodeStandard:
		return "osu!standard"
	case GamemodeTaiko:
		return "osu!taiko"
	case GamemodeCatch:
		return "osu!catch"
	case GamemodeMania:
		return "osu!mania"
	default:
		return "unknown"
	}
}

func (g Gamemode) String() string {
	return string(g)
}

type ScoreType string

const (
	ScoreTypeBest   ScoreType = "best"
	ScoreTypeFirsts ScoreType = "firsts"
	ScoreTypeRecent ScoreType = "recent"
)

func (s ScoreType) String() string {
	return string(s)
}

type GetUserBeatmapsType string

const (
	GetUserBeatmapsTypeFavourite  GetUserBeatmapsType = "favourite"
	GetUserBeatmapsTypeGraveyard  GetUserBeatmapsType = "graveyard"
	GetUserBeatmapsTypeLoved      GetUserBeatmapsType = "loved"
	GetUserBeatmapsTypeMostPlayed GetUserBeatmapsType = "most_played"
	GetUserBeatmapsTypePending    GetUserBeatmapsType = "pending"
	GetUserBeatmapsTypeRanked     GetUserBeatmapsType = "ranked"
)

func (t GetUserBeatmapsType) String() string {
	return string(t)
}
