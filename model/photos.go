package model

type PhotoBase struct {
	Key         int     `json:"key" default:"-1"`
	Name        string  `json:"name"`
	FileType    string  `json:"fileType"`
	Orientation int     `json:"orientation"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Circle      bool    `json:"circle,omitempty"`
	Round       int     `json:"round,omitempty"`
	Crop        int     `json:"crop,omitempty"`
	AspectRatio float32 `json:"aspectRatio,omitempty"`
	ThumbSize   int     `json:"thumbSize,omitempty"`
	PaintWidth  float32 `json:"paintWidth,omitempty"`
	PaintHeight float32 `json:"paintHeight,omitempty"`
	Dx          float32 `json:"dx,omitempty"`
	Dy          float32 `json:"dy,omitempty"`
}

type IconBase struct {
	Key    int     `json:"key"`
	Name   string  `json:"name"`
	Width  int     `json:"width,omitempty"`
	Height int     `json:"height,omitempty"`
	Dx     float32 `json:"dx,omitempty"`
	Dy     float32 `json:"dy,omitempty"`
	Color  int     `json:"color,omitempty"`
	Alpha  int     `json:"alpha,omitempty"`
}
