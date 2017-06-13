package robinhood

import (
	"net/http"
	"time"

	"astuart.co/clyde"
)

const (
	epBase       = "https://api.robinhood.com/"
	epLogin      = epBase + "api-token-auth/"
	epAccounts   = epBase + "accounts/"
	epQuotes     = epBase + "quotes/"
	epPortfolios = epBase + "portfolios/"
)

type Client struct {
	Token   string
	Account *Account
	*http.Client
}

func Dial(t TokenGetter) (*Client, error) {
	tkn, err := t.GetToken()
	if err != nil {
		return nil, err
	}

	return &Client{
		Token:  tkn,
		Client: &http.Client{Transport: clyde.HeaderRoundTripper{"Authorization": "Token " + tkn}},
	}, nil
}

type Meta struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
}