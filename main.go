package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

type MyArr struct {
	Song  string
	value float32
}

func Initialize() *spotify.Client {
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
	return spotify.New(httpClient)
}

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
	results, err := client.Search(ctx, "youth", spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatal(err)
	}

	// handle album results
	if results.Albums != nil {
		fmt.Println("Albums:")
		item := results.Albums.Albums[0]
		res, err := client.GetAlbumTracks(ctx, item.ID, spotify.Market("US"))
		arr := make([]MyArr, res.Total+1)
		var i int = 0

		if err != nil {
			fmt.Println("error getting tracks ....", err.Error())
		}
		for _, item := range res.Tracks {
			x, err := client.GetAudioFeatures(ctx, item.ID)
			if err != nil {
				fmt.Println("error getting audio features...", err.Error())
			}
			arr[i].Song = item.Name
			arr[i].value = x[0].Valence
			//fmt.Println(arr[i])
			i++
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].value < arr[j].value
		})
		for i := 0; i < res.Total; i++ {
			fmt.Println(arr[i].Song)
		}
		/*	for _, item := range results.Albums.Albums {
			fmt.Println("   ", item.Name)
		}*/
	}
}
