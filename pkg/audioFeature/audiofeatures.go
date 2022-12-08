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

func GetAcousticSongs(arr []Features) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].acoustic > arr[j].acoustic
	})
}

func GetInstrumentalSongs(arr []Features) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].instrumental > arr[j].instrumental
	})
}

func GetVocalSongs(arr []Features) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].vocal > arr[j].vocal
	})
}

func GetSongs(ctx context.Context, results *spotify.SearchResult, client *spotify.Client, emo string, num int) []string {
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

		switch emo {
		case "sad":
			GetSadSongs(arr)
		case "happy":
			GetHappySongs(arr)
		case "acoustic":
			GetAcousticSongs(arr)
		case "instrumental":
			GetInstrumentalSongs(arr)
		case "intense":
			GetIntenseSongs(arr)
		case "vocal":
			GetVocalSongs(arr)
		default:
			GetSadSongs(arr)
		}

		songArr := make([]string, res.Total+1)

		for i := 0; i < res.Total; i++ {
			songArr[i] = arr[i].Song
		}
		return songArr[:num]
	} else if results.Playlists != nil {
		item := results.Playlists.Playlists[0]
		res, err := client.GetPlaylistItems(ctx, item.ID, spotify.Market("US"))
		arr := make([]Features, res.Total+1)
		artist := make([]string, res.Total+1)

		var i int = 0

		if err != nil {
			fmt.Println("error getting tracks ....", err.Error())
		}
		for _, item := range res.Items {
			x, err := client.GetAudioFeatures(ctx, item.Track.Track.ID)
			if err != nil {
				fmt.Println("error getting audio features...", err.Error())
			}
			artist[i] = item.Track.Track.Artists[0].Name
			arr[i].Song = item.Track.Track.Name
			arr[i].valence = x[0].Valence
			arr[i].acoustic = x[0].Acousticness
			arr[i].danceable = x[0].Danceability
			arr[i].intense = x[0].Energy
			arr[i].instrumental = x[0].Instrumentalness
			arr[i].vocal = x[0].Speechiness
			i++
		}

		switch emo {
		case "sad":
			GetSadSongs(arr)
		case "happy":
			GetHappySongs(arr)
		case "acoustic":
			GetAcousticSongs(arr)
		case "instrumental":
			GetInstrumentalSongs(arr)
		case "intense":
			GetIntenseSongs(arr)
		case "vocal":
			GetVocalSongs(arr)
		default:
			GetSadSongs(arr)
		}

		songArr := make([]string, res.Total+1)

		for i := 0; i < res.Total; i++ {
			songArr[i] = arr[i].Song + "\t-\t" + artist[i]
		}
		return songArr[:num]
	}
	return nil
}
