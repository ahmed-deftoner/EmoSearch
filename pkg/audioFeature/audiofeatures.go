package audiofeature

import (
	"context"
	"fmt"
	"sort"

	"github.com/zmb3/spotify/v2"
)

type MyArr struct {
	Song  string
	value float32
}

func GetSadSongs(ctx context.Context, results *spotify.SearchResult, client *spotify.Client) {
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
