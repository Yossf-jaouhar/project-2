package mosdef

func Color(flag string) [3]int {
	RGB := [3]int{}
	if len(flag) > 8 && flag[0:8] == "--color=" {
		switch flag[8:] {
		case "red":
			RGB[0] = 255
			RGB[1] = 0
			RGB[2] = 0
		case "blue":
			RGB[0] = 0
			RGB[1] = 0
			RGB[2] = 255
		case "green":
			RGB[0] = 0
			RGB[1] = 255
			RGB[2] = 0
		case "coral":
			RGB[0] = 255
			RGB[1] = 127
			RGB[2] = 80
		case "pink":
			RGB[0] = 255
			RGB[1] = 0
			RGB[2] = 127
		case "purple":
			RGB[0] = 76
			RGB[1] = 0
			RGB[2] = 153
		case "yellow":
			RGB[0] = 255
			RGB[1] = 255
			RGB[2] = 0
		case "orange":
			RGB[0] = 255
			RGB[1] = 69
			RGB[2] = 0
		case "brown":
			RGB[0] = 139
			RGB[1] = 69
			RGB[2] = 19
		}
	}
	return RGB
}
