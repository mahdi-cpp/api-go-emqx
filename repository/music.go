package repository

import (
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/utils"
)

var musicDTO MusicDTO

type MusicDTO struct {
	Caption string  `json:"name"`
	Musics  []Music `json:"musics"`
}

type Music struct {
	Artist string          `json:"artist"`
	Track  string          `json:"track"`
	Cover  model.PhotoBase `json:"cover"`
}

func GetMusics(folder string) MusicDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto MusicDTO

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var music Music
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		music.Artist = utils.MovieNames[nameIndex]

		music.Cover = photos[index]
		music.Cover.Key = -1
		//music.Cover.Crop = 1
		//music.Cover.Round = 10
		music.Cover.ThumbSize = 1080
		//music.Cover.PaintWidth = dp(70)
		//music.Cover.PaintHeight = dp(120)

		dto.Musics = append(dto.Musics, music)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
