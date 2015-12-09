package laterpay

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignAndEncode(t *testing.T) {
	u, err := url.Parse("https://api.sandbox.laterpaytest.net/access?cp=234")
	assert.Nil(t, err)
	secret := []byte("mysecret")
	err = signURL(secret, "GET", u)
	assert.Nil(t, err)
}

func messageTest(t *testing.T, method, expected string, baseURL string, params url.Values) {
	u, err := url.Parse(baseURL)
	assert.Nil(t, err)
	msg, err := baseMessage(method, u, params)
	assert.Nil(t, err)
	t.Logf("Encoded MSG: %s\n", msg)
	assert.Equal(t, msg, expected, "they should be equal")
}

func TestCreateMessage(t *testing.T) {
	expected := "POST&https%3A%2F%2Fendpoint.com%2F%C4%85pi&par%25C4%2584m1%3Dvalu%25C4%2598"
	baseURL := "https://endpoint.com/ąpi"
	q := url.Values{}
	q.Set("parĄm1", "valuĘ")
	messageTest(t, "POST", expected, baseURL, q)
}

func TestCreateMessageSortCombime(t *testing.T) {
	expected := "POST&https%3A%2F%2Fendpoint.com%2Fapi&par%25C4%2584m1%3Dvalu%25C4%2598%26param2%3Dvalue2%26param2%3Dvalue3%26param3%3Dwith%2520a%2520space"
	baseURL := "https://endpoint.com/api"
	q := url.Values{}
	q.Set("parĄm1", "valuĘ")
	q.Set("param2", "value2")
	q.Add("param2", "value3")
	q.Set("param3", "with a space")
	messageTest(t, "POST", expected, baseURL, q)
}

func TestCreateMessageWrongMethod(t *testing.T) {
	u, err := url.Parse("https://foo.bar")
	assert.Nil(t, err)
	_, err = baseMessage("WRONG", u, url.Values{})
	assert.NotNil(t, err)
}

func TestSignAndEncode2(t *testing.T) {
	u, err := url.Parse("https://api.sandbox.laterpaytest.net/access?cp=234")
	assert.Nil(t, err)
	secret := []byte("mysecret")
	err = signURL(secret, "GET", u)
	assert.Nil(t, err)
}
