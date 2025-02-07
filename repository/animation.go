package repository

import (
	"github.com/mahdi-cpp/api-go-emqx/cache"
	"github.com/mahdi-cpp/api-go-emqx/model"
	"github.com/mahdi-cpp/api-go-emqx/utils"
)

var animationDTO AnimationDTO

type AnimationDTO struct {
	Caption    string      `json:"name"`
	Animations []Animation `json:"animations"`
}

type Animation struct {
	Name  string          `json:"name"`
	Photo model.PhotoBase `json:"photo"`
}

func GetAnimation(folder string) AnimationDTO {

	var file = "data.txt"
	var photos = cache.ReadOfFile(folder, file)
	var count = len(photos)
	var dto AnimationDTO

	//if count > 50 {
	//	count = 50
	//}

	var index = 0
	var nameIndex = 0

	for i := 0; i < count; i++ {
		var animation Animation
		if nameIndex >= len(utils.MovieNames) {
			nameIndex = 0
		}

		animation.Name = utils.MovieNames[nameIndex]

		animation.Photo = photos[index]
		animation.Photo.Key = -1
		animation.Photo.ThumbSize = 540
		animation.Photo.PaintWidth = dp(70)
		animation.Photo.PaintHeight = dp(120)

		dto.Animations = append(dto.Animations, animation)
		nameIndex++
		index++
	}

	index = 0

	return dto
}
