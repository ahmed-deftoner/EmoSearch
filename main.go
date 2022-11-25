package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

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

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return

	//init
	ctx := context.Background()
	client := getClient(ctx)

	//flags
	albumPtr := flag.Bool("album", false, "search for album")
	playlistPtr := flag.Bool("playlist", false, "search for playlist")

	numbPtr := flag.Int("num", 1, "number of songs")
	forkPtr := flag.String("emo", "sad", "get emotion")

	flag.Parse()

	fmt.Println("numb:", *albumPtr)
	fmt.Println("numb:", *playlistPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("tail:", flag.Args())

	//major  code
	if *albumPtr {
		results, err := client.Search(ctx, "youth", spotify.SearchTypeAlbum)
		if err != nil {
			log.Fatal(err)
		}
		audiofeature.GetSadSongs(ctx, results, client)
	}
	if *playlistPtr {
		results, err := client.Search(ctx, "youth", spotify.SearchTypePlaylist)
		if err != nil {
			log.Fatal(err)
		}
		audiofeature.GetSadSongs(ctx, results, client)
	}

}
