package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type OAuth struct {
	AuthURL string
}

func newOAuth() *OAuth {
	oauthClientID := os.Getenv("OAUTH_INSTAGRAM_CLIENT_ID")
	baseUrl := "https://api.instagram.com/oauth/authorize"
	redirectUrl := os.Getenv("OAUTH_INSTAGRAM_REDIRECT_URL")
	scope := "user_profile"
	combinedUrl := "%s?redirect_uri=%s&client_id=%s&scope=%s&response_type=code"
	authUrl := fmt.Sprintf(combinedUrl, baseUrl, url.QueryEscape(redirectUrl), oauthClientID, strings.ReplaceAll(url.QueryEscape(scope), "+", "%20"))

	return &OAuth{
		AuthURL: authUrl,
	}
}
