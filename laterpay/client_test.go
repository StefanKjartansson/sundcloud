package laterpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {

	c := LaterPayClient{
		Id:        "myid",
		secretKey: []byte("mysecret"),
		addURL:    "https://foo.bar/add",
		accessURL: "https://api.foo.bar/access",
		webRoot:   "https://foo.bar",
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
