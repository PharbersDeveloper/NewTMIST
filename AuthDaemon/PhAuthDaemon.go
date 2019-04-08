package AuthDaemon

import (
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strings"
)

type AuthClient struct {
	AuthURL  string
	TokenURL string
}

// 原有库提供的Token不会序列化Scope，因扩展
type PhToken struct {
	oauth2.Token
	Scope 	string `json:"scope"`
	AccountID string `json:"account_id"`
}

func (au AuthClient) NewAuthClientDaemon(args map[string]string) *AuthClient {
	return &AuthClient{
		AuthURL:     args["auth_url"],
		TokenURL:    args["token_url"],
	}
}

func  (au *AuthClient)  ConfigFromURIParameter(r *http.Request) *oauth2.Config {
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)

	config := &oauth2.Config {
		ClientID: 		findArrayByKey("client_id", queryForm),
		ClientSecret: 	findArrayByKey("client_secret", queryForm),
		RedirectURL: 	findArrayByKey("redirect_uri", queryForm),
		Scopes:			strings.Split(findArrayByKey("scope", queryForm), "#"),
		Endpoint:		oauth2.Endpoint{
			AuthURL: 	au.AuthURL,
			TokenURL: 	au.TokenURL,
			AuthStyle: 	oauth2.AuthStyleAutoDetect, //风格，等会再看
		},
	}

	return config
}

func findArrayByKey(key string, values url.Values) string {
	if r := values[key]; len(r) > 0 {
		return r[0]
	}
	return ""
}