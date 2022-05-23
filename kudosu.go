package osugo

import "time"

type KudosuHistoryAction string

const (
	KudosuHistoryActionGive       KudosuHistoryAction = "give"
	KudosuHistoryActionVoteGive   KudosuHistoryAction = "vote.give"
	KudosuHistoryActionReset      KudosuHistoryAction = "reset"
	KudosuHistoryActionVoteReset  KudosuHistoryAction = "vote.reset"
	KudosuHistoryActionRevoke     KudosuHistoryAction = "revoke"
	KudosuHistoryActionVoteRevoke KudosuHistoryAction = "vote.revoke"
)

type KudosuGiver struct {
	Url      string `json:"url"`
	Username string `json:"username"`
}

type KudosuPost struct {
	Url   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

type KudosuHistory struct {
	Id        int                 `json:"id"`
	Action    KudosuHistoryAction `json:"action"`
	Amount    int                 `json:"amount"`
	Model     string              `json:"model"`
	CreatedAt time.Time           `json:"created_at"`
	Giver     *KudosuGiver        `json:"giver,omitempty"`
	Post      *KudosuPost         `json:"post,omitempty"`
}
