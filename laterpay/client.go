package laterpay

import (
	"fmt"
	"net/url"

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
	AddURL    string
	AccessURL string
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
	u, err := c.makeURL(c.AddURL, "GET", &i)
	if err != nil {
		return "", err
	}
	return c.dialogURL(u), nil
}

type accessParams struct {
	Ids   []string `url:"article_id"`
	Token string   `url:"lptoken"`
}

func (c *LaterPayClient) Access(token string, ids ...string) map[string]bool {
	ap := accessParams{
		Token: token,
		Ids:   ids,
	}
	u, err := c.makeURL(c.AccessURL, "GET", ap)
	if err != nil {
		return nil
	}
	fmt.Println(u)
	return nil
}

type tokenParams struct {
	Redir    string `url:"redir"`
	Merchant string `url:"cp"`
}

func (c *LaterPayClient) GetTokenURL(redir string) (string, error) {
	tp := tokenParams{redir, c.Id}
	return c.makeURL("https://api.sandbox.laterpaytest.net/gettoken", "GET", tp)
}
