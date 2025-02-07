package config

type Layout struct {
	CellItemCount int
	AspectRatio   bool
	ThumbSize     int
}

var CurrentLayout = Layout{
	CellItemCount: 5,
	AspectRatio:   true,
	ThumbSize:     270,
}

func LayoutInit() {
	switch CurrentLayout.CellItemCount {
	case 1:
	case 2:
		CurrentLayout.ThumbSize = 540
		break
	case 3:
	case 4:
	case 5:
		CurrentLayout.ThumbSize = 270
		break
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
		CurrentLayout.ThumbSize = 135
		break
	case 11:
	case 12:
	case 13:
	case 14:
		CurrentLayout.ThumbSize = 70
		break
	}
}
