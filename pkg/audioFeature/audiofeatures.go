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

func GetSongs() {

}

func GetSadSongs(ctx context.Context, results *spotify.SearchResult, client *spotify.Client) []string {
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
			//fmt.Println(arr[i])
			i++
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].valence < arr[j].valence
		})
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
