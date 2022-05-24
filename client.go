package osugo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var ErrNoAuth = fmt.Errorf("no auth object")

type Client struct {
	id     string
	secret string
	auth   *ClientAuth

	http *http.Client
}

type ClientAuth struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func NewClient(id, secret string) (*Client, error) {
	client := &Client{
		http: http.DefaultClient,
	}
	err := client.FetchToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) FetchToken() error {
	data := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		Scope        string `json:"scope"`
	}{
		ClientID:     c.id,
		ClientSecret: c.secret,
		GrantType:    "client_credentials",
		Scope:        "public",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return err
	}

	resp, err := c.http.Post("https://osu.ppy.sh/oauth/token", "application/json", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var auth *ClientAuth

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, auth)
	return nil
}

func (c *Client) request(method string, path string, params map[string]string, in []byte) (body io.ReadCloser, err error) {
	if c.auth == nil {
		return nil, ErrNoAuth
	}

	req, err := http.NewRequest(method, BASE_URL+path, bytes.NewBuffer(in))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	req.Header.Set("User-Agent", "osugo")

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 401:
		err := c.FetchToken()
		if err != nil {
			return nil, err
		}
		return c.request(method, path, params, in)
	}
	body = resp.Body
	return
}
