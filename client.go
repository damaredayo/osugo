package osugo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var ErrNoAuth = fmt.Errorf("no auth object")

type Client struct {
	id     string
	secret string
	Auth   *ClientAuth

	http *http.Client
}

type ClientAuth struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func NewClient(id, secret string) (*Client, error) {
	client := &Client{
		http:   http.DefaultClient,
		id:     id,
		secret: secret,
	}
	err := client.FetchToken(id, secret)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) FetchToken(id, secret string) error {
	data := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		Scope        string `json:"scope"`
	}{
		ClientID:     id,
		ClientSecret: secret,
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

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &c.Auth)
	return nil
}

func (c *Client) request(method string, path string, params map[string]string, in []byte) (body io.ReadCloser, err error) {
	if c.Auth == nil {
		return nil, ErrNoAuth
	}

	req, err := http.NewRequest(method, BASE_URL+path, bytes.NewBuffer(in))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.Auth.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
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
	case 200:
	case 401:
		err := c.FetchToken(c.id, c.secret)
		if err != nil {
			return nil, err
		}
		return c.request(method, path, params, in)
	case 429:
		// professional ratelimiting solution :D
		time.Sleep(time.Second * 5)
		return c.request(method, path, params, in)

	default:
		return nil, fmt.Errorf("%s %s: %s", method, path, resp.Status)
	}
	body = resp.Body
	return
}
