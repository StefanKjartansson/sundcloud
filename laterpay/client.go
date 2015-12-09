package laterpay

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// const API version

// DOCS
type ItemDefinition struct {
	Id      string `url:"article_id"`
	Pricing string `url:"pricing"`
	URL     string `url:"url"`
	Title   string `url:"title"`
	Expiry  string `url:"expiry,omitempty"`
}

type LaterPayClient struct {
	// access cache with ttl.
	// parse access response.
	//client    http.Client
	Id        string
	SecretKey []byte
	APIRoot   string
	WebRoot   string
}

func (c *LaterPayClient) makeURL(baseURL, method string, params interface{}) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", nil
	}
	if params != nil {
		qs, err := query.Values(params)
		if err != nil {
			return "", nil
		}
		u.RawQuery = qs.Encode()
	}
	err = signURL(c.SecretKey, method, u)
	return u.String(), nil
}

func (c *LaterPayClient) dialogURL(u string) string {
	return fmt.Sprintf("%s/dialog-api?url=%s", c.WebRoot, url.QueryEscape(u))
}

func (c *LaterPayClient) Add(i ItemDefinition) (string, error) {
	u, err := c.makeURL(fmt.Sprintf("%s/dialog/add", c.WebRoot), "GET", &i)
	if err != nil {
		return "", err
	}
	return c.dialogURL(u), nil
}

type accessParams struct {
	Merchant string   `url:"cp"`
	Ids      []string `url:"article_id"`
	Token    string   `url:"lptoken"`
}

func (c *LaterPayClient) Access(token string, ids ...string) map[string]bool {
	ap := accessParams{
		Token:    token,
		Ids:      ids,
		Merchant: c.Id,
	}
	u, err := c.makeURL(fmt.Sprintf("%s/access", c.APIRoot), "GET", ap)
	if err != nil {
		return nil
	}
	fmt.Println(u)

	resp, err := http.Get(u)

	log.Printf("%q", resp)
	log.Printf("%q", err)

	return nil
}

type tokenParams struct {
	Redir     string `url:"redir"`
	Merchant  string `url:"cp"`
	TimeStamp int64  `url:"ts"`
}

func (c *LaterPayClient) GetTokenURL(redir string) (string, error) {
	tp := tokenParams{redir, c.Id, time.Now().Unix()}
	return c.makeURL("https://api.sandbox.laterpaytest.net/gettoken", "GET", tp)
}
