package parser

import (
	"context"
	"log"
	"os"
	"strings"
	"strconv"

	"github.com/joho/godotenv"
	audiofeature "github.com/spoifytest/pkg/audioFeature"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2auth"
	golang.org/x/oauth2/clientcredentials"


	godotenv.Load()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	return spotify.New(httpClient)
}

func HandleRequest(message string) []string {
	//init
	ctx := context.Background()
	client := getClient(ctx)

	tokens := strings.Split(message, ":")
	if tokens[1] == "a" {
		results, err := client.Search(ctx, tokens[0], spotify.SearchTypeAlbum)
		results, err :=client.Search(ctx, tokens[0], spotify.SearchTypeAlbum)
		if err != nil {
			og.Fatal(err)
		}
		return audiofeature.GetSong(ctx, results, client,  tokens[2], strconv.Aoi(tokens[3]))
	} else if tokens[1] == "p" {
		results, err :=client.Search(ctx, tokens[0], spotify.SearchTypePlaylist)
		if err != nil {
			og.Fatal(err)
		}
		return audiofeature.GetSongs(ctx, results, client,  tokens[2], strconv.Aoi(tokens[3]))
	}
	eturn nil
}
