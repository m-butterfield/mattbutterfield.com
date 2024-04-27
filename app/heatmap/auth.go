package heatmap

import (
	"context"
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/strava-api/swagger"
	"golang.org/x/oauth2"
	"log"
	"os"
)

var clientID = os.Getenv("STRAVA_CLIENT_ID")
var clientSecret = os.Getenv("STRAVA_CLIENT_SECRET")

func getNewToken() *oauth2.Token {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"activity:read_all"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.strava.com/oauth/authorize",
			TokenURL: "https://www.strava.com/api/v3/oauth/token",
		},
		RedirectURL: "http://localhost/exchange_token",
	}

	// use PKCE to protect against CSRF attacks
	// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
	verifier := oauth2.GenerateVerifier()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	token, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}
	return token
}

type tokenSaverSource struct {
	ds                 data.Store
	currentAccessToken string
	tokenSource        oauth2.TokenSource
}

func (t *tokenSaverSource) Token() (*oauth2.Token, error) {
	token, err := t.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	if token.AccessToken != t.currentAccessToken {
		if err := t.ds.UpdateStravaAccessToken(&data.StravaAccessToken{
			ID:           "main",
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			Expiry:       token.Expiry,
		}); err != nil {
			return nil, err
		}
	}
	return token, nil
}

func getOAuth2Token(ds data.Store) (oauth2.TokenSource, error) {
	token, err := ds.GetStravaAccessToken("main")
	if err != nil {
		return nil, err
	}

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://www.strava.com/api/v3/oauth/token",
		},
	}

	oauth2Token := &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	tokenSource := conf.TokenSource(context.Background(), oauth2Token)

	return &tokenSaverSource{
		ds:                 ds,
		currentAccessToken: token.AccessToken,
		tokenSource:        tokenSource,
	}, nil
}

func getStravaClient(ds data.Store) (*swagger.APIClient, context.Context, error) {
	token, err := getOAuth2Token(ds)
	if err != nil {
		return nil, nil, err
	}

	client := swagger.NewAPIClient(swagger.NewConfiguration())
	auth := context.WithValue(context.Background(), swagger.ContextOAuth2, token)

	return client, auth, nil
}
