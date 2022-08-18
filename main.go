package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	ctx := context.Background()
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
	client := spotify.New(httpClient)
	//major  code
	results, err := client.Search(ctx, "paramore", spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatal(err)
	}

	// handle album results
	if results.Albums != nil {
		fmt.Println("Albums:")
		item := results.Albums.Albums[0]
		res, err := client.GetAlbumTracks(ctx, item.ID, spotify.Market("US"))

		if err != nil {
			fmt.Println("error getting tracks ....", err.Error())
		}
		for _, item := range res.Tracks {
			fmt.Println(item.Name)
		}

		fmt.Println(item.ID.String())
		/*	for _, item := range results.Albums.Albums {
			fmt.Println("   ", item.Name)
		}*/
	}
}
