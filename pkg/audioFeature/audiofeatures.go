package audiofeature

import (
	"context"
	"fmt"
	"sort"

	"github.com/zmb3/spotify/v2"
)

type Features struct {
	Song         string
	valence      float32
	danceable    float32
	intense      float32
	instrumental float32
	vocal        float32
	acoustic     float32
}

func GetSadSongs(arr []Features) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].valence < arr[j].valence
	})
}

func GetHappySongs(arr []Features) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].valence > arr[j].valence
	})
}

func GetIntenseSongs(arr []Features) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].intense > arr[j].intense
	})
}

func GetSongs(ctx context.Context, results *spotify.SearchResult, client *spotify.Client) []string {
	if results.Albums != nil {
		item := results.Albums.Albums[0]
		res, err := client.GetAlbumTracks(ctx, item.ID, spotify.Market("US"))
		arr := make([]Features, res.Total+1)
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
			arr[i].valence = x[0].Valence
			arr[i].acoustic = x[0].Acousticness
			arr[i].danceable = x[0].Danceability
			arr[i].intense = x[0].Energy
			arr[i].instrumental = x[0].Instrumentalness
			arr[i].vocal = x[0].Speechiness
			i++
		}

		songArr := make([]string, res.Total+1)

		for i := 0; i < res.Total; i++ {
			songArr[i] = arr[i].Song
		}
		/*	for _, item := range results.Albums.Albums {
			fmt.Println("   ", item.Name)
		}*/
		return songArr
	}
	return nil
}
