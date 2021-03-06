package laterpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {

	c := LaterPayClient{
		Id:        "myid",
		SecretKey: []byte("mysecret"),
		AddURL:    "https://foo.bar/add",
		AccessURL: "https://api.foo.bar/access",
		WebRoot:   "https://foo.bar",
	}

	i := ItemDefinition{
		Id:      "14",
		Pricing: "EUR23",
		URL:     "https://my.server.com",
		Title:   "My Article",
	}

	url, err := c.Add(i)
	assert.Nil(t, err)
	t.Logf("Add URL: %q\n", url)

	c.Access("foobar213", "1", "2", "3")

}
