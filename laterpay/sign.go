package laterpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

var ErrMethodNotAllowed = errors.New("Method Not Allowed")
var ErrAlreadyEncoded = errors.New("URL already encoded")
var ALLOWED_METHODS = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"}

func baseMessage(method string, URL *url.URL, values url.Values) (string, error) {
	if !stringInSlice(method, ALLOWED_METHODS) {
		return "", ErrMethodNotAllowed
	}

	// signing sorts the q params alphabetically after encoding
	// them. A bit weird.
	p := strings.Split(url.QueryEscape(values.Encode()), "%26")
	log.Printf("Params: %q\n", p)
	for idx, s := range p {
		// for some reason spaces seem to be double encoded?
		p[idx] = strings.Replace(s, "%2B", "%2520", -1)
	}
	sort.Strings(p)

	encodedValues := strings.Join(p, "%26")

	msg := fmt.Sprintf("%s&%s&%s",
		method,
		url.QueryEscape(fmt.Sprintf("%s://%s%s", URL.Scheme, URL.Host, URL.Path)),
		encodedValues)

	return msg, nil
}

func sign(secret []byte, method string, u *url.URL) (string, error) {
	params := u.Query()

	_, ok := params["hmac"]
	if ok {
		return "", ErrAlreadyEncoded
	}
	_, ok = params["ts"]
	if !ok {
		params.Set("ts", fmt.Sprintf("%d", time.Now().Unix()))
	}

	// COPY
	v := url.Values{}
	for p, g := range params {
		for _, s := range g {
			v.Add(p, s)
		}
	}

	msg, err := baseMessage(method, u, v)
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New224, secret)
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func signURL(secret []byte, method string, u *url.URL) error {

	params := u.Query()
	_, ok := params["ts"]
	if !ok {
		params.Set("ts", fmt.Sprintf("%d", time.Now().Unix()))
	}

	signature, err := sign(secret, method, u)
	if err != nil {
		return err
	}
	//params.Set("hmac", signature)
	u.RawQuery = params.Encode() + fmt.Sprintf("&hmac=%s", signature)
	return nil
}

func verify(signature string, secret []byte, method string, u *url.URL) (bool, error) {
	expected, err := sign(secret, method, u)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(signature), []byte(expected)), nil
}
