package parser

import (
	"context"

	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	audiofeature "github.com/spotifytest/pkg/audioFeature"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)


func getClient(ctx context.Context) *spotify.Client {
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

	tokens := strings.Split(message, "-")
	arr := make([]string, 1)
	arr[0] = "incorrect input format. Type 'help' to see options"
	if len(tokens) != 4 {
		return arr
	}
	num, err := strconv.Atoi(tokens[3])
	if err != nil {
		log.Fatal(err)
	}
	if tokens[1] == "a" {
		results, err := client.Search(ctx, tokens[0], spotify.SearchTypeAlbum)
		if err != nil {
			log.Fatal(err)
		}
		return audiofeature.GetSongs(ctx, results, client,  tokens[2], num)
	} else if tokens[1] == "p" {
		results, err := client.Search(ctx, tokens[0], spotify.SearchTypePlaylist)
		
		if err != nil {  
			log.Fatal(err)
		}
		return audiofeature.GetSongs(ctx, results, client,  tokens[2], num)
	}
	return nil
}