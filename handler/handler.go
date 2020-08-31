package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// http client
var httpClient *http.Client

// Twitter client
var client *twitter.Client

func init() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("clientID"),
		ClientSecret: os.Getenv("clientSecret"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient = config.Client(oauth2.NoContext)
	client = twitter.NewClient(httpClient)
}

// Favorite is get twitter favorite list
func Favorite(c echo.Context) error {
	params := &twitter.FavoriteListParams{
		ScreenName: c.QueryParam("screen_name"),
	}
	// get favorites list
	tweet, resp, err := client.Favorites.List(params)
	fmt.Println(tweet)
	fmt.Println(resp)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusOK, tweet)
	}
	return c.JSON(http.StatusBadRequest, err)
}
