package repository

import (
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/utils"
)

type InstagramPostDTO struct {
	Caption string            `json:"caption"`
	Avatar  model.PhotoBase   `json:"avatar"`
	Photos  []model.PhotoBase `json:"photos"`
}

//curl -fsSL https://ollama.com/install.sh | sh

var instagramPostDTO1 InstagramPostDTO
var instagramPostDTO2 InstagramPostDTO
var instagramPostDTO3 InstagramPostDTO

func GetInstagram(folder string, avatar string) InstagramPostDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var localInstagramPostDTO InstagramPostDTO

	//if count > 10 {
	//	count = 10
	//}

	var index = 0
	var nameIndex = 0

	var photo = model.PhotoBase{}
	photo.Name = avatar
	photo.FileType = ".jpg"
	photo.Width = 50
	photo.Height = 50
	photo.ThumbSize = 270
	photo.Dx = 20
	photo.Dy = 20
	photo.Circle = true
	photo.Key = -1
	photo.PaintWidth = dp(33)
	photo.PaintHeight = dp(33)
	localInstagramPostDTO.Avatar = photo

	localInstagramPostDTO.Caption = "Mahdi"

	for i := 0; i < count; i++ {
		if nameIndex >= len(utils.FackNames) {
			nameIndex = 0
		}

		var photo = model.PhotoBase{}
		photo = photos[index]
		photo.ThumbSize = 540
		photo.Crop = 1
		photo.Round = 0
		photo.Key = -1
		photo.PaintWidth = 1000
		photo.PaintHeight = 1000

		localInstagramPostDTO.Photos = append(localInstagramPostDTO.Photos, photo)

		nameIndex++
		index++
	}

	index = 0

	return localInstagramPostDTO
}
