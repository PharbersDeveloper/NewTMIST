package AuthDaemon

import (
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type AuthClient struct {
	AuthURL  string
	TokenURL string
}

func (au AuthClient) NewAuthClientDaemon(args map[string]string) *AuthClient {
	return &AuthClient{
		AuthURL:     args["auth_url"],
		TokenURL:     args["token_url"],
	}
}

func  (au *AuthClient)  ConfigFromURIParameter(r *http.Request) *oauth2.Config {
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)

	config := &oauth2.Config {
		ClientID: 		queryForm["client_id"][0],
		ClientSecret: 	queryForm["client_secret"][0],
		RedirectURL: 	queryForm["redirect_uri"][0],
		Endpoint:		oauth2.Endpoint{
			AuthURL: 	au.AuthURL,
			TokenURL: 	au.TokenURL,
			AuthStyle: 	oauth2.AuthStyleAutoDetect, //风格，等会再看
		},
	}

	return config
}